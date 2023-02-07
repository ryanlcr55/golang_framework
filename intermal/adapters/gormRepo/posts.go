package gormRepo

import (
	"errors"
	"go_framework/intermal/respositories"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type PostModel struct {
	ID        uint64         `gorm:"primaryKey;autoIncrement;"`
	No        string         `gorm:"type:varchar(50);not null;uniqueIndex"`
	UserNo    string         `gorm:"type:varchar(50);not null;index"`
	Status    string         `gorm:"type:enum('enable', 'disable');default:'enable';not null;"`
	Title     string         `gorm:"type:varchar(255);not null;"`
	Content   string         `gorm:"type:varchar(255);not null"`
	ShowAt    datatypes.Date `gorm:"type:date;not null;unique Index"`
	CreatedAt time.Time      `gorm:"type:datetime;autoCreateTime;"`
	UpdatedAt time.Time      `gorm:"type:datetime;autoCreateTime;"`
}

type PostRepo struct {
	DB *gorm.DB
}

func (r PostRepo) WithTransaction(tx any) (respositories.IPostRepo, error) {
	gormTx, ok := tx.(*gorm.DB)
	if !ok && gormTx != nil {
		return nil, errors.New("transaction handler is illegal")
	}
	return PostRepo{
		DB: gormTx,
	}, nil
}

func (r PostRepo) FindByNo(no string) (respositories.Post, error) {
	var post PostModel
	return r.marshallPost(post), nil
}

func (r PostRepo) Create(post respositories.Post) error {

	return nil
}

func (r PostRepo) marshallPost(post PostModel) respositories.Post {
	return respositories.Post{
		ID:        post.ID,
		No:        post.No,
		UserNo:    post.UserNo,
		Status:    post.Status,
		Title:     post.Title,
		Content:   post.Content,
		ShowAt:    time.Time(post.ShowAt),
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
}
