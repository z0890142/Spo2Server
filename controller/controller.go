package controller

import (
	"net/http"
	Mongo "spo2_server/helper/Mongo"

	"github.com/gin-gonic/gin"
	_mongo "go.mongodb.org/mongo-driver/mongo"
)

var db *_mongo.Collection

func InitController() {
	db = Mongo.Connect_Mongo("")
}

// func List(w http.ResponseWriter, r *http.Request) {
// 	_, err := ioutil.ReadAll(io.LimitReader(r.Body, 1024)) //io.LimitReader限制大小
// 	if err != nil {
// 	}
// 	vars := mux.Vars(r)
// 	id := vars["deviceId"]
// 	services.ResponseWithJson(w, http.StatusOK, Mongo.FindOneByApi(db, id))

// }

// func ListId(w http.ResponseWriter, r *http.Request) {
// 	_, err := ioutil.ReadAll(io.LimitReader(r.Body, 1024)) //io.LimitReader限制大小
// 	if err != nil {
// 	}

// 	services.ResponseWithJson(w, http.StatusOK, Mongo.FindOnlyId(db))

// }

func List(c *gin.Context) {

	id := c.Param("deviceId")
	c.JSON(http.StatusOK, Mongo.FindOneByApi(db, id))
	// services.ResponseWithJson(w, http.StatusOK, Mongo.FindOneByApi(db, id))
}
func ListId(c *gin.Context) {

	c.JSON(http.StatusOK, Mongo.FindOnlyId(db))

	// services.ResponseWithJson(w, http.StatusOK, Mongo.FindOnlyId(db))

}
