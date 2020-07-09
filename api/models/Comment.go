package models

import (
	"github.com/jinzhu/gorm"
	"html"
	"strings"
	"time"
)

type Comment struct {
	ID        uint64    `json:"id" gorm:"primary_key; auto_increment"`
	UserID    uint32    `json:"user_id" gorm:"not null"`
	PostID    uint64    `json:"post_id" gorm:"not null"`
	Author    User      `json:"author"`
	Body      string    `json:"body" gorm:"text; not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func (comment *Comment) Prepare() {
	comment.ID = 0
	comment.Author = User{}
	comment.Body = html.EscapeString(strings.TrimSpace(comment.Body))
	comment.UpdatedAt = time.Now()
	comment.CreatedAt = time.Now()
}

func (comment *Comment) Validate() map[string]string {
	errorMessages := make(map[string]string)
	if comment.Body == "" {
		errorMessages["require_body"] = "require body"
	}
	return errorMessages
}

func (comment *Comment) SaveComment(db *gorm.DB) (*Comment, error) {
	var err error
	err = db.Debug().Model(&Comment{}).Create(comment).Error
	if err != nil {
		return &Comment{}, err
	}

	if comment.ID != 0 {
		err = db.Model(&Comment{}).Where("id=?", comment.ID).Take(comment).Error
		if err != nil {
			return &Comment{}, err
		}
	}
	return comment, nil
}

func (comment *Comment) GetCommentsByPost(db *gorm.DB, pid uint64) (*[]Comment, error) {
	var err error
	var comments []Comment

	db = db.Debug().Model(&Comment{}).Where("post_id=?", pid).Order("created_at desc").Find(&comments)
	if db.Error != nil {
		return &[]Comment{}, err
	}
	if len(comments) > 0 {
		for i, _ := range comments {
			err = db.Debug().Model(&User{}).Where("id=?", comments[i].UserID).Take(&comments[i].Author).Error
			if err != nil {
				return &[]Comment{}, err
			}
		}
	}
	return &comments, nil

}

func (comment *Comment) UpdateComment(db *gorm.DB) (*Comment, error) {
	var err error

	db = db.Debug().Model(&Comment{}).Where("id=?", comment.ID).UpdateColumns(
		map[string]interface{}{
			"body":       comment.Body,
			"updated_at": time.Now(),
		})
	if err = db.Error; err != nil {
		return &Comment{}, err
	}
	err = db.Debug().Model(&Comment{}).Where("id=?", comment.ID).Take(comment).Error
	if err != nil {
		return &Comment{}, err
	}
	return comment, nil
}

func (comment *Comment) DeleteComment(db *gorm.DB, cid uint64) (int64, error) {
	var err error

	err = db.Debug().Model(&Comment{}).Where("id=?", cid).Take(&Comment{}).Delete(&Comment{}).Error
	if err != nil {
		return 0, err
	}
	return db.RowsAffected, nil
}

func (comment *Comment) DeleteCommentsByPost(db *gorm.DB, pid uint64) (int64, error) {
	var err error
	var comments []Comment

	db = db.Debug().Model(&Comment{}).Where("post_id=?", pid).Find(&comments).Delete(&comments)
	if db.Error != nil {
		return 0, err
	}
	return db.RowsAffected, nil
}

func (comment *Comment) DeleteCommentsByUser(db *gorm.DB, uid uint32) (int64, error) {
	var err error
	var comments []Comment

	db = db.Debug().Model(&Comment{}).Where("user_id=?", uid).Find(&comments).Delete(&comments)
	if err = db.Error; err != nil {
		return 0, err
	}
	return db.RowsAffected, nil
}
