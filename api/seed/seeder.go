// simple seed package version
package seed

import (
	"github.com/handl3r/GoForum/api/models"
	"github.com/jinzhu/gorm"
	"log"
)

var users = []models.User{
	models.User{
		Username: "test1",
		Email:    "test1@gmail.com",
		Password: "password",
	},
	models.User{
		Username: "test2",
		Email:    "test2@gmail.com",
		Password: "password",
	},
}
var posts = []models.Post{
	models.Post{
		Title:   "title1",
		Content: "content1",
	},
	models.Post{
		Title:   "title2",
		Content: "content2",
	},
}

func Load(db *gorm.DB) {
	var err error
	db = db.DropTableIfExists(&models.User{}, &models.Post{}, &models.Comment{}, &models.Like{})
	if err = db.Error; err != nil {
		log.Fatalf("error when delete tables: %v", err)
	}

	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("can not auto migrate database: %v", err)
	}

	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("can not add foreign key to post table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("can not seed user: %v", err)
		}
		posts[i].AuthorID = users[i].ID
		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("can not seed post: %v", err)
		}
	}
}
