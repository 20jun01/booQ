package model

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// Item itemの構造体
type Item struct {
	gorm.Model
	Name        string `gorm:"type:varchar(64);not null" json:"name"`
	Type        int    `gorm:"type:int;not null" json:"type"`
	Code        string `gorm:"type:varchar(13);" json:"code"`
	Description string `gorm:"type:text;" json:"description"`
	ImgURL      string `gorm:"type:text;" json:"img_url"`
}

// TableName dbのテーブル名を指定する
func (item *Item) TableName() string {
	return "items"
}

// GetItems 全itemを取得する
func GetItems() ([]Item, error) {
	res := []Item{}
	db.Find(&res)
	return res, nil
}

// CreateItem 新しいItemを登録する
func CreateItem(item Item) (Item, error) {
	if item.Name == "" {
		return Item{}, errors.New("Nameが存在しません")
	}
	db.Create(&item)
	return item, nil
}
