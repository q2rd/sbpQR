package types

type OrderStatus struct {
	RequestUID  string `json:"rq_uid"`
	RequestTime string `json:"rq_tm"`
	TID         string `json:"tid"`
	OrderNumber string `json:"order_number"`
}
