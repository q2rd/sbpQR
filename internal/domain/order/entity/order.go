package entity

import (
	"github.com/q2rd/sbpQR/internal/config"
	"github.com/q2rd/sbpQR/pkg/utils"
)

type Order struct {
	RequestUID       string           `json:"rq_uid"`
	RequestCreatedAt string           `json:"rq_tm"`
	MemberID         string           `json:"member_id"`
	OrderUID         string           `json:"order_number"`
	OrderCreatedAt   string           `json:"order_created_date"`
	OrderParams      *[]OrderPosition `json:"order_params_type"`
	QrID             string           `json:"id_qr"`
	OrderSum         int              `json:"order_sum"`
	Currency         string           `json:"currency"`
	Description      string           `json:"description"`
	SBPMemberID      string           `json:"sbp_member_id"`
}

type OrderPosition struct {
	PositionName        string `json:"position_name"`
	PositionCount       int    `json:"position_count"`
	PositionPrice       int    `json:"position_price"`
	PositionDiscription string `json:"position_description"`
}

func (o *Order) GetSum() int {
	var sum int
	if o.OrderParams == nil {
		return sum
	}

	for _, position := range *o.OrderParams {
		sum += position.PositionCount * position.PositionPrice
	}
	return sum
}
func GetMockOrder(cfg *config.Config) *Order {
	order := &Order{
		RequestUID:       utils.GenerateCleanUUID(),
		RequestCreatedAt: utils.GenerateTimestamp(),
		MemberID:         cfg.MemberID,
		OrderUID:         "1",
		OrderCreatedAt:   utils.GenerateTimestamp(),
		OrderParams: &[]OrderPosition{
			{
				PositionName:        "parking_place",
				PositionCount:       2,
				PositionPrice:       10,
				PositionDiscription: "порковочное место автомобиля.",
			},
		},
		QrID:        cfg.TID,
		Currency:    "643",
		Description: "test order",
		SBPMemberID: cfg.SBPMemberID,
	}
	order.OrderSum = order.GetSum()
	return order
}
