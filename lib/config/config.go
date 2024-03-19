package config

import (
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type config struct {
	Site Site
	Db   Db

	Stdout Stdout
	Log    Log
	Error  Log
}

type Site struct {
	Port string
}

type Db struct {
	Type string
	Dns  string
}

type Stdout struct {
	MinLevel int `mapstructure:"min_level"`
	MaxLevel int `mapstructure:"max_level"`
}

type Log struct {
	File       string
	MaxSize    int `mapstructure:"max_size"` //定义映射关系
	MaxBackups int `mapstructure:"max_backups"`
	MaxAge     int `mapstructure:"max_age"`
	MinLevel   int `mapstructure:"min_level"`
	MaxLevel   int `mapstructure:"max_level"`
}

var C config

// 配置变更订阅事件
var events = make(map[string]func())

var once sync.Once

func init() {
	once.Do(load) //类似静态方法，只执行一次
}

func load() {
	//初始配置
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("ini")

	//查找并读取配置
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	//监听配置变化
	viper.OnConfigChange(func(event fsnotify.Event) {
		//fmt.Println("Change!")
		if err := viper.Unmarshal(&C); err != nil {
			panic(err)
		}
		callEvent()
	})
	viper.WatchConfig()

	if err := viper.Unmarshal(&C); err != nil {
		panic(err)
	}
}

func RegisterEvent(name string, callback func()) {
	events[name] = callback
}

func UnRegisterEvent(name string, callback func()) {
	delete(events, name)
}

func callEvent() {
	//遍历事件列表中所有回调函数
	for _, callback := range events {
		callback() //传入参数调用回调函数
	}
}
