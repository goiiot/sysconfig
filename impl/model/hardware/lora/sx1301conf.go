package lora

import (
	"encoding/json"
)

// GetDefaultPktFwdConfig get same configuration as noted int https://github.com/Lora-net/packet_forwarder/blob/master/lora_pkt_fwd/global_conf.json
func GetDefaultPktFwdConfig() *PktFwdConfig {
	ret := *defaultPktFwdConfig
	return &ret
}

var defaultPktFwdConfig *PktFwdConfig

func init() {
	defaultPktFwdConfig = new(PktFwdConfig)
	if err := json.Unmarshal([]byte(defPktFwdJSONConfig), defaultPktFwdConfig); err != nil {
		// TODO check json unmarshal error
	}

	// fmt.Println(defaultPktFwdConfig)
}

// PktFwdConfig struct
type PktFwdConfig struct {
	SX1301Conf struct {
		LorawanPublic bool `json:"lorawan_public"`
		Clksrc        int  `json:"clksrc"`
		LbtCfg        struct {
			Enable     bool `json:"enable"`
			RssiTarget int  `json:"rssi_target"`
			ChanCfg    []struct {
				FreqHz     int `json:"freq_hz"`
				ScanTimeUs int `json:"scan_time_us"`
			} `json:"chan_cfg"`
			Sx127XRssiOffset int `json:"sx127x_rssi_offset"`
		} `json:"lbt_cfg"`
		AntennaGain int `json:"antenna_gain"`
		Radio0      struct {
			Enable      bool    `json:"enable"`
			Type        string  `json:"type"`
			Freq        int     `json:"freq"`
			RssiOffset  float64 `json:"rssi_offset"`
			TxEnable    bool    `json:"tx_enable"`
			TxNotchFreq int     `json:"tx_notch_freq"`
			TxFreqMin   int     `json:"tx_freq_min"`
			TxFreqMax   int     `json:"tx_freq_max"`
		} `json:"radio_0"`
		Radio1 struct {
			Enable     bool    `json:"enable"`
			Type       string  `json:"type"`
			Freq       int     `json:"freq"`
			RssiOffset float64 `json:"rssi_offset"`
			TxEnable   bool    `json:"tx_enable"`
		} `json:"radio_1"`
		ChanMultiSF0 struct {
			Enable bool `json:"enable"`
			Radio  int  `json:"radio"`
			If     int  `json:"if"`
		} `json:"chan_multiSF_0"`
		ChanMultiSF1 struct {
			Enable bool `json:"enable"`
			Radio  int  `json:"radio"`
			If     int  `json:"if"`
		} `json:"chan_multiSF_1"`
		ChanMultiSF2 struct {
			Enable bool `json:"enable"`
			Radio  int  `json:"radio"`
			If     int  `json:"if"`
		} `json:"chan_multiSF_2"`
		ChanMultiSF3 struct {
			Enable bool `json:"enable"`
			Radio  int  `json:"radio"`
			If     int  `json:"if"`
		} `json:"chan_multiSF_3"`
		ChanMultiSF4 struct {
			Enable bool `json:"enable"`
			Radio  int  `json:"radio"`
			If     int  `json:"if"`
		} `json:"chan_multiSF_4"`
		ChanMultiSF5 struct {
			Enable bool `json:"enable"`
			Radio  int  `json:"radio"`
			If     int  `json:"if"`
		} `json:"chan_multiSF_5"`
		ChanMultiSF6 struct {
			Enable bool `json:"enable"`
			Radio  int  `json:"radio"`
			If     int  `json:"if"`
		} `json:"chan_multiSF_6"`
		ChanMultiSF7 struct {
			Enable bool `json:"enable"`
			Radio  int  `json:"radio"`
			If     int  `json:"if"`
		} `json:"chan_multiSF_7"`
		ChanLoraStd struct {
			Enable       bool `json:"enable"`
			Radio        int  `json:"radio"`
			If           int  `json:"if"`
			Bandwidth    int  `json:"bandwidth"`
			SpreadFactor int  `json:"spread_factor"`
		} `json:"chan_Lora_std"`
		ChanFSK struct {
			Enable    bool `json:"enable"`
			Radio     int  `json:"radio"`
			If        int  `json:"if"`
			Bandwidth int  `json:"bandwidth"`
			Datarate  int  `json:"datarate"`
		} `json:"chan_FSK"`
		TxLut0 struct {
			PaGain  int `json:"pa_gain"`
			MixGain int `json:"mix_gain"`
			RfPower int `json:"rf_power"`
			DigGain int `json:"dig_gain"`
		} `json:"tx_lut_0"`
		TxLut1 struct {
			PaGain  int `json:"pa_gain"`
			MixGain int `json:"mix_gain"`
			RfPower int `json:"rf_power"`
			DigGain int `json:"dig_gain"`
		} `json:"tx_lut_1"`
		TxLut2 struct {
			PaGain  int `json:"pa_gain"`
			MixGain int `json:"mix_gain"`
			RfPower int `json:"rf_power"`
			DigGain int `json:"dig_gain"`
		} `json:"tx_lut_2"`
		TxLut3 struct {
			PaGain  int `json:"pa_gain"`
			MixGain int `json:"mix_gain"`
			RfPower int `json:"rf_power"`
			DigGain int `json:"dig_gain"`
		} `json:"tx_lut_3"`
		TxLut4 struct {
			PaGain  int `json:"pa_gain"`
			MixGain int `json:"mix_gain"`
			RfPower int `json:"rf_power"`
			DigGain int `json:"dig_gain"`
		} `json:"tx_lut_4"`
		TxLut5 struct {
			PaGain  int `json:"pa_gain"`
			MixGain int `json:"mix_gain"`
			RfPower int `json:"rf_power"`
			DigGain int `json:"dig_gain"`
		} `json:"tx_lut_5"`
		TxLut6 struct {
			PaGain  int `json:"pa_gain"`
			MixGain int `json:"mix_gain"`
			RfPower int `json:"rf_power"`
			DigGain int `json:"dig_gain"`
		} `json:"tx_lut_6"`
		TxLut7 struct {
			PaGain  int `json:"pa_gain"`
			MixGain int `json:"mix_gain"`
			RfPower int `json:"rf_power"`
			DigGain int `json:"dig_gain"`
		} `json:"tx_lut_7"`
		TxLut8 struct {
			PaGain  int `json:"pa_gain"`
			MixGain int `json:"mix_gain"`
			RfPower int `json:"rf_power"`
			DigGain int `json:"dig_gain"`
		} `json:"tx_lut_8"`
		TxLut9 struct {
			PaGain  int `json:"pa_gain"`
			MixGain int `json:"mix_gain"`
			RfPower int `json:"rf_power"`
			DigGain int `json:"dig_gain"`
		} `json:"tx_lut_9"`
		TxLut10 struct {
			PaGain  int `json:"pa_gain"`
			MixGain int `json:"mix_gain"`
			RfPower int `json:"rf_power"`
			DigGain int `json:"dig_gain"`
		} `json:"tx_lut_10"`
		TxLut11 struct {
			PaGain  int `json:"pa_gain"`
			MixGain int `json:"mix_gain"`
			RfPower int `json:"rf_power"`
			DigGain int `json:"dig_gain"`
		} `json:"tx_lut_11"`
		TxLut12 struct {
			PaGain  int `json:"pa_gain"`
			MixGain int `json:"mix_gain"`
			RfPower int `json:"rf_power"`
			DigGain int `json:"dig_gain"`
		} `json:"tx_lut_12"`
		TxLut13 struct {
			PaGain  int `json:"pa_gain"`
			MixGain int `json:"mix_gain"`
			RfPower int `json:"rf_power"`
			DigGain int `json:"dig_gain"`
		} `json:"tx_lut_13"`
		TxLut14 struct {
			PaGain  int `json:"pa_gain"`
			MixGain int `json:"mix_gain"`
			RfPower int `json:"rf_power"`
			DigGain int `json:"dig_gain"`
		} `json:"tx_lut_14"`
		TxLut15 struct {
			PaGain  int `json:"pa_gain"`
			MixGain int `json:"mix_gain"`
			RfPower int `json:"rf_power"`
			DigGain int `json:"dig_gain"`
		} `json:"tx_lut_15"`
	} `json:"SX1301_conf"`
	GatewayConf struct {
		GatewayID          string `json:"gateway_ID"`
		ServerAddress      string `json:"server_address"`
		ServPortUp         int    `json:"serv_port_up"`
		ServPortDown       int    `json:"serv_port_down"`
		KeepaliveInterval  int    `json:"keepalive_interval"`
		StatInterval       int    `json:"stat_interval"`
		PushTimeoutMs      int    `json:"push_timeout_ms"`
		ForwardCrcValid    bool   `json:"forward_crc_valid"`
		ForwardCrcError    bool   `json:"forward_crc_error"`
		ForwardCrcDisabled bool   `json:"forward_crc_disabled"`
	} `json:"gateway_conf"`
}

