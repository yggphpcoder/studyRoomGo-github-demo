package gen

import "time"

// SrMemberPackageCardOperateLog 会员-套餐操作记录
type SrMemberPackageCardOperateLog struct {
	ID            int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'主键编码'"` // 主键编码
	MerchantID    int64      `gorm:"column:merchant_id;type:bigint;not null;default:0;comment:'商家id'"`            // 商家id
	MpID          int64      `gorm:"column:mp_id;type:bigint;not null;default:0;comment:'会员-套餐id'"`               // 会员-套餐id
	MemberID      int64      `gorm:"column:member_id;type:bigint;not null;default:0;comment:'会员id'"`              // 会员id
	PackageCardID int64      `gorm:"column:package_card_id;type:bigint;not null;default:0;comment:'套餐id'"`        // 套餐id
	Operate       int8       `gorm:"column:operate;type:tinyint;not null;default:0;comment:'操作'"`                 // 操作
	Extra         *string    `gorm:"column:extra;type:varchar(1000);comment:'额外信息'"`                              // 额外信息
	Remark        *string    `gorm:"column:remark;type:varchar(255);comment:'备注'"`                                // 备注
	CreatedAt     *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	CreateBy      *uint      `gorm:"column:create_by;type:int unsigned"`
}

// TableName get sql table name.获取数据库表名
func (m *SrMemberPackageCardOperateLog) TableName() string {
	return "sr_member_package_card_operate_log"
}

// SrMemberPackageCardOperateLogColumns get sql column name.获取数据库列名
var SrMemberPackageCardOperateLogColumns = struct {
	ID            string
	MerchantID    string
	MpID          string
	MemberID      string
	PackageCardID string
	Operate       string
	Extra         string
	Remark        string
	CreatedAt     string
	CreateBy      string
}{
	ID:            "id",
	MerchantID:    "merchant_id",
	MpID:          "mp_id",
	MemberID:      "member_id",
	PackageCardID: "package_card_id",
	Operate:       "operate",
	Extra:         "extra",
	Remark:        "remark",
	CreatedAt:     "created_at",
	CreateBy:      "create_by",
}
