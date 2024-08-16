package types

type PurchaseCreate struct {
	Quantity int    `json:"quantity" validate:"required"`
	UserId   string `json:"user_id" validate:"required"`
}
