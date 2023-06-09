package database

type TableName string

const (
	Appointment TableName = "appointment"
	Employee    TableName = "employee"
	Employer    TableName = "employer"
	Service     TableName = "service"
	Customer    TableName = "customer"
)
