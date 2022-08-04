package entity

import "time"

type Person struct {
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string `json:"firstname" gorm:"type:varchar(32)"`
	LastName  string `json:"lastname" gorm:"type:varchar(32)"`
	Age       uint8  `json:"age" binding:"gte=1,lte=150" `
	Email     string `json:"email" validate:"required,email" gorm:"type:varchar(256)"`
}

type Video struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title       string    `json:"title" binding:"min=2,max=100" gorm:"type:varchar(20)"`
	Description string    `json:"description" binding:"max=200" gorm:"type:varchar(200)"`
	Url         string    `json:"url" binding:"required,url" gorm:"type:varchar(256);UNIQUE"`
	Author      Person    `json:"author" binding:"required" gorm:"foreignKey:PersonID"`
	PersonID    uint64    `json:"-"`
	CreatedAt   time.Time `json:"-" gorm:"CURRENT_TIMESTAMP" jsonL:"created_at"`

	UpdatedAt   time.Time `json:"-" gorm:"CURRENT_TIMESTAMP" jsonL:"updated_at"`
}
