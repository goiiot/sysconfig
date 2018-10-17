package wifi

import (
	"bytes"
	"testing"
	"text/template"
)

func TestWPASupplicant(t *testing.T) {
	tplt, err := template.New("wpa_supplicant_conf").Funcs(funcMap).Parse(defWpaSupplicantConf)
	if err != nil {
		t.Errorf("failed to parse wpa_supplicant_conf template: %v", err)
	}
	buf := &bytes.Buffer{}
	data := &WPASupplicantConf{
		CtrlInterface: "/test/path",
		ApScan:        0,
		SaeGroups:     []int{1, 2, 3},
	}

	if err = tplt.Execute(buf, data); err != nil {
		t.Errorf("failed to execute wpa_supplicant_conf template: %v", err)
	}

	println(buf.String())
}

const defWpaSupplicantConf = `
{{ if .UpdateConfig }}update_config=1{{ end }}
{{ if HasValue .CtrlInterface }}ctrl_interface={{ .CtrlInterface }}{{ end }}
{{ if HasValue .EapolVersion }}eapol_version={{ .EapolVersion }}{{ end }}
{{ if HasValue .ApScan }}ap_scan={{ .ApScan }}{{ end }}
{{ if .PassiveScan }}passive_scan=1{{ end }}
{{ if HasValue .UserMpm }}user_mpm=1{{ end }}
{{ if HasValue .MaxPeerLinks }}max_peer_links={{ .MaxPeerLinks }}{{ end }}
{{ if HasValue .MeshMaxInactivity }}mesh_max_inactivity={{ .MeshMaxInactivity }}{{ end }}
{{ if .CertInCb }}cert_in_cb=1{{ end }}
`
const def = `
{{ if not eq .FastReauth 0 }}
fast_reauth={{ .FastReauth }}
{{ end }}
{{ if lt len .OpenscEnginePath 0 }}
opensc_engine_path={{ .OpenscEnginePath }}
{{ end }}
{{ if lt len .Pkcs11EnginePath 0 }}
pkcs11_engine_path={{ .Pkcs11EnginePath }}
{{ end }}
{{ if lt len .Pkcs11ModulePath 0 }}
pkcs11_module_path={{ .Pkcs11ModulePath }}
{{ end }}
{{ if lt len .OpensslCiphers 0 }}
openssl_ciphers={{ .OpensslCiphers }}
{{ end }}
{{ if lt len .LoadDynamicEap 0 }}
load_dynamic_eap={{ .LoadDynamicEap }}
{{ end }}
{{ if lt len .DriverParam 0 }}
driver_param={{ .DriverParam }}
{{ end }}
{{ if lt len .Country 0 }}
country={{ .Country }}
{{ end }}
{{ if not eq .Dot11RSNAConfigPMKLifetime 0 }}
dot11RSNAConfigPMKLifetime={{ .Dot11RSNAConfigPMKLifetime }}
{{ end }}
{{ if not eq .Dot11RSNAConfigPMKReauthThreshold 0 }}
dot11RSNAConfigPMKReauthThreshold={{ .Dot11RSNAConfigPMKReauthThreshold }}
{{ end }}
{{ if not eq .Dot11RSNAConfigSATimeout 0 }}
dot11RSNAConfigSATimeout={{ .Dot11RSNAConfigSATimeout }}
{{ end }}
{{ if lt len .UUID 0 }}
uuid={{ .UUID }}
{{ end }}
{{ if not eq .AutoUUID 0 }}
auto_uuid={{ .AutoUUID }}
{{ end }}
{{ if lt len .DeviceName 0 }}
device_name={{ .DeviceName }}
{{ end }}
{{ if lt len .Manufacturer 0 }}
manufacturer={{ .Manufacturer }}
{{ end }}
{{ if lt len .ModelName 0 }}
model_name={{ .ModelName }}
{{ end }}
{{ if not eq .ModelNumber 0 }}
model_number={{ .ModelNumber }}
{{ end }}
{{ if not eq .SerialNumber 0 }}
serial_number={{ .SerialNumber }}
{{ end }}
{{ if lt len .DeviceType 0 }}
device_type={{ .DeviceType }}
{{ end }}
{{ if lt len .OsVersion 0 }}
os_version={{ .OsVersion }}
{{ end }}
{{ if lt len .ConfigMethods 0 }}
config_methods={{ .ConfigMethods }}
{{ end }}
{{ if not eq .WpsCredProcessing 0 }}
wps_cred_processing={{ .WpsCredProcessing }}
{{ end }}
{{ if lt len .WpsVendorExtM1 0 }}
wps_vendor_ext_m1={{ .WpsVendorExtM1 }}
{{ end }}
{{ if not eq .WpsNfcDevPwID 0 }}
wps_nfc_dev_pw_id={{ .WpsNfcDevPwID }}
{{ end }}
{{ if lt len .WpsNfcDhPubkey 0 }}
wps_nfc_dh_pubkey={{ .WpsNfcDhPubkey }}
{{ end }}
{{ if lt len .WpsNfcDhPrivkey 0 }}
wps_nfc_dh_privkey={{ .WpsNfcDhPrivkey }}
{{ end }}
{{ if lt len .WpsNfcDevPw 0 }}
wps_nfc_dev_pw={{ .WpsNfcDevPw }}
{{ end }}
{{ if not eq .WpsPriority 0 }}
wps_priority={{ .WpsPriority }}
{{ end }}
{{ if not eq .BssMaxCount 0 }}
bss_max_count={{ .BssMaxCount }}
{{ end }}
{{ if lt len .Autoscan 0 }}
autoscan={{ .Autoscan }}
{{ end }}
{{ if not eq .FilterSsids 0 }}
filter_ssids={{ .FilterSsids }}
{{ end }}
{{ if not eq .P2PGoMaxInactivity 0 }}
p2p_go_max_inactivity={{ .P2PGoMaxInactivity }}
{{ end }}
{{ if not eq .P2PPassphraseLen 0 }}
p2p_passphrase_len={{ .P2PPassphraseLen }}
{{ end }}
{{ if not eq .P2PSearchDelay 0 }}
p2p_search_delay={{ .P2PSearchDelay }}
{{ end }}
{{ if not eq .Okc 0 }}
okc={{ .Okc }}
{{ end }}
{{ if not eq .Pmf 0 }}
pmf={{ .Pmf }}
{{ end }}
{{ if lt len .SaeGroups 0 }}
sae_groups={{ .SaeGroups }}
{{ end }}
{{ if not eq .DtimPeriod 0 }}
dtim_period={{ .DtimPeriod }}
{{ end }}
{{ if not eq .BeaconInt 0 }}
beacon_int={{ .BeaconInt }}
{{ end }}
{{ if lt len .ApVendorElements 0 }}
ap_vendor_elements={{ .ApVendorElements }}
{{ end }}
{{ if not eq .IgnoreOldScanRes 0 }}
ignore_old_scan_res={{ .IgnoreOldScanRes }}
{{ end }}
{{ if not eq .MacAddr 0 }}
mac_addr={{ .MacAddr }}
{{ end }}
{{ if not eq .RandAddrLifetime 0 }}
rand_addr_lifetime={{ .RandAddrLifetime }}
{{ end }}
{{ if not eq .PreassocMacAddr 0 }}
preassoc_mac_addr={{ .PreassocMacAddr }}
{{ end }}
{{ if not eq .GasRandMacAddr 0 }}
gas_rand_mac_addr={{ .GasRandMacAddr }}
{{ end }}
{{ if not eq .GasRandAddrLifetime 0 }}
gas_rand_addr_lifetime={{ .GasRandAddrLifetime }}
{{ end }}
{{ if not eq .GoAccessNetworkType 0 }}
go_access_network_type={{ .GoAccessNetworkType }}
{{ end }}
{{ if not eq .GoInternet 0 }}
go_internet={{ .GoInternet }}
{{ end }}
{{ if not eq .GoVenueGroup 0 }}
go_venue_group={{ .GoVenueGroup }}
{{ end }}
{{ if not eq .GoVenueType 0 }}
go_venue_type={{ .GoVenueType }}
{{ end }}
{{ if not eq .AutoInterworking 0 }}
auto_interworking={{ .AutoInterworking }}
{{ end }}
{{ if not eq .GasAddress3 0 }}
gas_address3={{ .GasAddress3 }}
{{ end }}
{{ if not eq .FtmResponder 0 }}
ftm_responder={{ .FtmResponder }}
{{ end }}
{{ if not eq .FtmInitiator 0 }}
ftm_initiator={{ .FtmInitiator }}
{{ end }}
{{ if not eq .MboCellCapa 0 }}
mbo_cell_capa={{ .MboCellCapa }}
{{ end }}
{{ if not eq .Oce 0 }}
oce={{ .Oce }}
{{ end }}
{{ if not eq .MemOnlyPsk 0 }}
mem_only_psk={{ .MemOnlyPsk }}
{{ end }}
{{ if not eq .ApMaxInactivity 0 }}
ap_max_inactivity={{ .ApMaxInactivity }}
{{ end }}
{{ if not eq .WpsDisabled 0 }}
wps_disabled={{ .WpsDisabled }}
{{ end }}
{{ if not eq .FilsDhGroup 0 }}
fils_dh_group={{ .FilsDhGroup }}
{{ end }}

{{ range .Network }}
network={
	{{ if lt len .SSID 0 }}
	ssid={{ .SSID }}
	{{ end }}
	{{ if lt len .Bssid 0 }}
	bssid={{ .Bssid }}
	{{ end }}
	{{ if not eq .ScanSSID 0 }}
	scan_ssid={{ .ScanSSID}}
	{{ end }}
	{{ if lt len .KeyMgmt 0 }}
	key_mgmt={{ .KeyMgmt }}
	{{ end }}
	{{ if lt len .Pairwise 0 }}
	pairwise={{ .Pairwise}}
	{{ end }}
	{{ if lt len .Group 0 }}
	group={{ .Group }}
	{{ end }}
	{{ if lt len .PSK 0 }}
	psk={{ .PSK }}
	{{ end }}
	{{ if lt len .Eap 0 }}
	eap={{ .Eap }}
	{{ end }}
	{{ if lt len .AnonymousIdentity 0 }}
	anonymous_identity={{ .AnonymousIdentity }}
	{{ end }}
	{{ if lt len .PacFile 0 }}
	pac_file={{ .PacFile }}
	{{ end }}
	{{ if not eq .EapolFlags 0 }}
	eapol_flags={{ .EapolFlags }}
	{{ end }}
	{{ if lt len .Identity 0 }}
	identity={{ .Identity}}
	{{ end }}
	{{ if lt len .Password 0 }}
	password={{ .Password}}
	{{ end }}
	{{ if lt len .Pcsc 0 }}
	pcsc={{ .Pcsc }}
	{{ end }}
	{{ if lt len .CACert 0 }}
	ca_cert={{ .CACert }}
	{{ end }}
	{{ if lt len .ClientCert 0 }}
	client_cert={{ .ClientCert }}
	{{ end }}
	{{ if lt len .PrivateKey 0 }}
	private_key={{ .PrivateKey }}
	{{ end }}
	{{ if lt len .PrivateKeyPasswd 0 }}
	private_key_passwd={{ .PrivateKeyPasswd }}
	{{ end }}
	{{ if lt len .Phase1 0 }}
	phase_1={{ .Phase1 }}
	{{ end }}
	{{ if lt len .Phase2 0 }}
	phase_2={{ .Phase2 }}
	{{ end }}
	{{ if lt len .CACert2 0 }}
	ca_cert_2={{ .CACert2 }}
	{{ end }}
	{{ if lt len .ClientCert20 }} client_cert_2={{ .ClientCert2}}
	{{ end }}
	{{ if lt len .PrivateKey20 }} private_key_2={{ .PrivateKey2}}
	{{ end }}
	{{ if lt len .PrivateKey2Passwd 0 }}
	private_key_2_passwd={{ .PrivateKey2Passwd }}
	{{ end }}
	{{ if lt len .Pin 0 }}
	pin={{ .Pin }}
	{{ end }}
	{{ if lt len .Proto 0 }}
	proto={{ .Proto }}
	{{ end }}
	{{ if not eq .Mode 0 }}
	mode={{ .Mode }}
	{{ end }}
	{{ if not eq .Frequency 0 }}
	frequency={{ .Frequency }}
	{{ end }}
	{{ if not eq .Priority 0 }}
	priority={{ .Priority}}
	{{ end }}
}
{{ end }}
`
