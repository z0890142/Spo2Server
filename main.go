package main

import (
	"fmt"
	controller "spo2_server/controller"
	M "spo2_server/helper/MQ"
	mysql "spo2_server/helper/Mysql"
	"strings"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"

	"github.com/spf13/viper"
)

func init() {
	//log.SetLevel(log.DebugLevel)
	InitConfig()
	mysql.CreateDbConn("mysql", "root@tcp(127.0.0.1:3306)/Spo2")
}

func main() {

	M.MqttClientInit("", "", "")

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	r.LoadHTMLGlob("./views/*.html")          // complains about /static being a dir on boot
	r.LoadHTMLFiles("./views/static/*/*")     //  load the static path
	r.Static("/static", "./views/static")     // use the loaded source
	r.StaticFile("/spo2", "views/index.html") // use the loaded source

	r.GET("/db/:deviceId", controller.List)
	r.GET("/id/list", controller.ListId)
	r.POST("/insert/tag", controller.InsertTag)

	// r.POST("/mysql/:deviceId", controller.Transfer)

	r.Run(":8087")
}

func InitConfig() {
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()         // read in environment variables that match
	viper.SetEnvPrefix("gorush") // will be uppercased automatically
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetConfigName("config")   // name of config file (without extension)
	viper.AddConfigPath("./config") // optionally look for config in the working directory

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
