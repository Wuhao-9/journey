package setting

import (
	"log"

	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	Run_mode  string
	Http_port uint16

	DictPath      string
	UserAgentPath string
	DefGRNum      uint
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("fail to load config file: %s\n", err)
	}
	Run_mode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
	loadServerConf()
	loadDirsearchConf()
}

func loadServerConf() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("fail to parse section: <%s>\n", "server")
	}

	Http_port = uint16(sec.Key("HTTP_PORT").MustUint(8888))
	if Http_port == 0 {
		log.Fatalf("fail to fetch http-port: <%d>", Http_port)
	}
}

func loadDirsearchConf() {
	sec, err := Cfg.GetSection("dirsearch")
	if err != nil {
		log.Fatalf("fail to parse section: <%s>\n", "dirsearch")
	}

	DictPath = sec.Key("DictPath").String()
	UserAgentPath = sec.Key("UserAgentPath").String()
	DefGRNum = sec.Key("DefaultGRNum").MustUint(10)
}
