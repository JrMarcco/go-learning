package cache_service

import (
	"go_learning/gin_example/pkg/e"
	"strconv"
	"strings"
)

type Tag struct {
	Id    int
	Name  string
	State int

	PageNo   int
	PageSize int
}

func (t *Tag) GetTagsKey() string {
	keys := []string{
		e.CACHE_TAG,
		"LIST",
	}

	if t.Name != "" {
		keys = append(keys, t.Name)
	}
	if t.State >= 0 {
		keys = append(keys, strconv.Itoa(t.State))
	}
	if t.PageNo >= 0 {
		keys = append(keys, strconv.Itoa(t.PageNo))
	}
	if t.PageSize >= 0 {
		keys = append(keys, strconv.Itoa(t.PageSize))
	}

	return strings.Join(keys, ":")
}
