package configure

import (
	"github.com/gin-gonic/gin"
	"github.com/goiiot/sysconfig/cmd/sysconfig/conf"
	"github.com/goiiot/sysconfig/impl/model"
)

func busConfigDecoder(typ, fmt string, data []byte) (interface{}, error) {
	return nil, nil
}

func busConfigConverter(typ, fmt string, postData []byte) ([]byte, error) {
	return nil, nil
}

func busStatusInstance() interface{} {
	return new(model.WifiStatus)
}

func initBusConfigure(v1 *gin.RouterGroup, config *conf.CommonConfigure) {
	if !config.Enabled {
		return
	}

	busConfig, busList, ok := initCommonData(config)
	if ok {
		v1.GET("/configure/bus", handleCommonList(busList))
		v1.GET("/configure/bus/:name", handleCommonGetInfo(busConfig, busConfigDecoder, busStatusInstance))
		v1.GET("/configure/bus/:name/conf", handleCommonGetConfig(busConfig, busConfigDecoder))
		v1.PUT("/configure/bus/:name/conf", handleCommonUpdateConfig(busConfig, busConfigConverter))
		v1.POST("/configure/bus/:name/action", handleCommonAction(busConfig, busStatusInstance))
	}
}
