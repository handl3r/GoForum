package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"html"
	"strings"
	"time"
)

type Post struct {
	ID        uint64    `json:"id" gorm:"primary_key;auto_increment"`
	Title     string    `json:"title" gorm:"size:255; not null; unique"`
	Content   string    `json:"content" gorm:"text; not null"`
	Author    User      `json:"author"`
	AuthorID  uint32    `json:"author_id" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	CreatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func (post *Post) Prepare() {
	post.Title = html.EscapeString(strings.TrimSpace(post.Title))
	post.Content = html.EscapeString(strings.TrimSpace(post.Content))
	post.Author = User{}
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()
}

func (post *Post) Validate() map[string]string {
	errorMessages := map[string]string{}
	if post.Title == "" {
		errorMessages["require_title"] = "require title"
	}
	if post.Content == "" {
		errorMessages["require_content"] = "require content"
	}
	if post.AuthorID < 1 {
		errorMessages["require_author"] = "require author"
	}
}

func (post *Post) SavePost(db *gorm.DB) (*Post, error) {
	var err error
	err = db.Debug().Model(&Post{}).Create(&post).Error
	if err != nil {
		return &Post{}, err
	}

	if post.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id=?", post.AuthorID).Take(&post.Author).Error
		if err != nil {
			return &Post{}, err
		}
	}
	return post, nil
}

func (post *Post) FindAllPosts(db *gorm.DB) (*[]Post, error) {
	var err error
	var posts []Post

	err = db.Debug().Model(&post).Limit(100).Order("created-at desc").Find(&posts).Error
	if err != nil {
		return &[]Post{}, err
	}

	if len(posts) > 0 {
		for i, _ := range posts {
			err = db.Debug().Model(&User{}).Where("id=?", posts[i].AuthorID).Take(&posts[i].Author).Error
			if err != nil {
				return &[]Post{}, err
			}
		}
	}
	return &posts, nil
}

func (post *Post) FindPostByID(db *gorm.DB, pid uint64) (*Post, error) {
	var err error

	err = db.Debug().Model(&Post{}).Where("id=?", pid).Take(&post).Error
	if err != nil {
		return &Post{}, errors.New("user not found")
	}

	err = db.Debug().Model(&User{}).Where("id=?", post.AuthorID).Take(&post.Author).Error
	if err != nil {
		return &Post{}, err
	}
	return post, nil
}

func (post *Post) UpdatePost(db *gorm.DB) (*Post, error) {
	var err error

	db = db.Debug().Model(&Post{}).Where("id=?", post.ID).UpdateColumns(
		map[string]interface{}{
			"title":      post.Title,
			"content":    post.Content,
			"updated_at": time.Now(),
		})
	if err = db.Error; err != nil {
		return &Post{}, err
	}

	err = db.Debug().Model(&User{}).Where("id=?", post.AuthorID).Take(&post.Author).Error
	if err != nil {
		return &Post{}, err
	}
	return post, nil
}

func (post *Post) DeletePost(db *gorm.DB) (int64, error) {
	var err error

	err = db.Debug().Model(&Post{}).Delete(&post).Error
	if err != nil {
		return 0, err
	}
	return db.RowsAffected, nil
}

func (post *Post) FindPostsOfUser(db *gorm.DB, uid uint32) (*[]Post, error) {
	var err error
	var user User
	var posts []Post

	err = db.Debug().Model(&User{}).Where("id=?", uid).Take(&user).Error
	if err != nil {
		return &[]Post{}, errors.New("user not found")
	}

	err = db.Debug().Model(&Post{}).Where("author_id=?", uid).Limit(100).
		Order("created_at desc").Take(&posts).Error
	if err != nil {
		return &[]Post{}, err
	}

	if len(posts) > 0 {
		for i, _ := range posts {
			posts[i].Author = user
		}
	}

	return &posts, nil
}

// Delete all posts of deleted user
func (post *Post) DeletePostsByUser(db *gorm.DB, uid uint32) (int64, error) {
	var err error
	var posts []Post

	err = db.Debug().Model(&Post{}).Where("author_id=?", uid).Find(&posts).Delete(&posts).Error
	if err != nil {
		return 0, err
	}
	return db.RowsAffected, nil
}
