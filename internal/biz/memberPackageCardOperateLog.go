package biz

type MemberPackageCardOperateLog struct {
	MpId      int64
	PackageId int64
	MemberId  int64
	Operate   int8
	Extra     interface{}
}

func (l *MemberPackageCardOperateLog) StatusToOperate(status int8) {
	switch status {
	case 1:
		l.Operate = MemberPackageCard_OperateActive
	}
}
