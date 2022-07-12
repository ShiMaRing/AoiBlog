package service

// TagReq 关于tag的请求
type TagReq struct {
	ID        uint32 `json:"id,omitempty" form:"id" `
	Name      string `json:"name,omitempty" form:"name" `
	State     uint8  `json:"state,omitempty" form:"state default=1" binding:"oneof=0 1"`
	CreatedBy string `json:"createdBy,omitempty" form:"createdBy" `
	UpdatedBy string `json:"updatedBy,omitempty" form:"updatedBy" `
}
