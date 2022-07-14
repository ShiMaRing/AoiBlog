package service

import "errors"

type AuthRequest struct {
	AppKey    string `json:"appKey,omitempty" form:"appKey" binding:"required"`
	AppSecret string `json:"appSecret,omitempty" form:"appSecret" binding:"required"`
}

func (svc *Service) CheckAuth(param *AuthRequest) error {
	auth, err := svc.dao.GetAuth(param.AppKey, param.AppSecret)
	if err != nil {
		return err
	}
	if auth.ID > 0 {
		return nil
	}
	return errors.New("Auth Fail !!")
}
