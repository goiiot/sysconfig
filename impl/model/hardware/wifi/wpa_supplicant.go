package wifi

// WpaSupplicantNetworkConfig network section inside wpa_supplicant.conf
type WpaSupplicantNetworkConfig struct {
	SSID              string `json:"ssid" confl:"ssid"`
	Bssid             string `json:"bssid" confl:"bssid"`
	ScanSSID          int    `json:"scan_ssid" confl:"scan_ssid"`
	KeyMgmt           string `json:"key_mgmt" confl:"key_mgmt"`
	Pairwise          string `json:"pairwise" confl:"pairwise"`
	Group             string `json:"group" confl:"group"`
	PSK               string `json:"psk" confl:"psk"`
	Eap               string `json:"eap" confl:"eap"`
	AnonymousIdentity string `json:"anonymous_identity" confl:"anonymous_identity"`
	PacFile           string `json:"pac_file" confl:"pac_file"`
	EapolFlags        int    `json:"eapol_flags" confl:"eapol_flags"`
	Identity          string `json:"identity" confl:"identity"`
	Password          string `json:"password" confl:"password"`
	Pcsc              string `json:"pcsc" confl:"pcsc"`
	CACert            string `json:"ca_cert" confl:"ca_cert"`
	ClientCert        string `json:"client_cert" confl:"client_cert"`
	PrivateKey        string `json:"private_key" confl:"private_key"`
	PrivateKeyPasswd  string `json:"private_key_passwd" confl:"private_key_passwd"`
	Phase1            string `json:"phase_1" confl:"phase1"`
	Phase2            string `json:"phase_2" confl:"phase2"`
	CACert2           string `json:"ca_cert_2" confl:"ca_cert2"`
	ClientCert2       string `json:"client_cert_2" confl:"client_cert2"`
	PrivateKey2       string `json:"private_key_2" confl:"private_key2"`
	PrivateKey2Passwd string `json:"private_key_2_passwd" confl:"private_key2_passwd"`
	Pin               string `json:"pin" confl:"pin"`
	Proto             string `json:"proto" confl:"proto"`
	Mode              int    `json:"mode" confl:"mode"`
	Frequency         int    `json:"frequency" confl:"frequency"`
	Priority          int    `json:"priority" confl:"priority"`
}

// WpaSupplicantConfig struct
type WpaSupplicantConfig struct {
	UpdateConfig  int                        `json:"update_config" confl:"update_config"`
	CtrlInterface string                     `json:"ctrl_interface" confl:"ctrl_interface"`
	Network       WpaSupplicantNetworkConfig `json:"network" confl:"network"`
}

// func (w *WpaSupplicantConfig) UnmarshalText(text []byte) error {
// 	return nil
// }
//
// func (w *WpaSupplicantConfig) MarshalText() ([]byte, error) {
// 	return nil, nil
// }
