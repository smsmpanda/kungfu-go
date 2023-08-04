package orm

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"myexample/basic/orm/model"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func getDb() (*gorm.DB, error) {
	dsn := "root:password@tcp(192.168.200.20:3306)/door?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}

func Gorun() {

	db, err := getDb()
	if err != nil {
		return
	}
	//create(db)
	alarm := &model.Alarm{ID: 1}
	db.First(alarm)

	o, err := json.Marshal(alarm)
	fmt.Println(string(o))
}

func create(db *gorm.DB) {
	//产生一个随机数
	defer func() {
		if err := recover(); err != nil {
			log.Fatalf("insert db occurred error detail: %v", err)
		}
	}()

	gid := rand.Intn(10)

	alarm := model.Alarm{ID: 0, Personid: gid, Name: "测试名字", Depid: gid, Depname: "测试部门", Times: time.Now(), Type: 1, Mark: "这是使用gorm方式写入数据"}
	result := db.Create(&alarm)
	fmt.Println("this Alarm id's is :", alarm.ID)
	fmt.Println("this effect row's number is ", result.RowsAffected)
}

func query(db *gorm.DB) interface{} {
	return db.First(&model.Alarm{})
}
