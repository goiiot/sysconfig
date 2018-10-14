export const loraDefaultConfig = {
  "SX1301_conf": {
    "lorawan_public": true,
    "clksrc": 1,
    "lbt_cfg": {
      "enable": false,
      "rssi_target": -80,
      "chan_cfg": [
        {
          "freq_hz": 867100000,
          "scan_time_us": 128
        },
        {
          "freq_hz": 867300000,
          "scan_time_us": 5000
        },
        {
          "freq_hz": 867500000,
          "scan_time_us": 128
        },
        {
          "freq_hz": 869525000,
          "scan_time_us": 128
        }
      ],
      "sx127x_rssi_offset": -4
    },
    "antenna_gain": 0,
    "radio_0": {
      "enable": true,
      "type": "SX1257",
      "freq": 867500000,
      "rssi_offset": -166,
      "tx_enable": true,
      "tx_notch_freq": 129000,
      "tx_freq_min": 863000000,
      "tx_freq_max": 870000000
    },
    "radio_1": {
      "enable": true,
      "type": "SX1257",
      "freq": 868500000,
      "rssi_offset": -166,
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
};

export const loraBridgeDefaultConfig ={
  "general": {
    "log_level": 0
  },
  "packet_forwarder": {
    "udp_bind": "",
    "skip_crc_check": false,
    "configuration": null
  },
  "backend": {
    "mqtt": {
      "uplink_topic_template": "",
      "downlink_topic_template": "",
      "stats_topic_template": "",
      "ack_topic_template": "",
      "config_topic_template": "",
      "marshaler": "protobuf",
      "auth": {
        "type": "{{ .Backend.MQTT.Auth.Type }}",
        "generic": {
          "server": "tcp://",
          "username": "test_username",
          "password": "test_password",
          "ca_cert": "",
          "tls_cert": "",
          "tls_key": "",
          "qos": 0,
          "clean_session": false,
          "client_id": "",
          "max_reconnect_interval": 0
        }
      },
      "server": "",
      "username": "",
      "password": "",
      "ca_cert": "",
      "tls_cert": "",
      "tls_key": "",
      "qos": 0,
      "clean_session": false,
      "client_id": "",
      "max_reconnect_interval": 0,
      "always_subscribe_ma_cs": null
    }
  },
  "metrics": {
    "prometheus": {
      "endpoint_enabled": false,
      "bind": "0.0.0.0:9090"
    }
  }
};
