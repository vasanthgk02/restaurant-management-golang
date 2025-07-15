package views

import "time"

type InvoiceViewFormat struct {
	Invoice_id       string
	Order_id         string
	Payment_method   string
	Payment_status   string
	Payment_due      any
	Table_number     any
	Payment_due_date time.Time
	Order_details    any
}
