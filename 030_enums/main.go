package main

import "fmt"

// type OrderStatus int

// const (
// 	PENDING OrderStatus = iota
// 	SHIPPED
// 	DELIVERED
// 	CANCELLED
// )

type OrderStatus string

const (
	PENDING   OrderStatus = "pending"
	SHIPPED   OrderStatus = "shipped"
	DELIVERED OrderStatus = "delivered"
	CANCELLED OrderStatus = "cancelled"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Order status changed to ", status)
}
func main() {
	changeOrderStatus(DELIVERED)
}
