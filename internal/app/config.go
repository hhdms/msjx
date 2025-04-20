package app

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App struct {
		Env  string `yaml:"env"`
		Port int    `yaml:"port"`
	} `yaml:"app"`
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		MaxConn  int    `yaml:"max_conn"`
	} `yaml:"database"`
}

var AppConfig Config

func LoadConfig() {
	f, err := os.Open("configs/config.yaml")
	if err != nil {
		log.Fatalf("打开配置文件失败: %v", err)
	}
	defer f.Close()

	d := yaml.NewDecoder(f)
	if err := d.Decode(&AppConfig); err != nil {
		log.Fatalf("解析配置文件失败: %v", err)
	}

	fmt.Println("配置文件加载成功")
}