const defPktFwdJSONConfig = `
{
    "SX1301_conf": {
        "lorawan_public": true,
        "clksrc": 1,
        "lbt_cfg": {
            "enable": false,
            "rssi_target": -80,
            "chan_cfg":[
                { "freq_hz": 867100000, "scan_time_us": 128 },
                { "freq_hz": 867300000, "scan_time_us": 5000 },
                { "freq_hz": 867500000, "scan_time_us": 128 },
                { "freq_hz": 869525000, "scan_time_us": 128 }
            ],
            "sx127x_rssi_offset": -4
        },
        "antenna_gain": 0,
        "radio_0": {
            "enable": true,
            "type": "SX1257",
            "freq": 867500000,
            "rssi_offset": -166.0,
            "tx_enable": true,
            "tx_notch_freq": 129000,
            "tx_freq_min": 863000000,
            "tx_freq_max": 870000000
        },
        "radio_1": {
            "enable": true,
            "type": "SX1257",
            "freq": 868500000,
            "rssi_offset": -166.0,
            "tx_enable": false
        },
        "chan_multiSF_0": {
        
            "enable": true,
            "radio": 1,
            "if": -400000
        },
        "chan_multiSF_1": {
        
            "enable": true,
            "radio": 1,
            "if": -200000
        },
        "chan_multiSF_2": {
        
            "enable": true,
            "radio": 1,
            "if": 0
        },
        "chan_multiSF_3": {
        
            "enable": true,
            "radio": 0,
            "if": -400000
        },
        "chan_multiSF_4": {
        
            "enable": true,
            "radio": 0,
            "if": -200000
        },
        "chan_multiSF_5": {
        
            "enable": true,
            "radio": 0,
            "if": 0
        },
        "chan_multiSF_6": {
        
            "enable": true,
            "radio": 0,
            "if": 200000
        },
        "chan_multiSF_7": {
        
            "enable": true,
            "radio": 0,
            "if": 400000
        },
        "chan_Lora_std": {
        
            "enable": true,
            "radio": 1,
            "if": -200000,
            "bandwidth": 250000,
            "spread_factor": 7
        },
        "chan_FSK": {
        
            "enable": true,
            "radio": 1,
            "if": 300000,
            "bandwidth": 125000,
            "datarate": 50000
        },
        "tx_lut_0": {
        
            "pa_gain": 0,
            "mix_gain": 8,
            "rf_power": -6,
            "dig_gain": 0
        },
        "tx_lut_1": {
        
            "pa_gain": 0,
            "mix_gain": 10,
            "rf_power": -3,
            "dig_gain": 0
        },
        "tx_lut_2": {
        
            "pa_gain": 0,
            "mix_gain": 12,
            "rf_power": 0,
            "dig_gain": 0
        },
        "tx_lut_3": {
        
            "pa_gain": 1,
            "mix_gain": 8,
            "rf_power": 3,
            "dig_gain": 0
        },
        "tx_lut_4": {
        
            "pa_gain": 1,
            "mix_gain": 10,
            "rf_power": 6,
            "dig_gain": 0
        },
        "tx_lut_5": {
        
            "pa_gain": 1,
            "mix_gain": 12,
            "rf_power": 10,
            "dig_gain": 0
        },
        "tx_lut_6": {
        
            "pa_gain": 1,
            "mix_gain": 13,
            "rf_power": 11,
            "dig_gain": 0
        },
        "tx_lut_7": {
        
            "pa_gain": 2,
            "mix_gain": 9,
            "rf_power": 12,
            "dig_gain": 0
        },
        "tx_lut_8": {
        
            "pa_gain": 1,
            "mix_gain": 15,
            "rf_power": 13,
            "dig_gain": 0
        },
        "tx_lut_9": {
        
            "pa_gain": 2,
            "mix_gain": 10,
            "rf_power": 14,
            "dig_gain": 0
        },
        "tx_lut_10": {
        
            "pa_gain": 2,
            "mix_gain": 11,
            "rf_power": 16,
            "dig_gain": 0
        },
        "tx_lut_11": {
        
            "pa_gain": 3,
            "mix_gain": 9,
            "rf_power": 20,
            "dig_gain": 0
        },
        "tx_lut_12": {
        
            "pa_gain": 3,
            "mix_gain": 10,
            "rf_power": 23,
            "dig_gain": 0
        },
        "tx_lut_13": {
        
            "pa_gain": 3,
            "mix_gain": 11,
            "rf_power": 25,
            "dig_gain": 0
        },
        "tx_lut_14": {
        
            "pa_gain": 3,
            "mix_gain": 12,
            "rf_power": 26,
            "dig_gain": 0
        },
        "tx_lut_15": {
        
            "pa_gain": 3,
            "mix_gain": 14,
            "rf_power": 27,
            "dig_gain": 0
        }
    },
    "gateway_conf": {
        "gateway_ID": "AA555A0000000000",
    
        "server_address": "localhost",
        "serv_port_up": 1680,
        "serv_port_down": 1680,
    
        "keepalive_interval": 10,
        "stat_interval": 30,
        "push_timeout_ms": 100,
    
        "forward_crc_valid": true,
        "forward_crc_error": false,
        "forward_crc_disabled": false
    }
}
`
