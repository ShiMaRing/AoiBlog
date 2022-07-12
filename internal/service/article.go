package service

// ArticleReq 提交的关于Article的请求
type ArticleReq struct {
	ID            uint   `json:"id" form:"id"`
	Title         string `json:"title" form:"title"`
	Desc          string `json:"desc" form:"desc"`
	Content       string `json:"content" form:"content"`
	CoverImageUrl string `json:"coverImageUrl" form:"coverImageUrl"`
	State         uint8  `json:"state" form:"state default=1"`
	CreatedBy     string `json:"createdBy" form:"createdBy"`
	UpdatedBy     string `json:"updatedBy" form:"updatedBy"`
	Page          int    `json:"page" form:"page"`
	Size          int    `json:"size" form:"size"`
}
