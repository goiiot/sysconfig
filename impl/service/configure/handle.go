package configure

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goiiot/sysconfig/cmd/sysconfig/conf"
	"github.com/goiiot/sysconfig/impl/log"
	"github.com/goiiot/sysconfig/impl/service/utils"
	"github.com/lytics/confl"
	"github.com/pelletier/go-toml"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

var (
	shell                  string
	errConfFmtNotSupported = errors.New("config file format not supported")
	errWPAConfigInvalid    = errors.New("wpa_supplicant.conf file invalid")
)

func InitServiceConfig(v1 *gin.RouterGroup, config *conf.ServiceConfigure) {
	if !config.Enabled {
		return
	}

	shell = config.Shell

	initLoraConfigure(v1, &config.Lora)
	initNetConfigure(v1, &config.Network)
	initPeriphConfigure(v1, &config.Periph)
	initBusConfigure(v1, &config.Bus)
}

func initCommonData(
	config *conf.CommonConfigure,
) (m map[string]*conf.CommonDevice, l []*conf.CommonDevice, hasData bool) {
	m = make(map[string]*conf.CommonDevice)
	l = make([]*conf.CommonDevice, 0)
	for i, v := range config.DeviceList {
		if v.Enabled {
			m[v.Name] = &config.DeviceList[i]
			l = append(l, &config.DeviceList[i])
		}
	}

	if len(l) < 1 {
		m = nil
		l = nil
		hasData = false
		return
	} else {
		hasData = true
		return
	}
}

func handleCommonList(list []*conf.CommonDevice) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		utils.RespOkJSON(ctx, list)
	}
}

func getConfig(
	m map[string]*conf.CommonDevice,
	decoder func(typ, fmt string, data []byte) (interface{}, error),
	ctx *gin.Context,
) (interface{}, error) {
	name := ctx.Param("name")
	dev, ok := m[name]
	if !ok {
		return nil, fmt.Errorf("no such device: %s", name)
	}

	content, err := ioutil.ReadFile(dev.ConfigFile)
	if err != nil {
		log.W("failed to read data form conf file", zap.String("file", dev.ConfigFile), zap.Error(err))
		return nil, fmt.Errorf("can't read config file: %v", err)
	}

	if dev.ConfigFmt == "text" {
		return string(content), nil
	}

	if decoder == nil {
		log.E("no decoder for non text conf file provided", zap.String("fmt", dev.ConfigFmt), zap.String("name", dev.Name))
		return nil, fmt.Errorf("no decoder for non text conf file provided")
	}

	data, err := decoder(dev.Type, dev.ConfigFmt, content)
	if err != nil {
		return nil, fmt.Errorf("can't parse config: %v", err)
	}

	return data, nil
}

func getStatus(
	m map[string]*conf.CommonDevice,
	ctx *gin.Context,
	statusInstance func() interface{},
) (interface{}, error) {
	name := ctx.Param("name")
	dev, ok := m[name]
	if !ok {
		return nil, fmt.Errorf("no such device: %s", name)
	}

	result, err := run(shell, dev.HelperScript, "status", name)
	if err != nil {
		return nil, err
	}

	status := statusInstance()
	err = confl.Unmarshal(result, status)
	if err != nil {
		log.E("failed to parse status to json", zap.Error(err))
		return nil, fmt.Errorf("parse status to json failed: %v", err)
	}

	return status, nil
}

func handleCommonGetInfo(
	m map[string]*conf.CommonDevice,
	decoder func(typ, fmt string, data []byte) (interface{}, error),
	statusInstance func() interface{},
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		config, err := getConfig(m, decoder, ctx)
		if err != nil {
			utils.RespErrJSON(ctx, http.StatusOK, 1, err.Error())
			return
		}

		status, err := getStatus(m, ctx, statusInstance)
		if err != nil {
			utils.RespErrJSON(ctx, http.StatusOK, 1, err.Error())
			return
		}

		utils.RespOkJSON(ctx, gin.H{
			"status": status,
			"config": config,
		})
	}
}

func handleCommonGetConfig(
	m map[string]*conf.CommonDevice,
	decoder func(typ, fmt string, data []byte) (interface{}, error),
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		config, err := getConfig(m, decoder, ctx)
		if err != nil {
			utils.RespErrJSON(ctx, http.StatusOK, 1, err.Error())
			return
		}

		utils.RespOkJSON(ctx, config)
	}
}

