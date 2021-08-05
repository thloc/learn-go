package model

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type Task struct {
	Id    int    `json:"id,omitempty" gorm:"column:id;"`
	Title string `json:"title" gorm:"column:title;"`
	Desc  string `json:"desc" gorm:"column:description;"`
	// Priority int    `json:"priority" gorm:"column: priority;"`
	// Status   int    `json:"status" gorm:"column: status;"`
}

func (Task) TableName() string {
	return "tasks"
}

func (res *Task) Validate() error {
	res.Title = strings.TrimSpace(res.Title)

	if len(res.Title) == 0 {
		return errors.New("Task title can't be blank")
	}

	return nil
}

func (task *Task) GetAll(db *gorm.DB) (interface{}, error) {
	tasks := []Task{}
	result := db.Table(task.TableName()).Where("status = ?", 1).Find(&tasks)

	if result.Error != nil {
		return nil, result.Error
	}

	return &tasks, nil
}

func (task *Task) CreateTask(db *gorm.DB, data *Task) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := db.Create(data).Error; err != nil {
		return err
	}

	return nil
}

func (task *Task) GetTaskByCondition(db *gorm.DB, conditions map[string]interface{}) (*Task, error) {
	var data Task

	if err := db.Where(conditions).
		First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
