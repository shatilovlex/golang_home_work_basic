package entity

type CreateOrderParams struct {
	UserID     int32   `json:"userId"`
	ProductIds []int32 `json:"productIds"`
}
