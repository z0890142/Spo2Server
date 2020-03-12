package Mysql

import (
	"fmt"
	"spo2_server/model"
	"time"

	_ "github.com/go-sql-driver/mysql" //前面加 _ 是為了只讓他執行init
	"github.com/jmoiron/sqlx"
)

type Spo2Data struct {
	Spo2 int    `json:"spo2"`
	Bpm  int    `json:"bpm"`
	Time string `json:"time"`
}

var db *sqlx.DB

func CreateDbConn(driveName string, dataSourceName string) (*sqlx.DB, error) {
	db, err := sqlx.Open(driveName, dataSourceName)
	db.SetConnMaxLifetime(100)
	db.SetMaxIdleConns(10)

	if err := db.Ping(); err != nil {

		return db, err
	}

	return db, err
}

func InsertTag(insertObject model.InsertTag) bool {
	startTime := time.Unix(insertObject.Start/1000, 0).Format("2006-01-02 15:04:05")
	endTime := time.Unix(insertObject.End/1000, 0).Format("2006-01-02 15:04:05")
	stmt, err := db.Prepare("update Spo2_Tag1 set Tag=? where Time Between ? and ?)")
	if err != nil {
		fmt.Println("InsertTag error" + err.Error())
		return false
	}
	_, err = stmt.Exec(1, startTime, endTime)
	if err != nil {
		fmt.Println("InsertTag error" + err.Error())
		return false
	}
	return true
}

func InsertDevice(deviceID string) bool {
	stmt, err := db.Prepare("insert into Device(DeviceID) values (?)")
	if err != nil {
		fmt.Println("InsertDevice error" + err.Error())
		return false
	}
	_, err = stmt.Exec(deviceID)
	if err != nil {
		fmt.Println("InsertDevice error" + err.Error())
		return false
	}
	return true
}

func GetDeviceIDList() []string {
	var DeviceIDList []string
	rows, err := db.Query("select DeviceID from Device")
	if err != nil {
		fmt.Println("GetDeviceIDList Error : " + err.Error())
	}
	for rows.Next() {
		var DeviceID string
		rows.Scan(&DeviceID)
		DeviceIDList = append(DeviceIDList, DeviceID)
	}
	return DeviceIDList
}

func GetSpo2Data(deviceId string) []Spo2Data {
	var Spo2DataList []Spo2Data
	rows, err := db.Query("select Spo2,Bpm,Time from Spo2 where DeviceID=?", deviceId)
	if err != nil {
		fmt.Println("GetSpo2Data Error : " + err.Error())
	}
	for rows.Next() {
		var Spo2Data Spo2Data
		// var timeStr string
		rows.Scan(&Spo2Data.Spo2, &Spo2Data.Bpm, &Spo2Data.Time)
		// tranTime, _ := time.Parse("2006-01-02 15:04:05", timeStr)
		// timeInt, _ := strconv.ParseInt(tranTime.Format("2006-01-02 15:04:05"), 10, 64)
		// Spo2Data.Time = timeInt * 1000
		Spo2DataList = append(Spo2DataList, Spo2Data)
	}
	return Spo2DataList
}

func InsertSpo2FromDevice(data model.MqResponse) {
	now := time.Now()
	time_UTC := now.In(time.FixedZone("CST", 28800))
	time := time_UTC.Format("2006-01-02 15:04:05")
	stmt, err := db.Prepare("insert into Spo2(DeviceID,Spo2,Bpm,Time) values(?,?,?,?)")
	_, err = stmt.Exec(data.DeviceId, data.Spo2, data.Bpm, time)
	if err != nil {
		panic(err)
	}

}
