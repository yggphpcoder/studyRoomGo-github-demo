package gen

// SrSettingText 配置富文本
type SrSettingText struct {
	ID        int64  `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'主键编码'"` // 主键编码
	SettingID int64  `gorm:"column:setting_id;type:bigint;not null;default:0;comment:'配置id'"`             // 配置id
	RichText  string `gorm:"column:rich_text;type:text;not null;comment:'富文本'"`                           // 富文本
}

// TableName get sql table name.获取数据库表名
func (m *SrSettingText) TableName() string {
	return "sr_setting_text"
}

// SrSettingTextColumns get sql column name.获取数据库列名
var SrSettingTextColumns = struct {
	ID        string
	SettingID string
	RichText  string
}{
	ID:        "id",
	SettingID: "setting_id",
	RichText:  "rich_text",
}
