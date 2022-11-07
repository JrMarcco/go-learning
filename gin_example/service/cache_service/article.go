package cache_service

import (
	"go_learning/gin_example/pkg/e"
	"strconv"
	"strings"
)

type Article struct {
	ID    int
	TagId int
	State int

	PageNo   int
	PageSize int
}

func (a *Article) GetArticleKey() string {
	return e.CACHE_ARTICEL + ":" + strconv.Itoa(a.ID)
}

func (a *Article) GetArticlesKey() string {
	keys := []string{
		e.CACHE_ARTICEL,
		"LIST",
	}

	if a.ID > 0 {
		keys = append(keys, strconv.Itoa(a.ID))
	}
	if a.TagId > 0 {
		keys = append(keys, strconv.Itoa(a.TagId))
	}
	if a.PageNo > 0 {
		keys = append(keys, strconv.Itoa(a.PageNo))
	}
	if a.PageSize > 0 {
		keys = append(keys, strconv.Itoa(a.PageSize))
	}

	return strings.Join(keys, ":")
}
