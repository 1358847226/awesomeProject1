package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	// 站点配置
	App struct {
		Name  string
		Port  string
		Debug bool
		Url   string
		Tcp  string
	}
	//数据库连接配置
	Conenction struct {
		Type     string
		Host     string
		Port	 string
		DataBase string
		Username string
		Password string
	}
}

var AppConfig Config

func init() {
	file, err := ioutil.ReadFile("./config/conf.yaml")
	if err != nil {
		log.Panicln("加载配置文件失败", err)
	}

	if err := yaml.Unmarshal(file, &AppConfig); err != nil {
		log.Fatalf("解析配置文件失败:", err)
	}
}
//本地
/*func Conf(conf string) string {
	if conf == "ip"{
		return "http://192.168.16.126:8080"
	}
	if conf == "server"{
		return "192.168.16.126:8080"
	}else{
		return "err"
	}
}*/

//服务器
func Conf(conf string) string {
	if conf == "ip"{
		return "https://card.itianwang.com"
	}
	if conf == "server"{
		return "127.0.0.1:8080"
	}else{
		return "err"
	}
}