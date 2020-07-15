package controller

import (
	strconv "strconv"
	db "tasksApp/models/db"
	entity "tasksApp/models/entity"

	"github.com/gin-gonic/gin"
)

// 状態
const (
	// 問題あり
	NotPurchased = 0
	// 解決済み
	Purchased = 1
)

// GetAllTasks 全件取得
func GetAllTasks(c *gin.Context) {
	resultTasks := db.GetAllTasks()

	c.JSON(200, resultTasks)
}

// GetTaskByID 特定のIDのレコードを取得
func GetTaskByID(c *gin.Context) {
	taskIDStr := c.Query("taskID")
	taskID, _ := strconv.Atoi(taskIDStr)
	resultTask := db.GetTaskByID(taskID)

	c.JSON(200, resultTask)
}

// CreateTask 新規登録
func CreateTask(c *gin.Context) {
	taskName := c.PostForm("taskName")
	taskMemo := c.PostForm("taskMemo")

	task := entity.Task{
		Name:  taskName,
		Memo:  taskMemo,
		State: NotPurchased,
	}

	db.InsertTask(&task)
}

// ChangeStateTask 状態の更新
func ChangeStateTask(c *gin.Context) {
	reqTaskID := c.PostForm("taskID")
	reqTaskState := c.PostForm("taskState")
	taskID, _ := strconv.Atoi(reqTaskID)
	taskState, _ := strconv.Atoi(reqTaskState)
	changeState := NotPurchased

	if taskState == NotPurchased {
		changeState = Purchased
	} else {
		changeState = NotPurchased
	}

	db.UpdateStateTask(taskID, changeState)
}

// DeleteTask 削除
func DeleteTask(c *gin.Context) {
	taskIDStr := c.PostForm("taskID")
	taskID, _ := strconv.Atoi(taskIDStr)

	db.DeleteTask(taskID)
}
