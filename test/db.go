package test

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

type Config struct {
	Host    string
	Port    uint16
	Usr     string
	Passwd  string
	DbName  string
	Charset string
}

var Db *sql.DB = nil
var DbErr error

func Open(config Config) {
	fmt.Printf("db config: host:%s,port:%d,usr:%s,passwd:%s,db name:%s\n", config.Host, config.Port, config.Usr, config.Passwd, config.DbName)
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		config.Usr, config.Passwd, config.Host, strconv.Itoa(int(config.Port)), config.DbName, config.Charset)
	Db, DbErr = sql.Open("mysql", dbDSN)
	if DbErr != nil {
		fmt.Println("mysql open fail error:" + DbErr.Error())
	}
	// 最大连接数
	Db.SetMaxOpenConns(100)
	// 闲置连接数
	Db.SetMaxIdleConns(20)
	// 最大连接周期
	Db.SetConnMaxLifetime(100 * time.Second)

	if DbErr = Db.Ping(); DbErr != nil {
		panic("数据库链接失败: " + DbErr.Error())
	}
}

type DeviceInfo struct {
	CameraIp      string `"camera_ip_addr"`//摄像头ip
	CameraFactory string `"camera_factory_string"`//摄像头品牌
	ScreenFactory string `"display_factory_string"`//显示屏品牌
}

func GetDeviceInfo() (DeviceInfo, error) {
	sqlCmd := "select camera_ip_addr,camera_factory_string,display_factory_string from device_info;"
	row :=Db.QueryRow(sqlCmd)
	var deviceInfo DeviceInfo
	var err error
	if err=row.Scan(&deviceInfo.CameraIp,&deviceInfo.CameraFactory,&deviceInfo.ScreenFactory);err!=nil {
		fmt.Printf("mysql:%s query fail\n",sqlCmd)
	}

	return deviceInfo,err
}
