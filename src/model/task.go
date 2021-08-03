package model

import (
	"gorm.io/gorm"
)

type Task struct {
	Id    int    `json:"id,omitempty" gorm:"column:id;"`
	Tilte string `json:"title" gorm:"column:title;"`
	Desc  string `json:"desc" gorm:"column:description;"`
	// Priority int    `json:"priority" gorm:"column: priority;"`
	// Status   int    `json:"status" gorm:"column: status;"`
}

func (Task) TableName() string {
	return "tasks"
}

func (task *Task) GetAll(db *gorm.DB) (interface{}, error) {
	tasks := []Task{}
	result := db.Table(task.TableName()).Where("status = ?", 1).Find(&tasks)

	if result.Error != nil {
		return nil, result.Error
	}

	return &tasks, nil
}
