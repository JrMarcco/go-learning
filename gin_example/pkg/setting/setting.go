package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type App struct {
	PageSize        int
	JwtSecret       string
	RuntimeRootPath string

	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string

	PrefixUrl      string
	ExportSavePath string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DbSetting = &Database{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

var cfg *ini.File

func Setup() {

	var err error
	//cfg, err = ini.Load("conf/app.ini")
	cfg, err = ini.Load("/home/jrmarcco/workspaces/projects/goland/go_learning/gin_example/conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v\n", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DbSetting)
	mapTo("redis", RedisSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second

	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}

func mapTo(sec string, v any) {
	if err := cfg.Section(sec).MapTo(v); err != nil {
		log.Fatalf("Map to %s err: %v\n", sec, err)
	}
}
