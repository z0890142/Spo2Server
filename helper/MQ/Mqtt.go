package MQTT

import (
	"encoding/json"
	"fmt"
	db "spo2_server/helper/Mongo"
	"spo2_server/model"
	"time"

	_mongo "go.mongodb.org/mongo-driver/mongo"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/viper"
)

var mongodb *_mongo.Collection

var f paho.MessageHandler = func(client paho.Client, msg paho.Message) {
	go messageHandler(&client, msg)

}

func MqttClientInit(address string, clientID string, apiKey string) {
	fmt.Println(viper.GetString("broker"))
	opts := paho.NewClientOptions().AddBroker(viper.GetString("broker")) //"tcp://localhost:1883"
	opts.SetClientID(clientID)
	opts.SetDefaultPublishHandler(f)
	opts.SetAutoReconnect(true)                                  //自动链接？！
	opts.SetMaxReconnectInterval(time.Duration(1) * time.Second) //自动链接间隔吧？！
	// opts.SetUsername(apiKey)
	// opts.SetPassword(apiKey)

	var lostf mqtt.ConnectionLostHandler = func(c mqtt.Client, err_ error) { //链接断开后的事件
		fmt.Println("mqtt disconnect")
	}
	opts.SetConnectionLostHandler(lostf)

	client := paho.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
	}
	mongodb = db.Connect_Mongo("123")
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

	db.MsgHandler(mongodb, mqResponse.DeviceId, mqResponse.Spo2, mqResponse.Bpm)

}
