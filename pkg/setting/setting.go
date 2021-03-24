package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

const (
	KEY_RUN_MODE      = "RUN_MODE"
	KEY_HTTP_PORT     = "HTTP_PORT"
	KEY_READ_TIMEOUT  = "READ_TIMEOUT"
	KEY_WRITE_TIMEOUT = "WRITE_TIMEOUT"
	KEY_TEMPLATE_PATH = "TEMPLATE_PATH"
	KEY_STATIC_PATH   = "STATIC_PATH"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	TemplatePath string
	StaticPath   string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Failed to load 'app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key(KEY_RUN_MODE).MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Failed to get section 'server': %v", err)
	}
	HTTPPort = sec.Key(KEY_HTTP_PORT).MustInt(8080)
	ReadTimeout = time.Duration(sec.Key(KEY_READ_TIMEOUT).MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key(KEY_WRITE_TIMEOUT).MustInt(60)) * time.Second
	TemplatePath = sec.Key(KEY_TEMPLATE_PATH).MustString("templates/*")
	StaticPath = sec.Key(KEY_STATIC_PATH).MustString("static")
}
