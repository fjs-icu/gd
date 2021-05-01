package main

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// json xml yaml 等格式
// 支持读写

// type Config struct {
// 	XMLName xml.Name    `json:"-"  xml:"config"`
// 	Redis   string      `json:"redis"  xml:"redis"`
// 	MySQL   MySQLConfig `json:"mysql"  `
// }
// type MySQLConfig struct {
// 	// `json,bson,xml,form:"field_1,omitempty" other:"value"`
// 	XMLName  xml.Name `json:"-" xml:"mysqlconfig"`
// 	Port     int      `json:"port" xml:"port"`
// 	Host     string   `json:"host" xml:"host"`
// 	Username string   `json:"username" xml:"username"`
// 	Password string   `json:"password" xml:"password"`
// }
type Config struct {
	XMLName xml.Name    `json:"-"  xml:"config"`
	Name    string      `xml:"name,attr"`
	Redis   string      `json:"redis"  xml:"redis"`
	MySQL   MySQLConfig `json:"mysql"  xml:"mysql" `
}
type MySQLConfig struct {
	// `json,bson,xml,form:"field_1,omitempty" other:"value"`
	Port     int    `json:"port" xml:"port"`
	Host     string `json:"host" xml:"host"`
	Username string `json:"username" xml:"username"`
	Password string `json:"password" xml:"password"`
}

func main() {
	// var config1 Config

	// vtoml := viper.New()
	// vtoml.SetConfigName("config")
	// vtoml.SetConfigType("yaml")
	// vtoml.AddConfigPath(".")

	// if err := vtoml.ReadInConfig(); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// vtoml.Unmarshal(&config1)
	// fmt.Println("read config.toml")
	vtom2 := viper.New()

	// vtom2.Set("", config1)
	vtom2.Set("Verbose", true)
	var config1 Config
	var ss MySQLConfig
	ss.Host = "host"
	ss.Password = "aa"
	ss.Port = 100
	ss.Username = "aa3"
	vtom2.Set("Config", config1)
	config1.MySQL = ss

	if err := vtom2.WriteConfigAs("./a.yaml"); err != nil {
		fmt.Println(err)
	}
	if err := vtom2.WriteConfigAs("./a.json"); err != nil {
		fmt.Println(err)
	}

	// if err := vtom2.WriteConfigAs("./a.toml"); err != nil {
	// 	fmt.Println("222", err)
	// }

	output, err := xml.MarshalIndent(config1, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	os.WriteFile("a.xml", output, 064)

	// fmt.Println("config: ", config1, "redis: ", config1.Redis)
	// vtoml.
	// vtoml := viper.New()
	// viper.SetConfigName("config")
	// viper.SetConfigType("yaml")
	// viper.AddConfigPath(".")

	// viper.Set("yaml", "this is a example of yaml")

	// viper.Set("redis.port", 4405)
	// viper.Set("redis.host", "127.0.0.1")

	// viper.Set("mysql.port", 3306)
	// viper.Set("mysql.host", "192.168.1.0")
	// viper.Set("mysql.username", "root123")
	// viper.Set("mysql.password", "root123")

	// if err := viper.WriteConfig(); err != nil {
	// 	fmt.Println(err)
	// }
}
