package model

import (
	"studyRoomGo/internal/data/gen"
)

type Order struct {
	*gen.SrOrder
	Total []*gen.SrOrderTotal `json:"total" gorm:"foreignKey:OrderID"`
}
