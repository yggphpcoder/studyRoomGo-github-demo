package config

import (
	"fmt"
	"io"
	"os"
	conf2 "studyRoomGo/internal/conf"
	"studyRoomGo/internal/data/db/mysql"
	"studyRoomGo/internal/data/gen"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Conf struct {
	Environment string
	Upload      conf2.Upload
	Config      conf2.Conf
	Data        struct {
		DB    *gorm.DB
		Redis *redis.Client
	}
}

var AppConf Conf

func NewConfig(c config.Config) {
	var bc conf2.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	AppConf.Upload.Uri = bc.Upload.GetUri()
	AppConf.Upload.Avatar = bc.Upload.Avatar
	AppConf.Upload.Qrcode = bc.Upload.Qrcode

	var loggerWriter io.Writer
	var err error
	if bc.Data.Database.LogFile == "" {
		loggerWriter = os.Stdout
	} else {
		loggerWriter, err = os.OpenFile(bc.Data.Database.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println(err)
		}
	}
	db, _ := mysql.NewGorm(bc.Data.Database.Source, loggerWriter)
	redisDB := redis.NewClient(&redis.Options{
		Addr:     bc.Data.Redis.Addr, // redis地址
		Password: bc.Data.Redis.Password,
		DB:       0, // 默认数据库
	})
	AppConf.Data.DB = db
	AppConf.Data.Redis = redisDB
	AppConf.Environment = bc.Environment
	AppConf.Config.Dir = bc.Conf.Dir

}

func Close() {
	db, _ := AppConf.Data.DB.DB()
	db.Close()
	fmt.Println("global config exit!")
}

func GetShopSetting(shopId int64) (map[string]string, error) {
	mgr := gen.SrSettingMgr(AppConf.Data.DB)
	var mapList = make(map[string]string)
	list, err := mgr.GetBatchFromShopID([]int64{shopId})
	if err != nil {
		return nil, err
	}
	for _, setting := range list {
		mapList[setting.Key] = setting.Value
	}
	return mapList, nil
}
