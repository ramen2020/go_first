package main

import (
	"log"
	"net/http"
	controller "tasksApp/controllers/controller"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	serve()
}

func serve() {
	router := gin.Default()
	router.Static("/views", "./views")
	router.StaticFS("/tasks", http.Dir("./views/static"))

	// ルーティング API
	router.GET("/allTasks/", controller.GetAllTasks)            // 全件取得
	router.GET("/task", controller.GetTaskByID)                 // 一件取得
	router.POST("/createTask", controller.CreateTask)           // 新規作成
	router.POST("/changeStateTask", controller.ChangeStateTask) // 状態変更
	router.POST("/deleteTask", controller.DeleteTask)           // 削除

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}
