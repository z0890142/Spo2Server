package MQTT

import (
	"encoding/json"
	"fmt"
	"spo2_server/helper/Mysql"
	"spo2_server/model"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/viper"
)

var DeviceMap map[string]string
var f paho.MessageHandler = func(client paho.Client, msg paho.Message) {
	go messageHandler(&client, msg)

}

func MqttClientInit(address string, clientID string, apiKey string) {
	DeviceMap = make(map[string]string)
	fmt.Println(viper.GetString("broker"))
	opts := paho.NewClientOptions().AddBroker(viper.GetString("broker")) //"tcp://localhost:1883"
	opts.SetClientID(clientID)
	opts.SetDefaultPublishHandler(f)
	opts.SetAutoReconnect(true)
	opts.SetMaxReconnectInterval(time.Duration(1) * time.Second)

	var lostf mqtt.ConnectionLostHandler = func(c mqtt.Client, err_ error) {
		fmt.Println("mqtt disconnect")
	}
	opts.SetConnectionLostHandler(lostf)

	client := paho.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	} else {
	}
	MqttSubscribe(client, "spo2")

}

func MqttSubscribe(client paho.Client, topic string) {
	if token := client.Subscribe(topic, 0, f); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}
	fmt.Println("Subscribe " + topic)

}

func MqttPublish(client paho.Client, topic string, msg string) {

	token := client.Publish(topic, 0, false, "jsonStr")
	token.Wait()
}

func messageHandler(clientP *paho.Client, msg paho.Message) {
	fmt.Printf("%s", msg.Payload())
	var mqResponse model.MqResponse

	json.Unmarshal(msg.Payload(), &mqResponse)
	if DeviceMap[mqResponse.DeviceId] == "" {
		Mysql.InsertDevice(mqResponse.DeviceId)
	}
	Mysql.InsertSpo2FromDevice(mqResponse)

}
