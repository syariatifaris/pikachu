package repo

import "github.com/syariatifaris/pikachu/app/module/page/model"

func NewPageRepository() PageRepository {
	return &pageRepository{}
}

type PageRepository interface {
	GetAll() []*model.Page
	Get(id int64) *model.Page
}

type pageRepository struct {
}

func (*pageRepository) GetAll() []*model.Page {
	return []*model.Page{}
}

func (*pageRepository) Get(id int64) *model.Page {
	return &model.Page{}
}
