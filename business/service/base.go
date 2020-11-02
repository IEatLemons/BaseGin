package service

import (
	resp "github.com/WeEatLemon/Go-Common/helper/responses"
	"github.com/IEatLemons/BaseGin/business/dao"
)

type UserService struct {
	D    *dao.BaseDao
	resp *resp.Resp
}
