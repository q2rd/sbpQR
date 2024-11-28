package types

type OrderParams struct {
	PositionName        string `json:"position_name"`
	PositionCount       int    `json:"position_count"`
	PositionSum         int    `json:"position_sum"`
	PositionDescription string `json:"position_description"`
}

type SBPOrderCreationRequest struct {
	RequestUID      string         `json:"rq_uid"`
	RequestTime     string         `json:"rq_tm"`
	MemberID        string         `json:"member_id"`
	OrderNumber     string         `json:"order_number"`
	OrderCreateDate string         `json:"order_create_date"`
	OrderParamsType []*OrderParams `json:"order_params_type"`
	IDQr            string         `json:"id_qr"`
	OrderSum        int            `json:"order_sum"`
	Currency        string         `json:"currency"`
	Description     string         `json:"description"`
	SbpMemberID     string         `json:"sbp_member_id"`
}

func CountSum(objects []*OrderParams) int {
	var sum int
	for _, v := range objects {
		sum += v.PositionSum
	}
	return sum
}
