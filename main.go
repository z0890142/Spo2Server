package main

import (
	"fmt"
	"strings"

	controller "spo2_server/controller"
	M "spo2_server/helper/MQ"

	"github.com/gin-gonic/gin"

	"github.com/spf13/viper"
)

func init() {
	//log.SetLevel(log.DebugLevel)
	InitConfig()
	controller.InitController()
}

func main() {
	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	M.MqttClientInit("", "", "")
	// // <-c

	// headers := handlers.AllowedHeaders([]string{"X-Request-With", "Content-Type", "Authorization"})
	// methods := handlers.AllowedMethods([]string{"GET", "PUT", "POST", "DELETE"})
	// origins := handlers.AllowedOrigins([]string{"*"})
	// router := routes.NewRouter()
	// http.ListenAndServe(":8087", handlers.CORS(headers, methods, origins)(router))
	r := gin.Default()
	r.LoadHTMLGlob("./views/*.html")          // complains about /static being a dir on boot
	r.LoadHTMLFiles("./views/static/*/*")     //  load the static path
	r.Static("/static", "./views/static")     // use the loaded source
	r.StaticFile("/spo2", "views/index.html") // use the loaded source

	r.GET("/db/:deviceId", controller.List)
	r.GET("/id/list", controller.ListId)

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
