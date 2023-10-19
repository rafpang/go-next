package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	if base.ID == uuid.Nil {
		base.ID = uuid.New()
	}
	return
}

type Project struct {
	Base
	ProjectID   uint
	ProjectName string
	Tasks       []Task
	Users       []User `gorm:"many2many:user_projects;"`
}

type Task struct {
	Base
	TaskID    uint
	TaskName  string
	DueDate   time.Time
	Status    string
	ProjectID uint
}

type User struct {
	Base
	Username       string
	HashedPassword string
	Projects       []Project `gorm:"many2many:user_projects;"`
}
