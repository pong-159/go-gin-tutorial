package repository

import (


	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"go-gin/tutorial-api/entity"
)

type VideoRepository interface {
	Save(entity.Video) 
	Update(entity.Video) 
	Delete(entity.Video)
	FindAll() []entity.Video
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewVideoRepository() VideoRepository {
	dsn := "host=localhost user=postgres password=admin dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("cannot connect to db  ")
	}
	
	db.AutoMigrate(&entity.Video{}, &entity.Person{})

	return &database{
		connection: db,
	}


	
}

func (db *database) CloseDB() {
	DB, err := db.connection.DB()
	if err != nil{
		panic("cannot close db  ")
	}
	errcls := DB.Close()
	if errcls != nil{
		panic("cannot close db  ")
	}
}

func (db *database) Save(video entity.Video) {
	db.connection.Create(&video)
}

func (db *database) Update(video entity.Video) {
	db.connection.Save(&video)
}

func (db *database) Delete(video entity.Video) {
	db.connection.Delete(&video)
}

func (db *database) FindAll() []entity.Video {
	var videos []entity.Video
	db.connection.Preload("Author").Find(&videos)
	return videos
}
