package main

import (
	M "ble_mqtt/helper/MQ"
	routes "ble_mqtt/routers"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/handlers"

	"github.com/spf13/viper"
)

func init() {
	//log.SetLevel(log.DebugLevel)
}

func main() {
	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	M.MqttClientInit("", "", "")
	// <-c

	headers := handlers.AllowedHeaders([]string{"X-Request-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "PUT", "POST", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	router := routes.NewRouter()
	http.ListenAndServe(":8087", handlers.CORS(headers, methods, origins)(router))
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
	connectString := "sqlserver://" + viper.GetString("user") + ":" + viper.GetString("password") +
		"@" + viper.GetString("host") + ":" + viper.GetString("port") + "?database=" + viper.GetString("db") + "&connection+timeout=30"

	viper.Set("connectString", connectString) // same result as next line

}
