package configure

import (
	"encoding/json"
	"errors"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/goiiot/sysconfig/cmd/sysconfig/conf"
	"github.com/goiiot/sysconfig/impl/model"
	"github.com/goiiot/sysconfig/impl/model/hardware/wifi"
)

var (
	netSectionLookup    = regexp.MustCompile(`network.*?=.*?{`)
	netSectionEndLookup = regexp.MustCompile(`}`)
)

var (
	errWifiTypeNotSupported = errors.New("unsupported wifi type")
)

const (
	typeWPASupplicant = "wpa_supplicant"
)

func getWifiInstanceByType(typ string) interface{} {
	switch typ {
	case typeWPASupplicant:
		return new(wifi.WpaSupplicantConfig)
	default:
		return nil
	}
}

// func formatWpaSupplicantConf(data []byte) (interface{}, error) {
// 	// find first network section and make it a array
// 	startLoc := netSectionLookup.FindIndex(data)
// 	if startLoc == nil {
// 		// only global fields presents
// 		return unmarshal("ucl", data, make(map[string]interface{}))
// 	}
//
// 	// embed network sections into an array
// 	data = append(data[:startLoc[1]-1], append([]byte{'['}, data[startLoc[1]-1:]...)...)
// 	data = append(data[:startLoc[1]], netSectionLookup.ReplaceAll(data[startLoc[1]:], []byte("{"))...)
// 	data = netSectionEndLookup.ReplaceAll(data[:], []byte("},"))
//
// 	lastEndLoc := []int{0, 0}
// 	endLoc := netSectionEndLookup.FindIndex(data)
// 	for ; endLoc != nil && endLoc[1] > lastEndLoc[1] && endLoc[1]+1 < len(data); endLoc = netSectionEndLookup.FindIndex(data[endLoc[1]+1:]) {
// 		lastEndLoc = endLoc
// 	}
//
// 	if lastEndLoc != nil && len(lastEndLoc) == 2 {
// 		data = append(data[:lastEndLoc[1]], append(data[lastEndLoc[1]:], ']')...)
// 		return unmarshal("ucl", data, make(map[string]interface{}))
// 	} else {
// 		return nil, errWPAConfigInvalid
// 	}
// }

func wifiConfigDecoder(typ, format string, data []byte) (interface{}, error) {
	w := getWifiInstanceByType(typ)
	if w == nil {
		return nil, errWifiTypeNotSupported
	}
	return unmarshal(format, data, w)
}

func wifiConfigConverter(typ, format string, postData []byte) ([]byte, error) {
	w := getWifiInstanceByType(typ)
	if w == nil {
		return nil, errWifiTypeNotSupported
	}

	if err := json.Unmarshal(postData, w); err != nil {
		return nil, err
	}

	return marshal(format, w)
}

func wifiStatusInstance() interface{} {
	return new(model.WifiStatus)
}

func ifaceConfigDecoder(typ, format string, data []byte) (interface{}, error) {
	return nil, nil
}

func ifaceConfigConverter(typ, format string, postData []byte) ([]byte, error) {
	return nil, nil
}

func ifaceStatusInstance() interface{} {
	return new(model.Status)
}

func cellConfigDecoder(typ, format string, data []byte) (interface{}, error) {
	return nil, nil
}

func cellConfigConverter(typ, format string, postData []byte) ([]byte, error) {
	return nil, nil
}

func cellStatusInstance() interface{} {
	return new(model.Status)
}

func initNetConfigure(v1 *gin.RouterGroup, config *conf.ServiceConfigureNetwork) {
	if !config.Enabled {
		return
	}

	if config.WifiConfig.Enabled {
		wifiConfig, wifiList, ok := initCommonData(&config.WifiConfig)
		if ok {
			v1.GET("/configure/net/wifi", handleCommonList(wifiList))
			v1.GET("/configure/net/wifi/:name", handleCommonGetInfo(wifiConfig, wifiConfigDecoder, wifiStatusInstance))
			v1.GET("/configure/net/wifi/:name/conf", handleCommonGetConfig(wifiConfig, wifiConfigDecoder))
			v1.PUT("/configure/net/wifi/:name/conf", handleCommonUpdateConfig(wifiConfig, wifiConfigConverter))
			v1.POST("/configure/net/wifi/:name/action", handleCommonAction(wifiConfig, wifiStatusInstance))
		}
	}

	if config.Interfaces.Enabled {
		ifaceConfig, ifaceList, ok := initCommonData(&config.Interfaces)
		if ok {
			v1.GET("/configure/net/iface", handleCommonList(ifaceList))
			v1.GET("/configure/net/iface/:name", handleCommonGetInfo(ifaceConfig, ifaceConfigDecoder, ifaceStatusInstance))
			v1.GET("/configure/net/iface/:name/conf", handleCommonGetConfig(ifaceConfig, ifaceConfigDecoder))
			v1.PUT("/configure/net/iface/:name/conf", handleCommonUpdateConfig(ifaceConfig, ifaceConfigConverter))
			v1.POST("/configure/net/iface/:name/action", handleCommonAction(ifaceConfig, ifaceStatusInstance))
		}
	}

	if config.CellularConfig.Enabled {
		cellConfig, cellList, ok := initCommonData(&config.CellularConfig)
		if ok {
			v1.GET("/configure/net/cell", handleCommonList(cellList))
			v1.GET("/configure/net/cell/:name", handleCommonGetInfo(cellConfig, cellConfigDecoder, cellStatusInstance))
			v1.GET("/configure/net/cell/:name/conf", handleCommonGetConfig(cellConfig, cellConfigDecoder))
			v1.PUT("/configure/net/cell/:name/conf", handleCommonUpdateConfig(cellConfig, cellConfigConverter))
			v1.POST("/configure/net/cell/:name/action", handleCommonAction(cellConfig, cellStatusInstance))
		}
	}
}
