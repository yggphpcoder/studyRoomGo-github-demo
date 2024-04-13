package main

import (
	"flag"
	"fmt"
	"os"
	config2 "studyRoomGo/common/config"
	"studyRoomGo/common/equipment"
	"studyRoomGo/common/equipment/accesscontrol"
	"studyRoomGo/common/equipment/mqtt"
	"studyRoomGo/common/equipment/smyoo"
	"studyRoomGo/internal/conf"
	"time"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
	)
}

func main() {
	flag.Parse()
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}
	defer config2.Close()
	config2.NewConfig(c) //全局通用配置

	runFile := fmt.Sprintf("/var/app/logs/studyRoomGo/run_%s_%s.log", config2.AppConf.Environment, time.Now().Format(time.DateOnly))
	f, err := os.OpenFile(runFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		f = os.Stdout
	}
	logger := log.With(log.NewStdLogger(f),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)

	equipment.NewServer(
		accesscontrol.NewServer(
			accesscontrol.WithKey("accesscontrol"),
		),
		mqtt.NewServer(
			mqtt.WithKey("mqtt"),
		),
		smyoo.NewServer(
			smyoo.WithKey("smyoo"),
		),
	)
	defer equipment.Close()

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	//dianping.CreateTuangouList("a87c1216347443cd")
	//dianping.SaveAccessToken("a87c1216347443cd")
	//dianping.Scanprepare("a87c1216347443cd")
	app, cleanup, err := wireApp(bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
