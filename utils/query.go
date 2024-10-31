package utils

const (
	CUSTOMER_INSERT = "INSERT INTO Customer(name, address_id, status) VALUES($1, $2, $3)"
	// CUSTOMER_LIST   = "SELECT * FROM CUSTOMER"
	DRIVER_INSERT = "INSERT INTO Driver(name, address_id) VALUES($1, $2)"
	ORDER_INSERT  = "INSERT INTO Orders(customer_id, driver_id, order_date, price) VALUES($1, $2, $3, $4)"
)

const (
	ADDRESS_INSERT = "INSERT INTO Address(street, city, province) VALUES($1, $2, $3)"
)

const (
	//view order month
	VIEW_ORDER_MONTH = "SELECT DATE_TRUNC('month', order_date) AS month, COUNT(*) AS total_order FROM Orders GROUP BY month ORDER BY month"
	//view order month customer
	VIEW_ORDER_CS_MONTH = "SELECT C.name AS customer_name, DATE_TRUNC('month', order_date) AS month, COUNT(*) AS total_orders FROM Orders O JOIN Customer C ON O.customer_id = C.id GROUP BY C.name, month ORDER BY month, total_orders DESC"
	// vier addres order summary
	VIEW_ADDRESS_ORDER = "SELECT A.street, A.city, A.Province, COUNT(O.id)"
)
