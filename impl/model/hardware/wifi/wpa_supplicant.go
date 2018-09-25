package wifi

var (
	defaultWpaSupplicantConf = &WPASupplicantConf{
		UpdateConfig:                      true,
		CtrlInterface:                     "/var/run/wpa_supplicant",
		EapolVersion:                      1,
		ApScan:                            1,
		PassiveScan:                       false,
		UserMpm:                           1,
		MaxPeerLinks:                      99,
		MeshMaxInactivity:                 300,
		CertInCb:                          true,
		FastReauth:                        true,
		OpensslCiphers:                    "DEFAULT:!EXP:!LOW",
		Dot11RSNAConfigPMKLifetime:        432000,
		Dot11RSNAConfigPMKReauthThreshold: 70,
		Dot11RSNAConfigSATimeout:          60,
		AutoUUID:                          false,
		BssMaxCount:                       200,
		P2PDisabled:                       true,
		P2PGoMaxInactivity:                300,
		P2PPassphraseLen:                  8,
		P2PSearchDelay:                    500,
		DtimPeriod:                        2,
		BeaconInt:                         100,
		RandAddrLifetime:                  60,
		GasRandAddrLifetime:               60,
		GoInternet:                        1,
		GoVenueGroup:                      7,
		GoVenueType:                       1,
	}
)

type WPASupplicantCredConf struct {
	Temporary                 bool     `json:"temporary"`
	Priority                  int      `json:"priority"`
	Pcsc                      bool     `json:"pcsc"`
	Realm                     string   `json:"realm"`
	Username                  string   `json:"username"`
	Password                  string   `json:"password"`
	CaCert                    string   `json:"ca_cert"`
	ClientCert                string   `json:"client_cert"`
	PrivateKey                string   `json:"private_key"`
	PrivateKeyPasswd          string   `json:"private_key_passwd"`
	Imsi                      string   `json:"imsi"`
	Milenage                  string   `json:"milenage"`
	Domain                    string   `json:"domain"`
	RoamingConsortium         string   `json:"roaming_consortium"`
	RequiredRoamingConsortium string   `json:"required_roaming_consortium"`
	RoamingConsortiums        []string `json:"roaming_consortiums"`
	Eap                       []string `json:"eap"`
	Phase1                    string   `json:"phase1"`
	Phase2                    string   `json:"phase2"`
	ExcludedSSID              []string `json:"excluded_ssid"`
	RoamingPartner            string   `json:"roaming_partner"`
	UpdateIdentifier          string   `json:"update_identifier"`
	ProvisioningSP            string   `json:"provisioning_sp"`
	MaxBSSLoad                int      `json:"max_bss_load"`
	ReqConnCapab              bool     `json:"req_conn_capab"`
	Ocsp                      int      `json:"ocsp"`
	SimNum                    string   `json:"sim_num"`
}

// WPASupplicantNetConf network section inside wpa_supplicant.conf
type WPASupplicantNetConf struct {
	SSID              string `json:"ssid"`
	Bssid             string `json:"bssid"`
	ScanSSID          int    `json:"scan_ssid"`
	KeyMgmt           string `json:"key_mgmt"`
	Pairwise          string `json:"pairwise"`
	Group             string `json:"group"`
	PSK               string `json:"psk"`
	Eap               string `json:"eap"`
	AnonymousIdentity string `json:"anonymous_identity"`
	PacFile           string `json:"pac_file"`
	EapolFlags        int    `json:"eapol_flags"`
	Identity          string `json:"identity"`
	Password          string `json:"password"`
	Pcsc              string `json:"pcsc"`
	CACert            string `json:"ca_cert"`
	ClientCert        string `json:"client_cert"`
	PrivateKey        string `json:"private_key"`
	PrivateKeyPasswd  string `json:"private_key_passwd"`
	Phase1            string `json:"phase_1"`
	Phase2            string `json:"phase_2"`
	CACert2           string `json:"ca_cert_2"`
	ClientCert2       string `json:"client_cert_2"`
	PrivateKey2       string `json:"private_key_2"`
	PrivateKey2Passwd string `json:"private_key_2_passwd"`
	Pin               string `json:"pin"`
	Proto             string `json:"proto"`
	Mode              int    `json:"mode"`
	Frequency         int    `json:"frequency"`
	Priority          int    `json:"priority"`
}

