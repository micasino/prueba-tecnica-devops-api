package domain

import (
	"context"
	"time"
)

type Devops struct {
	ID            int32     `gorm:"primary_key;column:id;autoIncrement;" json:"id"`
	Message       string    `gorm:"column:message;type:VARCHAR;size:100;" json:"message"`
	To            string    `gorm:"column:to;type:VARCHAR;size:100;" json:"to"`
	From          string    `gorm:"column:from;type:VARCHAR;size:200;" json:"from"`
	TimeToLifeSec int       `gorm:"column:timeToLifeSec;type:integer;" json:"timeToLifeSec"`
	DateCreate    time.Time `gorm:"column:date_create;autoCreateTime:true;" json:"date_create"`
	DateDeleted   time.Time `json:"date_deleted" copier:"must"`
}

// TableName sets the insert table name for this struct type
func (c *Devops) TableName() string {
	return "Devops"
}

type DevopsRepository interface {
	FindAll(ctx context.Context) ([]Devops, error)
	FindByID(ctx context.Context, id uint) (Devops, error)
	Save(ctx context.Context, devops Devops) (Devops, error)
	SaveAll(ctx context.Context, devops []Devops) (string, error)
	Update(ctx context.Context, devops Devops) (Devops, error)
	Delete(ctx context.Context, devops Devops) error
}
