package gen

import "time"

// SrEquipment 设备信息
type SrEquipment struct {
	ID            int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'主键编码'"` // 主键编码
	Pid           int64      `gorm:"column:pid;type:bigint;not null;default:0;comment:'父键'"`                      // 父键
	Name          string     `gorm:"column:name;type:varchar(255);not null;default:'';comment:'设备名'"`             // 设备名
	TypeEquipment int8       `gorm:"column:type_equipment;type:tinyint;not null;default:0;comment:'设备类型'"`        // 设备类型
	TypeRelate    int8       `gorm:"column:type_relate;type:tinyint;not null;default:0;comment:'关联类型'"`           // 关联类型
	RelateID      int64      `gorm:"column:relate_id;type:bigint;not null;default:0;comment:'关联id'"`              // 关联id
	ShopID        int64      `gorm:"column:shop_id;type:bigint;not null;default:0;comment:'店铺id'"`                // 店铺id
	Code          string     `gorm:"column:code;type:varchar(50);not null;default:'';comment:'设备码'"`              // 设备码
	Mac           string     `gorm:"column:mac;type:varchar(50);not null;default:'';comment:'mac地址'"`             // mac地址
	Port          string     `gorm:"column:port;type:varchar(50);not null;default:'';comment:'端口数'"`              // 端口数
	Status        int8       `gorm:"column:status;type:tinyint;not null;default:0;comment:'可用状态'"`                // 可用状态
	Operation     int8       `gorm:"column:operation;type:tinyint;not null;default:0;comment:'操作状态'"`             // 操作状态
	Sort          int        `gorm:"column:sort;type:int;not null;default:0;comment:'排序'"`                        // 排序
	Remark        *string    `gorm:"column:remark;type:varchar(255);comment:'备注'"`                                // 备注
	CreatedAt     *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt     *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt     *time.Time `gorm:"column:deleted_at;type:timestamp;comment:'删除时间'"` // 删除时间
	CreateBy      *uint      `gorm:"column:create_by;type:int unsigned"`
	UpdateBy      *uint      `gorm:"column:update_by;type:int unsigned"`
}

// TableName get sql table name.获取数据库表名
func (m *SrEquipment) TableName() string {
	return "sr_equipment"
}

// SrEquipmentColumns get sql column name.获取数据库列名
var SrEquipmentColumns = struct {
	ID            string
	Pid           string
	Name          string
	TypeEquipment string
	TypeRelate    string
	RelateID      string
	ShopID        string
	Code          string
	Mac           string
	Port          string
	Status        string
	Operation     string
	Sort          string
	Remark        string
	CreatedAt     string
	UpdatedAt     string
	DeletedAt     string
	CreateBy      string
	UpdateBy      string
}{
	ID:            "id",
	Pid:           "pid",
	Name:          "name",
	TypeEquipment: "type_equipment",
	TypeRelate:    "type_relate",
	RelateID:      "relate_id",
	ShopID:        "shop_id",
	Code:          "code",
	Mac:           "mac",
	Port:          "port",
	Status:        "status",
	Operation:     "operation",
	Sort:          "sort",
	Remark:        "remark",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	DeletedAt:     "deleted_at",
	CreateBy:      "create_by",
	UpdateBy:      "update_by",
}
