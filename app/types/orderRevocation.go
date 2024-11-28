package types

type OrderRevocation struct {
	RequestUID  string `json:"rq_uid"`
	RequestTime string `json:"rq_tm"`
	OrderID     string `json:"order_id"` // id закакза который вернул сбербанк при создании
}
