package biz

import "time"

type MemberPackageCardUseLog struct {
	MpID     int64
	MemberID int64
	SeatID   int64
	ShopID   int64
	Value    int
	Remark   string
	Extra    interface{}
	CreatAt  *time.Time
}
