package controllers

import (
	resp "github.com/WeEatLemon/Go-Common/helper/responses"
	"github.com/WeEatLemon/Go-Common/middle"
	"github.com/gin-gonic/gin"
)

var Resp *resp.Resp

/* 基础路由 */
func BaseRoute(g *gin.RouterGroup) {
	M := middle.New()
	g.Use(M.Cors(), M.AuthLanguage(), M.AuthPlatform())
	Resp = resp.InitResp(M.Language)

	Base := g.Group("/")

	RegisteredRouter(Base)
}

/* 注册相关模块的路由 */
func RegisteredRouter(g *gin.RouterGroup) {

}
