package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	Mysql "spo2_server/helper/Mysql"
	"spo2_server/model"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

var dbTag *sqlx.DB

func init() {
	Mysql.CreateDbConn("mysql", "root@tcp(127.0.0.1:3306)/Spo2_Tag1")

}
func List(c *gin.Context) {

	id := c.Param("deviceId")
	c.JSON(http.StatusOK, Mysql.GetSpo2Data(id))

}
func ListId(c *gin.Context) {
	c.JSON(http.StatusOK, Mysql.GetDeviceIDList())
}
func InsertTag(c *gin.Context) {
	var insertObject model.InsertTag
	body, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(body, &insertObject)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	Mysql.InsertTag(insertObject)

	c.JSON(http.StatusOK, Mysql.GetDeviceIDList())

}