// converter converts json body to device config file format if config file format is not text
func handleCommonUpdateConfig(
	m map[string]*conf.CommonDevice,
	converter func(typ, fmt string, postData []byte) ([]byte, error),
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Param("name")
		dev, ok := m[name]
		if !ok {
			utils.RespErrJSON(ctx, http.StatusNotFound, 1, fmt.Sprintf("no such device: %s", name))
			return
		}

		data, err := ctx.GetRawData()
		if err != nil {
			log.W("failed to read data form request", zap.Error(err))
			utils.RespErrJSON(ctx, http.StatusBadRequest, 1, fmt.Sprintf("failed to read request body: %v", err))
			return
		}

		if dev.ConfigFmt == "text" {
			if err := ioutil.WriteFile(dev.ConfigFile, data, 0644); err != nil {
				log.E("failed to write to config", zap.String("file", dev.ConfigFile), zap.Error(err))
				utils.RespErrJSON(ctx, http.StatusInternalServerError, 1, fmt.Sprintf("failed to write to config %s: %v", dev.ConfigFile, err))
				return
			}

			utils.RespOkJSON(ctx)
			return
		}

		if converter == nil {
			log.E("no converter for non text conf file provided", zap.String("fmt", dev.ConfigFmt), zap.String("name", dev.Name))
			utils.RespErrJSON(ctx, http.StatusInternalServerError, 1, "no converter for non text conf file provided")
			return
		}

		toStore, err := converter(dev.Type, dev.ConfigFmt, data)
		if err != nil {
			utils.RespErrJSON(ctx, http.StatusInternalServerError, 1, "")
			return
		}

		if err = ioutil.WriteFile(dev.ConfigFile, toStore, 0644); err != nil {
			log.E("failed to write to config", zap.String("file", dev.ConfigFile), zap.Error(err))
			utils.RespErrJSON(ctx, http.StatusInternalServerError, 1, fmt.Sprintf("failed to write to config %s: %v", dev.ConfigFile, err))
			return
		}

		utils.RespOkJSON(ctx)
	}
}

func handleCommonAction(
	m map[string]*conf.CommonDevice,
	statusInstance func() interface{},
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Param("name")
		dev, ok := m[name]
		if !ok {
			utils.RespErrJSON(ctx, http.StatusNotFound, 1, fmt.Sprintf("no such device: %s", name))
			return
		}
		action := ctx.Query("action")

		switch action {
		case "start", "stop", "restart":
			// good to go
			result, err := run(shell, dev.HelperScript, action, name)
			if err != nil {
				log.E("failed to execute action", zap.String("name", name), zap.String("action", action), zap.Error(err))
				utils.RespErrJSON(ctx, http.StatusInternalServerError, 1, fmt.Sprintf("failed to execute %s action %s: %v", dev.Name, action, err))
				return
			}

			utils.RespOkJSON(ctx, string(result))
			return
		case "status":
			// convert output ucl to json after
			status, err := getStatus(m, ctx, statusInstance)
			if err != nil {
				utils.RespErrJSON(ctx, http.StatusOK, 1, err.Error())
				return
			}

			utils.RespOkJSON(ctx, status)
			return
		default:
			utils.RespErrJSON(ctx, http.StatusBadRequest, 1, fmt.Sprintf("unsupported action: %s", action))
			return
		}
	}
}

func marshal(fmt string, in interface{}) ([]byte, error) {
	var f func(interface{}) ([]byte, error)
	switch fmt {
	case "json":
		f = json.Marshal
	case "toml":
		f = toml.Marshal
	case "yaml":
		f = yaml.Marshal
	case "ucl":
		f = confl.Marshal
	default:
		return nil, errConfFmtNotSupported
	}

	return f(in)
}

func unmarshal(fmt string, data []byte, out interface{}) (interface{}, error) {
	var f func([]byte, interface{}) error
	switch fmt {
	case "json":
		f = json.Unmarshal
	case "toml":
		f = toml.Unmarshal
	case "yaml":
		f = yaml.Unmarshal
	case "ucl":
		f = confl.Unmarshal
	default:
		return nil, errConfFmtNotSupported
	}

	return out, f(data, out)
}
