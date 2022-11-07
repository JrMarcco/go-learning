package tag_service

import (
	"encoding/json"
	"go_learning/gin_example/models"
	"go_learning/gin_example/pkg/gredis"
	"go_learning/gin_example/pkg/logging"
	"go_learning/gin_example/service/cache_service"
)

type Tag struct {
	ID         int
	Name       string
	State      int
	CreatedBy  string
	ModifiedBy string

	PageNo   int
	PageSize int
}

func (t *Tag) ExistByName() (bool, error) {
	return models.ExistTagByName(t.Name)
}

func (t *Tag) ExistByID() (bool, error) {
	return models.ExistTagByID(t.ID)
}

func (t *Tag) Add() error {
	return models.AddTag(t.Name, t.State, t.CreatedBy)
}

func (t *Tag) Edit() error {
	data := make(map[string]any)
	data["modified_by"] = t.ModifiedBy
	data["name"] = t.Name

	if t.State >= 0 {
		data["state"] = t.State
	}

	return models.EditTag(t.ID, data)
}

func (t *Tag) Delete() error {
	return models.DeleteTag(t.ID)
}

func (t *Tag) GetAll() ([]*models.Tag, error) {
	var tags, cacheTags []*models.Tag

	cache := cache_service.Tag{
		State:    t.State,
		PageNo:   t.PageNo,
		PageSize: t.PageSize,
	}

	key := cache.GetTagsKey()
	if gredis.Exists(key) {
		if data, err := gredis.Get(key); err == nil {
			if err := json.Unmarshal(data, &cacheTags); err == nil {
				return cacheTags, nil
			}
		} else {
			logging.Info(err)
		}
	}

	var err error
	tags, err = models.GetTags(t.PageNo, t.PageSize, t.getMaps())
	if err != nil {
		return nil, err
	}

	if err = gredis.Set(key, tags, 3600); err != nil {
		logging.Info(err)
	}

	return tags, nil
}

func (t *Tag) getMaps() map[string]any {
	maps := make(map[string]any)
	maps["deleted_on"] = 0

	if t.Name != "" {
		maps["name"] = t.Name
	}

	if t.State >= 0 {
		maps["state"] = t.State
	}

	return maps
}
