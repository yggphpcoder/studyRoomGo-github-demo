package data

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"studyRoomGo/internal/biz"
	"studyRoomGo/internal/conf"
	"studyRoomGo/internal/data/db/mysql"

	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/log"
)

// nil没有默认的类型，尽管它是多个类型的零值，必须显式或隐式指定每个nil的明确类型。
// // 明确.
// _ = (*struct{})(nil)
// _ = []int(nil)
// _ = map[int]bool(nil)
// _ = chan string(nil)
// _ = (func())(nil)
// _ = interface{}(nil)
// // 隐式.
// var _ *struct{} = nil
// var _ []int = nil
// var _ map[int]bool = nil
// var _ chan string = nil
// var _ func() = nil
// var _ interface{} = nil
var _ biz.MemberRepo = (*MemberRepo)(nil)
var _ biz.PackageCardRepo = (*PackageCardRepo)(nil)
var _ biz.ShopRepo = (*ShopRepo)(nil)
var _ biz.ConfigRepo = (*ConfigRepo)(nil)
var _ biz.OrderRepo = (*OrderRepo)(nil)
var _ biz.PaymentRepo = (*PaymentRepo)(nil)
var _ biz.CouponRepo = (*CouponRepo)(nil)
var _ biz.CourseRepo = (*CourseRepo)(nil)
var _ biz.StudentCourseRepo = (*StudentCourseRepo)(nil)
var _ biz.CourseSchedulingRepo = (*CourseSchedulingRepo)(nil)
var _ biz.ClassSchedulingRepo = (*ClassSchedulingRepo)(nil)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewMemberRepo,
	NewMemberPackageCardRepo,
	NewPackageCardRepo,
	NewShopRepo,
	NewSeatRepo,
	NewSeatBookRepo,
	NewConfigRepo,
	NewOrderRepo,
	NewPaymentRepo,
	NewEquipmentRepo,
	NewProfitSharingOrderRepo,
	NewCouponRepo,
	NewCourseRepo,
	NewStudentRepo,
	NewStudentCourseRepo,
	NewCourseSchedulingRepo,
	NewClassSchedulingRepo,
)

// Data .
type Data struct {
	db    *gorm.DB
	redis *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	var loggerWriter io.Writer
	if c.Database.LogFile == "" {
		loggerWriter = os.Stdout
	} else {
		loggerWriter, _ = os.OpenFile(c.Database.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	}
	db, _ := mysql.NewGorm(c.Database.Source, loggerWriter)
	redisDB := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,     // redis地址
		Password: c.Redis.Password, // 密码
		DB:       0,                // 使用默认数据库
	})
	if c.Debug {
		db = db.Debug()
	}
	d := &Data{
		db:    db,
		redis: redisDB,
	}
	cleanup := func() {
		d, _ := db.DB()
		redisDB.Close()
		err := d.Close()
		if err != nil {
			return
		}
		log.NewHelper(logger).Info("closing the data resources")
	}

	return d, cleanup, nil
}

var Omit = []string{
	"created_at", "updated_at", "deleted_at", "create_by", "update_by",
}

const (
	sortAId = iota
	sortDId
	sortASort
	sortDSort
	sortDictSort
	sortDictASort
	sortCreatAtSort
	sortCreatAtDSort
	sortStatusASort
	sortStatusDSort
	sortBookStartDayASort
	sortStartDateASort
	sortStartDateDSort
)

func SortType(i int, a ...string) string {
	var r string
	var (
		format       string
		formatSplice []string
		prefix       []any
		orderSplice  []any
	)

	if a != nil {
		format = "`%v`.%v %v"
		prefix = []any{a[0]}
	} else {
		format = "%v %v"
		a = append(a, "")
	}

	f := func(s ...string) {
		for i := 0; i < len(s); i = i + 2 {
			orderSplice = append(orderSplice, append(prefix, s[i], s[i+1])...)
			formatSplice = append(formatSplice, format)

		}
	}

	switch i {
	case sortAId:
		f("id", "asc")
	case sortDId:
		f("id", "desc")
	case sortASort:
		f("sort", "asc", "id", "asc")
	case sortDSort:
		f("sort", "desc", "id", "desc")
	case sortDictSort:
		f("dict_sort", "desc", "dict_code", "desc")
	case sortDictASort:
		f("dict_sort", "desc", "dict_code", "desc")
	case sortCreatAtSort:
		f("created_at", "asc", "id", "desc")
	case sortCreatAtDSort:
		f("created_at", "desc", "id", "desc")
	case sortStatusASort:
		f("status", "asc", "id", "desc")
	case sortStatusDSort:
		f("status", "desc", "id", "desc")
	case sortBookStartDayASort:
		f("booking_start_day", "asc", "booking_start", "asc")

	case sortStartDateASort:
		f("start_date", "asc", "start_time", "asc")

	case sortStartDateDSort:
		f("start_date", "desc", "start_time", "asc")
	}

	r = fmt.Sprintf(strings.Join(formatSplice, ","), orderSplice...)

	return r
}

func extra(e interface{}) *string {
	switch e.(type) {
	case string:
		str := fmt.Sprint(e)
		return &str
	case map[string]interface{}:
		bb, _ := json.Marshal(e)
		str := string(bb)
		return &str

	}
	return nil
}
