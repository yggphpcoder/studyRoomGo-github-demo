package biz

import "studyRoomGo/internal/data/gen"

type SeatLock struct {
	*gen.SrSeatLock
}

type SearchSeatLockingRequest struct {
	Id               int64
	SeatId           int64
	ShopId           int64
	MemberId         int64
	Status           []int32
	OrMemberId       int64
	GteLockDateEnd   string
	GteLockTimeEnd   string
	GteLockDateStart string
	GteLockTimeStart string
	LteLockDateEnd   string
	LteLockTimeEnd   string
	LteLockDateStart string
	LteLockTimeStart string
	UseTime          string
	Sort             int
	Page             int
	PageSize         int

	Join bool
}
