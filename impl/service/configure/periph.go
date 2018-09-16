package configure

import (
	"github.com/gin-gonic/gin"
	"github.com/goiiot/sysconfig/cmd/sysconfig/conf"
	"github.com/goiiot/sysconfig/impl/model"
)

func periphConfigDecoder(typ, format string, data []byte) (interface{}, error) {
	return nil, nil
}

func periphConfigConvert(typ, format string, postData []byte) ([]byte, error) {
	return nil, nil
}

func periphStatusInstance() interface{} {
	return new(model.Status)
}

func initPeriphConfigure(v1 *gin.RouterGroup, config *conf.CommonConfigure) {
	if !config.Enabled {
		return
	}

	periphConfig, periphList, ok := initCommonData(config)
	if ok {
		v1.GET("/configure/periph", handleCommonList(periphList))
		v1.GET("/configure/periph/:name", handleCommonGetInfo(periphConfig, periphConfigDecoder, periphStatusInstance))
		v1.GET("/configure/periph/:name/conf", handleCommonGetConfig(periphConfig, periphConfigDecoder))
		v1.PUT("/configure/periph/:name/conf", handleCommonUpdateConfig(periphConfig, periphConfigConvert))
		v1.POST("/configure/periph/:name/action", handleCommonAction(periphConfig, periphStatusInstance))
	}
}
