package models

import (
	"github.com/set2002satoshi/my-site-api/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api/pkg/module/customs/types"
)

type BlogAndCategoryEntity struct {
	Id         types.IDENTIFICATION `gorm:"primaryKey"`
	BlogId     types.IDENTIFICATION
	CategoryId types.IDENTIFICATION
}

func NewBlogAndCategoryEntity(
	id,
	blogId,
	categoryId int,
) (*BlogAndCategoryEntity, error) {
	BaC := new(BlogAndCategoryEntity)
	var err error
	err = errors.Combine(err, BaC.setId(id))
	err = errors.Combine(err, BaC.setBlogId(blogId))
	err = errors.Combine(err, BaC.setCategoryId(categoryId))
	if err != nil {
		return new(BlogAndCategoryEntity), err
	}
	return BaC, nil
}

func (bac *BlogAndCategoryEntity) GetId() types.IDENTIFICATION {
	return bac.Id
}

func (bac *BlogAndCategoryEntity) GetBlogId() types.IDENTIFICATION {
	return bac.BlogId
}

func (bac *BlogAndCategoryEntity) GetCategoryId() types.IDENTIFICATION {
	return bac.CategoryId
}

func (bac *BlogAndCategoryEntity) setId(id int) error {
	i, err := types.NewIDENTIFICATION(id)
	if err != nil {
		return err
	}
	bac.Id = i
	return nil
}

func (bac *BlogAndCategoryEntity) setBlogId(blogId int) error {
	i, err := types.NewIDENTIFICATION(blogId)
	if err != nil {
		return err
	}
	bac.BlogId = i
	return nil
}

func (bac *BlogAndCategoryEntity) setCategoryId(categoryId int) error {
	i, err := types.NewIDENTIFICATION(categoryId)
	if err != nil {
		return err
	}
	bac.CategoryId = i
	return nil
}
