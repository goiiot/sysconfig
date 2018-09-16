package conf

import "time"

const (
	cfgServiceShellEnabled        = "service.shell.enabled"
	cfgServiceFileUploadEnabled   = "service.file.upload_enabled"
	cfgServiceFileDownloadEnabled = "service.file.download_enabled"
	cfgServiceFileDefaultPath     = "service.file.default_path"
	// cfgServiceConfigureLoraEnabled     = "service.configure.lora.enabled"
	// cfgServiceConfigureNetworkEnabled  = "service.configure.network.enabled"
	// cfgServiceConfigureNetworkEthernet = "service.configure.network.ethernet"
	// cfgServiceConfigureNetworkCellular = "service.configure.network.cellular"
	// cfgServiceConfigureBusRS485        = "service.configure.bus.rs485"
	// cfgServiceConfigureBusCAN          = "service.configure.bus.can"
)

var serviceFlags = []Flag{
	boolFlag(&config.Service.Shell.Enabled, cfgServiceShellEnabled, "enable web shell or not"),
	boolFlag(&config.Service.File.EnableUpload, cfgServiceFileUploadEnabled, "enable file upload or not"),
	boolFlag(&config.Service.File.EnableDownload, cfgServiceFileDownloadEnabled, "enable file download or not"),
	stringFlag(&config.Service.File.DefaultPath, cfgServiceFileDefaultPath, "", "", "default dir for uploaded files"),
	// mapArrayFlag(cfgServiceConfigureNetworkWifi, "wifi devices"),
	// stringSliceFlag(&config.Service.Configure.Network.WifiList, cfgServiceConfigureNetworkWifi, "wifi devices"),
}

type CommonDevice struct {
	Name         string `yaml:"name" json:"name"`
	Type         string `yaml:"type" json:"type"`
	Enabled      bool   `yaml:"enabled" json:"enabled"`
	ConfigFile   string `yaml:"conf_file" json:"config_file"`
	ConfigFmt    string `yaml:"conf_fmt" json:"config_fmt"`
	HelperScript string `yaml:"helper_script" json:"helper_script"`
}

type CommonConfigure struct {
	Enabled    bool           `yaml:"enabled" json:"enabled"`
	DeviceList []CommonDevice `yaml:"devs" json:"device_list"`
}

// ServiceConfigureNetwork service.configure.net
type ServiceConfigureNetwork struct {
	Enabled        bool            `yaml:"enabled"`
	Interfaces     CommonConfigure `yaml:"interfaces"`
	CellularConfig CommonConfigure `yaml:"cellular"`
	WifiConfig     CommonConfigure `yaml:"wifi"`
}

// ServiceConfigure service.configure
type ServiceConfigure struct {
	Enabled bool                    `yaml:"enabled"`
	Shell   string                  `yaml:"shell"`
	Lora    CommonConfigure         `yaml:"lora"`
	Network ServiceConfigureNetwork `yaml:"network"`
	Bus     CommonConfigure         `yaml:"bus"`
	Periph  CommonConfigure         `yaml:"periph"`
}

// ServiceMonitoring service.monitoring
type ServiceMonitoring struct {
	Enabled         bool          `yaml:"enabled"`
	RefreshInterval time.Duration `yaml:"refresh_interval"`
}

// ServiceShell service.shell
type ServiceShell struct {
	Enabled      bool   `yaml:"enabled"`
	DefaultShell string `yaml:"default_shell"`
}

// ServiceFile service.file
type ServiceFile struct {
	EnableUpload   bool   `yaml:"upload_enabled"`
	EnableDownload bool   `yaml:"download_enabled"`
	DefaultPath    string `yaml:"default_path"`
}

// ServicePower service.power
type ServicePower struct {
	EnableReboot   bool `yaml:"reboot_enabled"`
	EnableShutdown bool `yaml:"shutdown_enabled"`
}

// ServiceConfig service
type ServiceConfig struct {
	Shell      ServiceShell      `yaml:"shell"`
	File       ServiceFile       `yaml:"file"`
	Monitoring ServiceMonitoring `yaml:"monitoring"`
	Power      ServicePower      `yaml:"power"`
	Configure  ServiceConfigure  `yaml:"configure"`
}
