package types

type TicketCreate struct {
	Name       string `json:"name" validate:"required"`
	Desc       string `json:"desc" validate:"required"`
	Allocation *int   `json:"allocation" validate:"required"`
}

type TicketResponse struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Allocation int    `json:"allocation"`
}
