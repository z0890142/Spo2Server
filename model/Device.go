package model

type Device struct {
	DeviceId string       `json:"deviceid"`
	First    string       `json:"first"`
	Last     string       `json:"last"`
	Data     []MqResponse `json:"data"`
}
