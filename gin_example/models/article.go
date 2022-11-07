package models

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title         string `json:"title"`
	CoverImageUrl string `json:"cover_image_url"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         int    `json:"state"`
}

func ExistArticleByID(id int) (bool, error) {
	var article Article
	err := db.Select("id").Where("id = ?", id).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	return article.ID > 0, nil
}

func GetArticleTotal(maps any) (int, error) {
	var cnt int
	if err := db.Model(&Article{}).Where(maps).Count(&cnt).Error; err != nil {
		return 0, err
	}
	return cnt, nil
}

func GetArticles(pageNo int, pageSize int, maps any) ([]*Article, error) {
	var articles []*Article
	var err error

	if pageSize > 0 && pageNo > 0 {
		err = db.Preloads("Tag").Where(maps).Offset(pageNo).Limit(pageSize).Find(&articles).Error
	} else {
		err = db.Preloads("Tag").Where(maps).Find(&articles).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return articles, nil
}

func GetArticle(id int) (*Article, error) {
	var article Article

	err := db.Where("id = ?", id).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	err = db.Model(&article).Related(&article.Tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &article, nil
}

func AddArticle(data map[string]any) error {

	article := Article{
		TagID:         data["tag_id"].(int),
		Title:         data["title"].(string),
		CoverImageUrl: data["cover_image_url"].(string),
		Desc:          data["desc"].(string),
		Content:       data["content"].(string),
		CreatedBy:     data["created_by"].(string),
		State:         data["state"].(int),
	}

	if err := db.Create(&article).Error; err != nil {
		return err
	}
	return nil
}

func EditArticle(id int, data any) error {
	if err := db.Model(&Article{}).Where("id = ?", id).Update(data).Error; err != nil {
		return err
	}
	return nil
}

func DeleteArticle(id int) error {
	if err := db.Where("id = ?", id).Delete(&Article{}).Error; err != nil {
		return err
	}
	return nil
}
