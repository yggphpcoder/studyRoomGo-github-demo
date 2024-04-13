package biz

type SeatBookOperateLog struct {
	SeatBookId    int64
	ShopId        int64
	SeatId        int64
	MemberId      int64
	Operate       int8
	Remark        string
	CreateByAdmin bool
	Extra         interface{}
}

func (l *SeatBookOperateLog) StatusToOperate(status int8) {
	switch status {
	case SeatBook_StatusWait:
		l.Operate = SeatBook_OperateBooked
	case SeatBook_StatusNoSeat:
		l.Operate = SeatBook_OperateWaiting
	case SeatBook_StatusSeated:
		l.Operate = SeatBook_OperateSeated
	case SeatBook_StatusWaitNext:
		l.Operate = SeatBook_OperateWaitNext
	case SeatBook_StatusComplete:
		l.Operate = SeatBook_OperateComplete
	case SeatBook_StatusCustomerCancel:
		l.Operate = SeatBook_OperateCustomerCancel
	case SeatBook_StatusSystemCancel:
		l.Operate = SeatBook_OperateSystemCancel
	}
}