// WPASupplicantConf struct
type WPASupplicantConf struct {
	UpdateConfig                      bool                    `json:"update_config"`
	CtrlInterface                     string                  `json:"ctrl_interface"`
	EapolVersion                      int                     `json:"eapol_version"`
	ApScan                            int                     `json:"ap_scan"`
	PassiveScan                       bool                    `json:"passive_scan"`
	UserMpm                           int                     `json:"user_mpm"`
	MaxPeerLinks                      int                     `json:"max_peer_links"`
	MeshMaxInactivity                 int                     `json:"mesh_max_inactivity"`
	CertInCb                          bool                    `json:"cert_in_cb"`
	FastReauth                        bool                    `json:"fast_reauth"`
	OpenscEnginePath                  string                  `json:"opensc_engine_path"`
	Pkcs11EnginePath                  string                  `json:"pkcs11_engine_path"`
	Pkcs11ModulePath                  string                  `json:"pkcs11_module_path"`
	OpensslCiphers                    string                  `json:"openssl_ciphers"`
	LoadDynamicEap                    string                  `json:"load_dynamic_eap"`
	DriverParam                       string                  `json:"driver_param"`
	Country                           string                  `json:"country"`
	Dot11RSNAConfigPMKLifetime        int                     `json:"dot11RSNAConfigPMKLifetime"`
	Dot11RSNAConfigPMKReauthThreshold int                     `json:"dot11RSNAConfigPMKReauthThreshold"`
	Dot11RSNAConfigSATimeout          int                     `json:"dot11RSNAConfigSATimeout"`
	UUID                              string                  `json:"uuid"`
	AutoUUID                          bool                    `json:"auto_uuid"`
	DeviceName                        string                  `json:"device_name"`
	Manufacturer                      string                  `json:"manufacturer"`
	ModelName                         string                  `json:"model_name"`
	ModelNumber                       string                  `json:"model_number"`
	SerialNumber                      string                  `json:"serial_number"`
	DeviceType                        string                  `json:"device_type"`
	OsVersion                         string                  `json:"os_version"`
	ConfigMethods                     []string                `json:"config_methods"`
	WpsCredProcessing                 int                     `json:"wps_cred_processing"`
	WpsVendorExtM1                    string                  `json:"wps_vendor_ext_m1"`
	WpsNfcDevPwID                     int                     `json:"wps_nfc_dev_pw_id"`
	WpsNfcDhPubkey                    string                  `json:"wps_nfc_dh_pubkey"`
	WpsNfcDhPrivkey                   string                  `json:"wps_nfc_dh_privkey"`
	WpsNfcDevPw                       string                  `json:"wps_nfc_dev_pw"`
	WpsPriority                       int                     `json:"wps_priority"`
	BssMaxCount                       int                     `json:"bss_max_count"`
	Autoscan                          string                  `json:"autoscan"`
	FilterSsids                       bool                    `json:"filter_ssids"`
	P2PDisabled                       bool                    `json:"p2p_disabled"`
	P2PGoMaxInactivity                int                     `json:"p2p_go_max_inactivity"`
	P2PPassphraseLen                  int                     `json:"p2p_passphrase_len"`
	P2PSearchDelay                    int                     `json:"p2p_search_delay"`
	Okc                               int                     `json:"okc"`
	Pmf                               int                     `json:"pmf"`
	SaeGroups                         []int                   `json:"sae_groups"`
	DtimPeriod                        int                     `json:"dtim_period"`
	BeaconInt                         int                     `json:"beacon_int"`
	ApVendorElements                  string                  `json:"ap_vendor_elements"`
	IgnoreOldScanRes                  int                     `json:"ignore_old_scan_res"`
	ScanCurFreq                       int                     `json:"scan_cur_freq"`
	MacAddr                           int                     `json:"mac_addr"`
	RandAddrLifetime                  int                     `json:"rand_addr_lifetime"`
	PreassocMacAddr                   int                     `json:"preassoc_mac_addr"`
	GasRandMacAddr                    int                     `json:"gas_rand_mac_addr"`
	GasRandAddrLifetime               int                     `json:"gas_rand_addr_lifetime"`
	GoInterworking                    bool                    `json:"go_interworking"`
	GoAccessNetworkType               int                     `json:"go_access_network_type"`
	GoInternet                        int                     `json:"go_internet"`
	GoVenueGroup                      int                     `json:"go_venue_group"`
	GoVenueType                       int                     `json:"go_venue_type"`
	Hessid                            string                  `json:"hessid"`
	AutoInterworking                  int                     `json:"auto_interworking"`
	GasAddress3                       int                     `json:"gas_address3"`
	FtmResponder                      int                     `json:"ftm_responder"`
	FtmInitiator                      int                     `json:"ftm_initiator"`
	MboCellCapa                       int                     `json:"mbo_cell_capa"`
	Oce                               int                     `json:"oce"`
	MemOnlyPsk                        int                     `json:"mem_only_psk"`
	ApMaxInactivity                   int                     `json:"ap_max_inactivity"`
	WpsDisabled                       int                     `json:"wps_disabled"`
	FilsDhGroup                       int                     `json:"fils_dh_group"`
	Credential                        []WPASupplicantCredConf `json:"creds"`
	Network                           []WPASupplicantNetConf  `json:"networks"`
}

// func (w *WPASupplicantConf) UnmarshalText(text []byte) error {
// 	return nil
// }
//
// func (w *WPASupplicantConf) MarshalText() ([]byte, error) {
// 	return nil, nil
// }
