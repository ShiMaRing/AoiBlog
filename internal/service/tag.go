package service

import (
	"Aoi/internal/model"
	"fmt"
)

// TagReq 关于tag的请求
type TagReq struct {
	ID        uint32 `json:"id,omitempty" form:"id" `
	Name      string `json:"name,omitempty" form:"name" `
	State     uint8  `json:"state,omitempty" form:"state,default=1" binding:"oneof=0 1"`
	CreatedBy string `json:"createdBy,omitempty" form:"createdBy" `
	UpdatedBy string `json:"updatedBy,omitempty" form:"updatedBy" `
	Page      int    `json:"page,omitempty" form:"page"`
	Size      int    `json:"size,omitempty" form:"size"`
}

func (svc *Service) CountTag(para *TagReq) (int, error) {
	return svc.dao.CountTag(para.Name, para.State)
}

func (svc *Service) GetList(para *TagReq) ([]*model.Tag, error) {
	fmt.Println("service state: ", para.State)
	return svc.dao.ListTag(para.State, para.Page, para.Size)
}

func (svc *Service) UpdateTag(para *TagReq) error {
	return svc.dao.UpdateTag(uint(para.ID), para.Name, para.UpdatedBy, para.State)
}
func (svc *Service) CreateTag(para *TagReq) error {
	return svc.dao.CreateTag(para.Name, para.CreatedBy, para.State)
}
func (svc *Service) DeleteTag(para *TagReq) error {
	return svc.dao.DeleteTag(uint(para.ID))
}
