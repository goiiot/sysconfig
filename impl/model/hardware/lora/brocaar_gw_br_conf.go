package lora

// The MIT License (MIT)
//
// Copyright (c) 2018 Orne Brocaar
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

import (
	"time"

	"github.com/brocaar/lorawan"
)

// GenericConfig defines the generic configuration.
type GenericConfig struct {
	Server               string        `json:"server"`
	Username             string        `json:"username"`
	Password             string        `json:"password"`
	CACert               string        `mapstructure:"ca_cert" json:"ca_cert"`
	TLSCert              string        `mapstructure:"tls_cert" json:"tls_cert"`
	TLSKey               string        `mapstructure:"tls_key" json:"tls_key"`
	QOS                  uint8         `mapstructure:"qos" json:"qos"`
	CleanSession         bool          `mapstructure:"clean_session" json:"clean_session"`
	ClientID             string        `mapstructure:"client_id" json:"client_id"`
	MaxReconnectInterval time.Duration `mapstructure:"max_reconnect_interval" json:"max_reconnect_interval"`
}

// BackendAuthConfig holds the MQTT pub-sub backend auth configuration.
type BackendAuthConfig struct {
	Type    string        `json:"type"`
	Generic GenericConfig `json:"generic"`
}

// BackendConfig holds the MQTT pub-sub backend configuration.
type BackendConfig struct {
	UplinkTopicTemplate   string            `mapstructure:"uplink_topic_template" json:"uplink_topic_template"`
	DownlinkTopicTemplate string            `mapstructure:"downlink_topic_template" json:"downlink_topic_template"`
	StatsTopicTemplate    string            `mapstructure:"stats_topic_template" json:"stats_topic_template"`
	AckTopicTemplate      string            `mapstructure:"ack_topic_template" json:"ack_topic_template"`
	ConfigTopicTemplate   string            `mapstructure:"config_topic_template" json:"config_topic_template"`
	Marshaler             string            `mapstructure:"marshaler" json:"marshaler"`
	Auth                  BackendAuthConfig `json:"auth"`

	// for backwards compatibility
	Server               string        `json:"server"`
	Username             string        `json:"username"`
	Password             string        `json:"password"`
	CACert               string        `mapstructure:"ca_cert" json:"ca_cert"`
	TLSCert              string        `mapstructure:"tls_cert" json:"tls_cert"`
	TLSKey               string        `mapstructure:"tls_key" json:"tls_key"`
	QOS                  uint8         `mapstructure:"qos" json:"qos"`
	CleanSession         bool          `mapstructure:"clean_session" json:"clean_session"`
	ClientID             string        `mapstructure:"client_id" json:"client_id"`
	MaxReconnectInterval time.Duration `mapstructure:"max_reconnect_interval" json:"max_reconnect_interval"`

	AlwaysSubscribeMACs []lorawan.EUI64 `mapstructure:"-" json:"always_subscribe_ma_cs"`
}

type PFConfiguration struct {
	MAC            lorawan.EUI64 `mapstructure:"-" json:"mac"`
	MACString      string        `mapstructure:"mac" json:"mac_string"`
	BaseFile       string        `mapstructure:"base_file" json:"base_file"`
	OutputFile     string        `mapstructure:"output_file" json:"output_file"`
	RestartCommand string        `mapstructure:"restart_command" json:"restart_command"`
	Version        string        `mapstructure:"-" json:"version"`
}

// BrocaarGWBRConfig defines the configuration structure.
type BrocaarGWBRConfig struct {
	General struct {
		LogLevel int `mapstructure:"log_level" json:"log_level"`
	} `json:"general"`

	PacketForwarder struct {
		UDPBind      string `mapstructure:"udp_bind" json:"udp_bind"`
		SkipCRCCheck bool   `mapstructure:"skip_crc_check" json:"skip_crc_check"`

		Configuration []PFConfiguration `mapstructure:"configuration" json:"configuration"`
	} `mapstructure:"packet_forwarder" json:"packet_forwarder"`

	Backend struct {
		MQTT BackendConfig `json:"mqtt"`
	} `json:"backend"`
	Metrics struct {
		Prometheus struct {
			EndpointEnabled bool   `mapstructure:"endpoint_enabled" json:"endpoint_enabled"`
			Bind            string `json:"bind"`
		} `json:"prometheus"`
	} `json:"metrics"`
}
