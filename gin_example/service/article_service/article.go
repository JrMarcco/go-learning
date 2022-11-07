package article_service

import (
	"encoding/json"
	"go_learning/gin_example/models"
	"go_learning/gin_example/pkg/gredis"
	"go_learning/gin_example/pkg/logging"
	"go_learning/gin_example/service/cache_service"
)

type Article struct {
	ID            int
	TagID         int
	Title         string
	CoverImageUrl string
	Desc          string
	Content       string
	CreatedBy     string
	ModifiedBy    string
	State         int

	PageNo   int
	PageSize int
}

func (a *Article) Add() error {
	article := map[string]any{
		"tag_id":          a.TagID,
		"title":           a.Title,
		"cover_image_url": a.CoverImageUrl,
		"desc":            a.Desc,
		"content":         a.Content,
		"created_by":      a.CreatedBy,
		"state":           a.State,
	}

	if err := models.AddArticle(article); err != nil {
		return err
	}
	return nil
}

func (a *Article) Edit() error {
	return models.EditArticle(a.ID, map[string]any{
		"tag_id":          a.TagID,
		"title":           a.Title,
		"cover_image_url": a.CoverImageUrl,
		"desc":            a.Desc,
		"content":         a.Content,
		"modified_by":     a.ModifiedBy,
		"state":           a.State,
	})
}

func (a *Article) Get() (*models.Article, error) {
	var cacheArticle *models.Article

	cache := cache_service.Article{ID: a.ID}
	key := cache.GetArticleKey()
	if gredis.Exists(key) {
		if data, err := gredis.Get(key); err == nil {
			if err = json.Unmarshal(data, &cacheArticle); err == nil {
				return cacheArticle, nil
			}
		} else {
			logging.Info(err)
		}
	}

	article, err := models.GetArticle(a.ID)
	if err != nil {
		return nil, err
	}
	if err := gredis.Set(key, article, 3600); err != nil {
		logging.Info(err)
	}
	return article, nil
}

func (a *Article) GetAll() ([]*models.Article, error) {
	var articles, cacheArticles []*models.Article

	cache := cache_service.Article{
		TagId: a.TagID,
		State: a.State,

		PageNo:   a.PageNo,
		PageSize: a.PageSize,
	}

	key := cache.GetArticlesKey()
	if gredis.Exists(key) {
		if data, err := gredis.Get(key); err == nil {
			if err = json.Unmarshal(data, &cacheArticles); err == nil {
				return cacheArticles, nil
			}
		} else {
			logging.Info(err)
		}
	}

	var err error
	articles, err = models.GetArticles(a.PageNo, a.PageSize, a.getMaps())
	if err != nil {
		return nil, err
	}

	if err = gredis.Set(key, articles, 3600); err != nil {
		logging.Info(err)
	}

	return articles, nil
}

func (a *Article) Delete() error {
	return models.DeleteArticle(a.ID)
}

func (a *Article) ExistByID() (bool, error) {
	return models.ExistArticleByID(a.ID)
}

func (a *Article) Count() (int, error) {
	return models.GetArticleTotal(a.getMaps())
}

func (a *Article) getMaps() map[string]any {
	maps := make(map[string]any)
	maps["deleted_on"] = 0
	if a.State != -1 {
		maps["state"] = a.State
	}
	if a.TagID != -1 {
		maps["tag_id"] = a.TagID
	}
	return maps
}
