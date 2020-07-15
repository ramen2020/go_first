package db

import (
	"fmt"
	config "tasksApp/config"
	entity "tasksApp/models/entity"

	"github.com/jinzhu/gorm"
)

// ConnectDb DB接続
func ConnectDb() *gorm.DB {
	DBMS := config.Config.DBMS
	USER := config.Config.USER
	PASS := config.Config.PASS
	PROTOCOL := config.Config.PROTOCOL
	DBNAME := config.Config.DBNAME
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.LogMode(true)
	db.SingularTable(true)
	db.AutoMigrate(&entity.Task{})

	fmt.Println("db connected: ", &db)
	return db
}

// GetAllTasks 全件取得
func GetAllTasks() []entity.Task {
	tasks := []entity.Task{}
	ConnectDb().Order("ID asc").Find(&tasks)
	defer ConnectDb().Close()

	return tasks
}

// GetTaskByID １件取得
func GetTaskByID(taskID int) []entity.Task {
	task := []entity.Task{}
	ConnectDb().First(&task, taskID)
	defer ConnectDb().Close()

	return task
}

// InsertTask 新規登録
func InsertTask(registerTask *entity.Task) {
	fmt.Printf("%+v", registerTask)
	ConnectDb().Create(&registerTask)
	defer ConnectDb().Close()
}

// UpdateStateTask 状態の更新
func UpdateStateTask(taskID int, taskState int) {
	task := []entity.Task{}
	ConnectDb().Model(&task).Where("ID = ?", taskID).Update("State", taskState)
	defer ConnectDb().Close()
}

// DeleteTask 削除
func DeleteTask(taskID int) {
	task := []entity.Task{}
	ConnectDb().Delete(&task, taskID)
	defer ConnectDb().Close()
}
