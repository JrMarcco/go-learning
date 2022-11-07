package models

import "github.com/jinzhu/gorm"

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTags(pageNo, pageSize int, maps any) ([]*Tag, error) {
	var tags []*Tag
	var err error

	if pageSize > 0 && pageNo > 0 {
		err = db.Where(maps).Find(&tags).Offset(pageNo).Limit(pageSize).Error
	} else {
		err = db.Where(maps).Find(&tags).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return tags, nil
}

func GetTagTotal(maps any) (int, error) {
	var cnt int
	if err := db.Model(&Tag{}).Where(maps).Count(&cnt).Error; err != nil {
		return 0, err
	}
	return cnt, nil
}

func ExistTagByName(name string) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("name = ?", name).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	return tag.ID > 0, nil
}

func AddTag(name string, state int, createBy string) error {
	tag := Tag{
		Name:      name,
		State:     state,
		CreatedBy: createBy,
	}

	if err := db.Create(&tag).Error; err != nil {
		return err
	}
	return nil
}

func ExistTagByID(id int) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("id = ?", id).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	return tag.ID > 0, nil
}

func EditTag(id int, data any) error {
	if err := db.Model(&Tag{}).Where("id = ?", id).Update(data).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTag(id int) error {
	if err := db.Where("id = ?", id).Delete(&Tag{}).Error; err != nil {
		return err
	}
	return nil
}

func CleanAllTag() (bool, error) {
	if err := db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Tag{}).Error; err != nil {
		return false, err
	}
	return true, nil
}
