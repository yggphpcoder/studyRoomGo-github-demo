package gen

import "time"

// SrSeat 座位信息
type SrSeat struct {
	ID          int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'主键编码'"` // 主键编码
	Pid         int64      `gorm:"column:pid;type:bigint;not null;default:0;comment:'父id'"`                     // 父id
	Name        string     `gorm:"column:name;type:varchar(255);not null;default:'';comment:'座位名'"`             // 座位名
	TypeSeat    int8       `gorm:"column:type_seat;type:tinyint;not null;default:0;comment:'座位类型'"`             // 座位类型
	ShopID      int64      `gorm:"column:shop_id;type:bigint;not null;default:0;comment:'店铺id'"`                // 店铺id
	ShopAreaID  int64      `gorm:"column:shop_area_id;type:bigint;not null;default:0;comment:'店铺分区id'"`         // 店铺分区id
	Status      int8       `gorm:"column:status;type:tinyint;not null;default:0;comment:'可用状态'"`                // 可用状态
	Sort        int        `gorm:"column:sort;type:int;not null;default:0;comment:'排序'"`                        // 排序
	Remark      *string    `gorm:"column:remark;type:varchar(255);comment:'备注'"`                                // 备注
	MapPosition *string    `gorm:"column:map_position;type:varchar(255);comment:'座位图坐标'"`                       // 座位图坐标
	CreatedAt   *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt   *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt   *time.Time `gorm:"column:deleted_at;type:timestamp;comment:'删除时间'"` // 删除时间
	CreateBy    *uint      `gorm:"column:create_by;type:int unsigned"`
	UpdateBy    *uint      `gorm:"column:update_by;type:int unsigned"`
}

// TableName get sql table name.获取数据库表名
func (m *SrSeat) TableName() string {
	return "sr_seat"
}

// SrSeatColumns get sql column name.获取数据库列名
var SrSeatColumns = struct {
	ID          string
	Pid         string
	Name        string
	TypeSeat    string
	ShopID      string
	ShopAreaID  string
	Status      string
	Sort        string
	Remark      string
	MapPosition string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
	CreateBy    string
	UpdateBy    string
}{
	ID:          "id",
	Pid:         "pid",
	Name:        "name",
	TypeSeat:    "type_seat",
	ShopID:      "shop_id",
	ShopAreaID:  "shop_area_id",
	Status:      "status",
	Sort:        "sort",
	Remark:      "remark",
	MapPosition: "map_position",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
	CreateBy:    "create_by",
	UpdateBy:    "update_by",
}
