package controller

import (
	"api/services"
	Mongo "ble_mqtt/helper/Mongo"
	"io"
	"io/ioutil"
	"net/http"

	_mongo "go.mongodb.org/mongo-driver/mongo"
)

var db *_mongo.Collection

func init() {
	db = Mongo.Connect_Mongo("")
}

func List(w http.ResponseWriter, r *http.Request) {
	_, err := ioutil.ReadAll(io.LimitReader(r.Body, 1024)) //io.LimitReader限制大小
	if err != nil {
	}

	services.ResponseWithJson(w, http.StatusOK, Mongo.FindOneByApi(db, "d3:76:ce:a8:16:d2"))

}

func ListId(w http.ResponseWriter, r *http.Request) {
	_, err := ioutil.ReadAll(io.LimitReader(r.Body, 1024)) //io.LimitReader限制大小
	if err != nil {
	}

	services.ResponseWithJson(w, http.StatusOK, Mongo.FindOnlyId(db, "d3:76:ce:a8:16:d2"))

}
