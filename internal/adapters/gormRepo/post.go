package gormRepo

import (
	"context"
	"errors"
	"go_framework/internal/app/respositories"
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

func (r PostRepo) FindByNo(ctx context.Context, no string) (respositories.Post, error) {
	var post PostModel
	return r.marshallPost(&post), nil
}

func (r PostRepo) Create(ctx context.Context, post *respositories.Post) error {
	data := r.unmarshallPost(post)
	return r.DB.Create(&data).Error
}

func (r PostRepo) marshallPost(postModel *PostModel) respositories.Post {
	return respositories.Post{
		ID:        postModel.ID,
		No:        postModel.No,
		UserNo:    postModel.UserNo,
		Status:    postModel.Status,
		Title:     postModel.Title,
		Content:   postModel.Content,
		ShowAt:    time.Time(postModel.ShowAt),
		CreatedAt: postModel.CreatedAt,
		UpdatedAt: postModel.UpdatedAt,
	}
}

func (r PostRepo) unmarshallPost(post *respositories.Post) PostModel {
	return PostModel{
		ID:        post.ID,
		No:        post.No,
		UserNo:    post.UserNo,
		Status:    post.Status,
		Title:     post.Title,
		Content:   post.Content,
		ShowAt:    datatypes.Date(post.ShowAt),
		UpdatedAt: post.UpdatedAt,
	}
}

func NewGormPostRepo(DB *gorm.DB) PostRepo {
	return PostRepo{
		DB: DB,
	}
}
