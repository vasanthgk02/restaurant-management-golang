package controller

import (
	"context"
	"net/http"
	"restaurant_management/database"
	"restaurant_management/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrderItemPack struct {
	Table_id    *string            `json:"table_id" validate:"required"`
	Order_items []models.OrderItem `json:"order_items"`
}

var orderItemCollection *mongo.Collection = database.OpenCollection(database.Client, "orderItem")

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		result, err := orderItemCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while retriving orderItems"})
			return
		}
		var allOrderItems []bson.M
		if err := result.All(ctx, &allOrderItems); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, allOrderItems)
	}
}

func GetOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		orderItemId := c.Param("id")
		var orderItem models.OrderItem

		if err := orderItemCollection.FindOne(ctx, bson.M{"order_item_id": orderItemId}).Decode(&orderItem); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, orderItem)
	}
}

func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.TODO(), 100*time.Second)
		defer cancel()

		var orderItemPack OrderItemPack
		var order models.Order

		if err := c.BindJSON(&orderItemPack); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(orderItemPack)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		orderItemsToBeInserted := []any{}
		order.Table_id = orderItemPack.Table_id
		order_id := OrderItemOrderCreator(order)

		validationErrItem := []any{}
		for _, item := range orderItemPack.Order_items {
			item.Order_id = order_id

			validationErr := validate.Struct(item)

			if validationErr != nil {
				validationErrItem = append(validationErrItem, item)
			}

			item.ID = primitive.NewObjectID()
			item.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			item.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			item.Order_item_id = item.ID.Hex()
			orderItemsToBeInserted = append(orderItemsToBeInserted, item)
		}
		if len(validationErrItem) >= 1 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create order", "items": validationErrItem})
			return
		}
		result, err := orderItemCollection.InsertMany(ctx, orderItemsToBeInserted)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while inserting records"})
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var orderItem models.OrderItem
		orderItemId := c.Param("id")

		if err := c.BindJSON(&orderItem); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while inserting records"})
			return
		}

		var updateObj bson.D
		if orderItem.Quantity != nil {
			updateObj = append(updateObj, bson.E{Key: "quantity", Value: orderItem.Quantity})
		}
		if orderItem.Food_id != nil {
			updateObj = append(updateObj, bson.E{Key: "food_id", Value: orderItem.Food_id})
		}
		orderItem.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{Key: "update_at", Value: orderItem.Updated_at})

		upsert := true
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}
		filter := bson.M{"order_item_id": orderItemId}

		update, updateErr := orderItemCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{Key: "$set", Value: updateObj},
			},
			&opt,
		)
		if updateErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while update data"})
			return
		}
		c.JSON(http.StatusOK, update)
	}
}

func GetOrderItemsByOrderId() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderId := c.Param("id")
		allOrderItems, err := ItemsByOrder(orderId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while fetching data"})
			return
		}
		c.JSON(http.StatusOK, allOrderItems)
	}
}

func ItemsByOrder(id string) (orderItems []bson.M, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	// match
	matchStage := bson.D{
		{
			Key: "$match", Value: bson.D{
				{Key: "order_id", Value: id},
			},
		},
	}

	// food
	lookupFoodStage := bson.D{
		{
			Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "food"},
				{Key: "localField", Value: "food_id"},
				{Key: "foreignField", Value: "food_id"},
				{Key: "as", Value: "food"},
			},
		},
	}
	unwinFoodStage := bson.D{
		{
			Key: "$unwind", Value: bson.D{
				{Key: "path", Value: "$food"},
				{Key: "preserveNullAndEmptyArrays", Value: true},
			},
		},
	}

	// order
	lookupOrderStage := bson.D{
		{
			Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "order"},
				{Key: "localField", Value: "order_id"},
				{Key: "foreignField", Value: "order_id"},
				{Key: "as", Value: "order"},
			},
		},
	}
	unwindOrderStage := bson.D{
		{
			Key: "$unwind", Value: bson.D{
				{Key: "path", Value: "$order"},
				{Key: "preserveNullAndEmptyArrays", Value: true},
			},
		},
	}

	// table
	lookupTableStage := bson.D{
		{
			Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "table"},
				{Key: "localField", Value: "order.table_id"},
				{Key: "foreignField", Value: "table_id"},
				{Key: "as", Value: "table"},
			},
		},
	}
	unwindTableStage := bson.D{
		{
			Key: "$unwind", Value: bson.D{
				{Key: "path", Value: "$table"},
				{Key: "preserveNullAndEmptyArrays", Value: true},
			},
		},
	}

	// project1
	projectStage1 := bson.D{
		{
			Key: "$project", Value: bson.D{
				{Key: "_id", Value: 0},
				{Key: "amount", Value: "$food.price"},
				{Key: "total_count", Value: 1},
				{Key: "food_name", Value: "$food.name"},
				{Key: "food_image", Value: "$food.food_image"},
				{Key: "table_number", Value: "$table.table_number"},
				{Key: "table_id", Value: "$table.table_id"},
				{Key: "order_id", Value: "$order.order_id"},
				{Key: "price", Value: "$food.price"},
				{Key: "quantity", Value: 1},
			},
		},
	}

	// group
	groupStage := bson.D{
		{Key: "$group", Value: bson.D{
			{
				Key: "_id", Value: bson.D{
					{Key: "order_id", Value: "$order_id"},
					{Key: "table_id", Value: "$table_id"},
					{Key: "table_number", Value: "$table_number"},
				},
			},
			{
				Key: "payment_due", Value: bson.D{
					{Key: "$sum", Value: "$amount"},
				},
			},
			{
				Key: "total_count", Value: bson.D{
					{Key: "$sum", Value: 1},
				},
			},
			{
				Key: "order_items", Value: bson.D{
					{Key: "$push", Value: "$$ROOT"},
				},
			},
		},
		},
	}

	// project2
	projectStage2 := bson.D{
		{
			Key: "$project", Value: bson.D{
				{Key: "id", Value: 1},
				{Key: "payment_due", Value: 1},
				{Key: "total_count", Value: 1},
				{Key: "table_number", Value: "$_id.table_number"},
				{Key: "order_items", Value: 1},
			},
		},
	}

	// Aggregate
	result, err := orderItemCollection.Aggregate(
		ctx,
		mongo.Pipeline{
			matchStage,
			lookupFoodStage,
			unwinFoodStage,
			lookupOrderStage,
			unwindOrderStage,
			lookupTableStage,
			unwindTableStage,
			projectStage1,
			groupStage,
			projectStage2,
		},
	)

	if err != nil {
		panic(err)
	}

	if err := result.All(ctx, &orderItems); err != nil {
		panic(err)
	}
	return
}
