package entity

import "time"

type Order struct {
	OrderID        string
	OrderCreatedAt time.Time
	OrderParams    *[]OrderPosition
}

type OrderPosition struct {
	PositionName        string
	PositionCount       int
	PositionPrice       int
	PositionDiscription string
}

func (o *Order) Sum() int {
	var sum int
	if o.OrderParams == nil {
		return sum
	}

	for _, position := range *o.OrderParams {
		sum += position.PositionCount * position.PositionPrice
	}
	return sum
}
