package entity

import (
	"github.com/q2rd/sbpQR/internal/config"
	"github.com/q2rd/sbpQR/pkg/utils"
)

type OrderCreateReq struct {
	RequestUID      string           `json:"rq_uid"`
	RequestTime     string           `json:"rq_tm"`
	MemberID        string           `json:"member_id"`
	OrderNumber     string           `json:"order_number"`
	OrderCreateDate string           `json:"order_create_date"`
	OrderParams     []*OrderPosition `json:"order_params_type"`
	IDQr            string           `json:"id_qr"`
	OrderSum        int              `json:"order_sum"`
	Currency        string           `json:"currency"`
	Description     string           `json:"description"`
	SbpMemberID     string           `json:"sbp_member_id"`
}

type OrderPosition struct {
	PositionName        string `json:"position_name"`
	PositionCount       int    `json:"position_count"`
	PositionPrice       int    `json:"position_price"`
	PositionDiscription string `json:"position_description"`
}

func (o *OrderCreateReq) GetSum() int {
	var sum int
	if o.OrderParams == nil {
		return sum
	}

	for _, position := range o.OrderParams {
		sum += position.PositionCount * position.PositionPrice
	}
	return sum
}

func GetMockOrder(cfg *config.Config) *OrderCreateReq {
	order := &OrderCreateReq{
		RequestUID:      utils.GenerateCleanUUID(),
		RequestTime:     utils.GenerateTimestamp(),
		MemberID:        cfg.MemberID,
		OrderNumber:     "1",
		OrderCreateDate: utils.GenerateTimestamp(),
		OrderParams: []*OrderPosition{
			{
				PositionName:        "parking_place",
				PositionCount:       2,
				PositionPrice:       10,
				PositionDiscription: "порковочное место автомобиля.",
			},
		},
		IDQr:        cfg.TID,
		Currency:    "643",
		Description: "test order",
		SbpMemberID: cfg.SBPMemberID,
	}
	order.OrderSum = order.GetSum()
	return order
}

type OrderStatusReq struct {
	RequestUID  string `json:"rq_uid"`
	RequestTime string `json:"rq_tm"`
	SberOrdeID  string `json:"order_id"`
	TID         string `json:"tid"`
	CRMOrderID  string `json:"partner_order_number"`
}

type OrderRevocationReq struct {
	RequestUID  string `json:"rq_uid"`
	RequestTime string `json:"rq_tm"`
	OrderID     string `json:"order_id"`
}
