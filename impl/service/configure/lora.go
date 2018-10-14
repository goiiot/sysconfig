package configure

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/goiiot/sysconfig/cmd/sysconfig/conf"
	"github.com/goiiot/sysconfig/impl/model"
	"github.com/goiiot/sysconfig/impl/model/hardware/lora"
)

var (
	errLoraTypeNotSupported = errors.New("unsupported lora device type")
)

const (
	typeBrocaarLoraGatewayBridge = "brocaar/lora-gateway-bridge"
	typePacketForwarder          = "pkt_forwarder"
)

func getInstanceFromType(typ string) interface{} {
	switch typ {
	case typePacketForwarder:
		return new(lora.PktFwdConfig)
	case typeBrocaarLoraGatewayBridge:
		m := make(map[string]interface{})
		return &m
	default:
		return nil
	}
}

func loraConfigDecoder(typ, formant string, data []byte) (interface{}, error) {
	d := getInstanceFromType(typ)
	if d == nil {
		return nil, errLoraTypeNotSupported
	}

	return unmarshal(formant, data, d)
}

func loraConfigConverter(typ, format string, postData []byte) ([]byte, error) {
	d := getInstanceFromType(typ)
	if d == nil {
		return nil, errLoraTypeNotSupported
	}

	err := json.Unmarshal(postData, d)
	if err != nil {
		return nil, err
	}

	return marshal(format, d)
}

func loraStatusInstance() interface{} {
	return new(model.LoraStatus)
}

func initLoraConfigure(v1 *gin.RouterGroup, config *conf.CommonConfigure) {
	if !config.Enabled {
		return
	}

	loraConfig, loraList, ok := initCommonData(config)
	if ok {
		v1.GET("/configure/lora", handleCommonList(loraList))
		v1.GET("/configure/lora/:name", handleCommonGetInfo(loraConfig, loraConfigDecoder, loraStatusInstance))
		v1.GET("/configure/lora/:name/conf", handleCommonGetConfig(loraConfig, loraConfigDecoder))
		v1.PUT("/configure/lora/:name/conf", handleCommonUpdateConfig(loraConfig, loraConfigConverter))
		v1.POST("/configure/lora/:name/action", handleCommonAction(loraConfig, loraStatusInstance))
	}
}
