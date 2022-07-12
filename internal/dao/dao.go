package dao

import (
	"Aoi/internal/model"
	"Aoi/pkg/app"
	"gorm.io/gorm"
)

type Dao struct {
	db *gorm.DB
}

func NewDB(db *gorm.DB) *Dao {
	return &Dao{db: db}
}

func (d *Dao) CountTag(name string, state uint8) (int, error) {
	tag := model.Tag{
		Name:  name,
		State: state,
	}
	return tag.Count(d.db)

}

func (d *Dao) ListTag(state uint8, page, size int) ([]*model.Tag, error) {
	tag := model.Tag{
		State: state,
	}
	offset := app.GetPageOffset(page, size)
	return tag.List(d.db, offset, size)
}

func (d *Dao) CreateTag(name, createdBy string, state uint8) error {
	return model.Tag{
		Model: gorm.Model{},
		Common: model.Common{
			CreatedBy: createdBy,
		},
		Name:  name,
		State: state,
	}.Create(d.db)
}

func (d *Dao) UpdateTag(id uint, name, updateBy string, state uint8) error {
	return model.Tag{
		Model: gorm.Model{
			ID: id,
		},
		Common: model.Common{
			UpdatedBy: updateBy,
		},
		Name:  name,
		State: state,
	}.Update(d.db)

}
func (d *Dao) DeleteTag(id uint) error {
	return model.Tag{
		Model: gorm.Model{ID: id},
	}.Delete(d.db)
}
