package main

import (
	"fmt"
	"github.com/IEatLemons/BaseGin/api"
	"github.com/IEatLemons/BaseGin/config"
	"github.com/IEatLemons/BaseGin/link/base"
	"github.com/IEatLemons/BaseGin/logger"
	"github.com/WeEatLemon/Go-Common/helper"
	"github.com/WeEatLemon/Go-Common/middle"
	"github.com/WeEatLemon/Go-Common/signal"
	"go.uber.org/zap"
)

func main() {
	vHandle := config.New()
	path, _ := helper.GetProjectRoot()
	err := vHandle.InitConfig(path)
	if err != nil {
		logger.Error("init config failed", zap.String("error", err.Error()))
		panic("init config failed:" + err.Error())
		return
	}
	// init logger
	if err := logger.InitLogger(&vHandle.Config.Log); err != nil {
		logger.Error("init logger failed", zap.String("error", err.Error()))
		panic("init logger failed:" + err.Error())
		return
	}

	// init MySQL
	if err := base.InitEngineGroup(&vHandle.Config.MysqlMaster, &vHandle.Config.MysqlSlave); err != nil {
		logger.Error("init mysql failed", zap.String("error", err.Error()))
		panic("init mysql failed:" + err.Error())
		return
	}
	// init redis
	if err := base.InitRedis(&vHandle.Config.Redis); err != nil {
		logger.Error("init redis failed", zap.String("error", err.Error()))
		panic("init redis failed:" + err.Error())
		return
	}

	// init Middle
	middle.Init(base.Redis, vHandle.Config.InternalHost, vHandle.Config.InternalIp)

	// init router
	r := api.SetupApi()
	// listen and signal
	if err := signal.Listen(fmt.Sprintf(":%d", vHandle.Config.Server.Port), r); err != nil {
		logger.Error("listen and signal failed", zap.String("error", err.Error()))
		panic("init redis failed:" + err.Error())
		return
	}

}
