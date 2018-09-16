package wifi

import "testing"

func TestWPASupplicant(t *testing.T) {
	// os.Env()
}

const defWpaSupplicantConf = `
##### Example wpa_supplicant configuration file ###############################
#
# This file describes configuration file format and lists all available option.
# Please also take a look at simpler configuration examples in 'examples'
# subdirectory.
#
# Empty lines and lines starting with # are ignored

# NOTE! This file may contain password information and should probably be made
# readable only by root user on multiuser systems.

# Note: All file paths in this configuration file should use full (absolute,
# not relative to working directory) path in order to allow working directory
# to be changed. This can happen if wpa_supplicant is run in the background.

# Whether to allow wpa_supplicant to update (overwrite) configuration
#
# This option can be used to allow wpa_supplicant to overwrite configuration
# file whenever configuration is changed (e.g., new network block is added with
# wpa_cli or wpa_gui, or a password is changed). This is required for
# wpa_cli/wpa_gui to be able to store the configuration changes permanently.
# Please note that overwriting configuration file will remove the comments from
# it.
#update_config=1

# global configuration (shared by all network blocks)
#
# Parameters for the control interface. If this is specified, wpa_supplicant
# will open a control interface that is available for external programs to
# manage wpa_supplicant. The meaning of this string depends on which control
# interface mechanism is used. For all cases, the existence of this parameter
# in configuration is used to determine whether the control interface is
# enabled.
#
# For UNIX domain sockets (default on Linux and BSD): This is a directory that
# will be created for UNIX domain sockets for listening to requests from
# external programs (CLI/GUI, etc.) for status information and configuration.
# The socket file will be named based on the interface name, so multiple
# wpa_supplicant processes can be run at the same time if more than one
# interface is used.
# /var/run/wpa_supplicant is the recommended directory for sockets and by
# default, wpa_cli will use it when trying to connect with wpa_supplicant.
#
# Access control for the control interface can be configured by setting the
# directory to allow only members of a group to use sockets. This way, it is
# possible to run wpa_supplicant as root (since it needs to change network
# configuration and open raw sockets) and still allow GUI/CLI components to be
# run as non-root users. However, since the control interface can be used to
# change the network configuration, this access needs to be protected in many
# cases. By default, wpa_supplicant is configured to use gid 0 (root). If you
# want to allow non-root users to use the control interface, add a new group
# and change this value to match with that group. Add users that should have
# control interface access to this group. If this variable is commented out or
# not included in the configuration file, group will not be changed from the
# value it got by default when the directory or socket was created.
#
# When configuring both the directory and group, use following format:
# DIR=/var/run/wpa_supplicant GROUP=wheel
# DIR=/var/run/wpa_supplicant GROUP=0
# (group can be either group name or gid)
#
# For UDP connections (default on Windows): The value will be ignored. This
# variable is just used to select that the control interface is to be created.
# The value can be set to, e.g., udp (ctrl_interface=udp)
#
# For Windows Named Pipe: This value can be used to set the security descriptor
# for controlling access to the control interface. Security descriptor can be
# set using Security Descriptor String Format (see http://msdn.microsoft.com/
# library/default.asp?url=/library/en-us/secauthz/security/
# security_descriptor_string_format.asp). The descriptor string needs to be
# prefixed with SDDL=. For example, ctrl_interface=SDDL=D: would set an empty
# DACL (which will reject all connections). See README-Windows.txt for more
# information about SDDL string format.
#
ctrl_interface=/var/run/wpa_supplicant

# IEEE 802.1X/EAPOL version
# wpa_supplicant is implemented based on IEEE Std 802.1X-2004 which defines
# EAPOL version 2. However, there are many APs that do not handle the new
# version number correctly (they seem to drop the frames completely). In order
# to make wpa_supplicant interoperate with these APs, the version number is set
# to 1 by default. This configuration value can be used to set it to the new
# version (2).
# Note: When using MACsec, eapol_version shall be set to 3, which is
# defined in IEEE Std 802.1X-2010.
eapol_version=1

# AP scanning/selection
# By default, wpa_supplicant requests driver to perform AP scanning and then
# uses the scan results to select a suitable AP. Another alternative is to
# allow the driver to take care of AP scanning and selection and use
# wpa_supplicant just to process EAPOL frames based on IEEE 802.11 association
# information from the driver.
# 1: wpa_supplicant initiates scanning and AP selection; if no APs matching to
#    the currently enabled networks are found, a new network (IBSS or AP mode
#    operation) may be initialized (if configured) (default)
# 0: driver takes care of scanning, AP selection, and IEEE 802.11 association
#    parameters (e.g., WPA IE generation); this mode can also be used with
#    non-WPA drivers when using IEEE 802.1X mode; do not try to associate with
#    APs (i.e., external program needs to control association). This mode must
#    also be used when using wired Ethernet drivers (including MACsec).
# 2: like 0, but associate with APs using security policy and SSID (but not
#    BSSID); this can be used, e.g., with ndiswrapper and NDIS drivers to
#    enable operation with hidden SSIDs and optimized roaming; in this mode,
#    the network blocks in the configuration file are tried one by one until
#    the driver reports successful association; each network block should have
#    explicit security policy (i.e., only one option in the lists) for
#    key_mgmt, pairwise, group, proto variables
# Note: ap_scan=2 should not be used with the nl80211 driver interface (the
# current Linux interface). ap_scan=1 is optimized work working with nl80211.
# For finding networks using hidden SSID, scan_ssid=1 in the network block can
# be used with nl80211.
# When using IBSS or AP mode, ap_scan=2 mode can force the new network to be
# created immediately regardless of scan results. ap_scan=1 mode will first try
# to scan for existing networks and only if no matches with the enabled
# networks are found, a new IBSS or AP mode network is created.
ap_scan=1

# Whether to force passive scan for network connection
#
# By default, scans will send out Probe Request frames on channels that allow
# active scanning. This advertise the local station to the world. Normally this
# is fine, but users may wish to do passive scanning where the radio should only
# listen quietly for Beacon frames and not send any Probe Request frames. Actual
# functionality may be driver dependent.
#
# This parameter can be used to force only passive scanning to be used
# for network connection cases. It should be noted that this will slow
# down scan operations and reduce likelihood of finding the AP. In
# addition, some use cases will override this due to functional
# requirements, e.g., for finding an AP that uses hidden SSID
# (scan_ssid=1) or P2P device discovery.
#
# 0:  Do normal scans (allow active scans) (default)
# 1:  Do passive scans.
#passive_scan=0

# MPM residency
# By default, wpa_supplicant implements the mesh peering manager (MPM) for an
# open mesh. However, if the driver can implement the MPM, you may set this to
# 0 to use the driver version. When AMPE is enabled, the wpa_supplicant MPM is
# always used.
# 0: MPM lives in the driver
# 1: wpa_supplicant provides an MPM which handles peering (default)
#user_mpm=1

# Maximum number of peer links (0-255; default: 99)
# Maximum number of mesh peering currently maintained by the STA.
#max_peer_links=99

# Timeout in seconds to detect STA inactivity (default: 300 seconds)
#
# This timeout value is used in mesh STA to clean up inactive stations.
#mesh_max_inactivity=300

# cert_in_cb - Whether to include a peer certificate dump in events
# This controls whether peer certificates for authentication server and
# its certificate chain are included in EAP peer certificate events. This is
# enabled by default.
#cert_in_cb=1

# EAP fast re-authentication
# By default, fast re-authentication is enabled for all EAP methods that
# support it. This variable can be used to disable fast re-authentication.
# Normally, there is no need to disable this.
fast_reauth=1

# OpenSSL Engine support
# These options can be used to load OpenSSL engines in special or legacy
# modes.
# The two engines that are supported currently are shown below:
# They are both from the opensc project (http://www.opensc.org/)
# By default the PKCS#11 engine is loaded if the client_cert or
# private_key option appear to be a PKCS#11 URI, and these options
# should not need to be used explicitly.
# make the opensc engine available
#opensc_engine_path=/usr/lib/opensc/engine_opensc.so
# make the pkcs11 engine available
#pkcs11_engine_path=/usr/lib/opensc/engine_pkcs11.so
# configure the path to the pkcs11 module required by the pkcs11 engine
#pkcs11_module_path=/usr/lib/pkcs11/opensc-pkcs11.so

# OpenSSL cipher string
#
# This is an OpenSSL specific configuration option for configuring the default
# ciphers. If not set, the value configured at build time ("DEFAULT:!EXP:!LOW"
# by default) is used.
# See https://www.openssl.org/docs/apps/ciphers.html for OpenSSL documentation
# on cipher suite configuration. This is applicable only if wpa_supplicant is
# built to use OpenSSL.
#openssl_ciphers=DEFAULT:!EXP:!LOW

# Dynamic EAP methods
# If EAP methods were built dynamically as shared object files, they need to be
# loaded here before being used in the network blocks. By default, EAP methods
# are included statically in the build, so these lines are not needed
#load_dynamic_eap=/usr/lib/wpa_supplicant/eap_tls.so
#load_dynamic_eap=/usr/lib/wpa_supplicant/eap_md5.so

# Driver interface parameters
# This field can be used to configure arbitrary driver interface parameters. The
# format is specific to the selected driver interface. This field is not used
# in most cases.
#driver_param="field=value"

# Country code
# The ISO/IEC alpha2 country code for the country in which this device is
# currently operating.
#country=US

# Maximum lifetime for PMKSA in seconds; default 43200
#dot11RSNAConfigPMKLifetime=43200
# Threshold for reauthentication (percentage of PMK lifetime); default 70
#dot11RSNAConfigPMKReauthThreshold=70
# Timeout for security association negotiation in seconds; default 60
#dot11RSNAConfigSATimeout=60

# Wi-Fi Protected Setup (WPS) parameters

# Universally Unique IDentifier (UUID; see RFC 4122) of the device
# If not configured, UUID will be generated based on the mechanism selected with
# the auto_uuid parameter.
#uuid=12345678-9abc-def0-1234-56789abcdef0

# Automatic UUID behavior
# 0 = generate static value based on the local MAC address (default)
# 1 = generate a random UUID every time wpa_supplicant starts
#auto_uuid=0

# Device Name
# User-friendly description of device; up to 32 octets encoded in UTF-8
#device_name=Wireless Client

# Manufacturer
# The manufacturer of the device (up to 64 ASCII characters)
#manufacturer=Company

# Model Name
# Model of the device (up to 32 ASCII characters)
#model_name=cmodel

# Model Number
# Additional device description (up to 32 ASCII characters)
#model_number=123

# Serial Number
# Serial number of the device (up to 32 characters)
#serial_number=12345

# Primary Device Type
# Used format: <categ>-<OUI>-<subcateg>
# categ = Category as an integer value
# OUI = OUI and type octet as a 4-octet hex-encoded value; 0050F204 for
#       default WPS OUI
# subcateg = OUI-specific Sub Category as an integer value
# Examples:
#   1-0050F204-1 (Computer / PC)
#   1-0050F204-2 (Computer / Server)
#   5-0050F204-1 (Storage / NAS)
#   6-0050F204-1 (Network Infrastructure / AP)
#device_type=1-0050F204-1

# OS Version
# 4-octet operating system version number (hex string)
#os_version=01020300

# Config Methods
# List of the supported configuration methods
# Available methods: usba ethernet label display ext_nfc_token int_nfc_token
#	nfc_interface push_button keypad virtual_display physical_display
#	virtual_push_button physical_push_button
# For WSC 1.0:
#config_methods=label display push_button keypad
# For WSC 2.0:
#config_methods=label virtual_display virtual_push_button keypad

# Credential processing
#   0 = process received credentials internally (default)
#   1 = do not process received credentials; just pass them over ctrl_iface to
#	external program(s)
#   2 = process received credentials internally and pass them over ctrl_iface
#	to external program(s)
#wps_cred_processing=0

# Vendor attribute in WPS M1, e.g., Windows 7 Vertical Pairing
# The vendor attribute contents to be added in M1 (hex string)
#wps_vendor_ext_m1=000137100100020001

# NFC password token for WPS
# These parameters can be used to configure a fixed NFC password token for the
# station. This can be generated, e.g., with nfc_pw_token. When these
# parameters are used, the station is assumed to be deployed with a NFC tag
# that includes the matching NFC password token (e.g., written based on the
# NDEF record from nfc_pw_token).
#
#wps_nfc_dev_pw_id: Device Password ID (16..65535)
#wps_nfc_dh_pubkey: Hexdump of DH Public Key
#wps_nfc_dh_privkey: Hexdump of DH Private Key
#wps_nfc_dev_pw: Hexdump of Device Password

# Priority for the networks added through WPS
# This priority value will be set to each network profile that is added
# by executing the WPS protocol.
#wps_priority=0

# Maximum number of BSS entries to keep in memory
# Default: 200
# This can be used to limit memory use on the BSS entries (cached scan
# results). A larger value may be needed in environments that have huge number
# of APs when using ap_scan=1 mode.
#bss_max_count=200

# Automatic scan
# This is an optional set of parameters for automatic scanning
# within an interface in following format:
#autoscan=<autoscan module name>:<module parameters>
# autoscan is like bgscan but on disconnected or inactive state.
# For instance, on exponential module parameters would be <base>:<limit>
#autoscan=exponential:3:300
# Which means a delay between scans on a base exponential of 3,
# up to the limit of 300 seconds (3, 9, 27 ... 300)
# For periodic module, parameters would be <fixed interval>
#autoscan=periodic:30
# So a delay of 30 seconds will be applied between each scan.
# Note: If sched_scan_plans are configured and supported by the driver,
# autoscan is ignored.

# filter_ssids - SSID-based scan result filtering
# 0 = do not filter scan results (default)
# 1 = only include configured SSIDs in scan results/BSS table
#filter_ssids=0

# Password (and passphrase, etc.) backend for external storage
# format: <backend name>[:<optional backend parameters>]
#ext_password_backend=test:pw1=password|pw2=testing


# Disable P2P functionality
# p2p_disabled=1

# Timeout in seconds to detect STA inactivity (default: 300 seconds)
#
# This timeout value is used in P2P GO mode to clean up
# inactive stations.
#p2p_go_max_inactivity=300

# Passphrase length (8..63) for P2P GO
#
# This parameter controls the length of the random passphrase that is
# generated at the GO. Default: 8.
#p2p_passphrase_len=8

# Extra delay between concurrent P2P search iterations
#
# This value adds extra delay in milliseconds between concurrent search
# iterations to make p2p_find friendlier to concurrent operations by avoiding
# it from taking 100% of radio resources. The default value is 500 ms.
#p2p_search_delay=500

# Opportunistic Key Caching (also known as Proactive Key Caching) default
# This parameter can be used to set the default behavior for the
# proactive_key_caching parameter. By default, OKC is disabled unless enabled
# with the global okc=1 parameter or with the per-network
# proactive_key_caching=1 parameter. With okc=1, OKC is enabled by default, but
# can be disabled with per-network proactive_key_caching=0 parameter.
#okc=0

# Protected Management Frames default
# This parameter can be used to set the default behavior for the ieee80211w
# parameter for RSN networks. By default, PMF is disabled unless enabled with
# the global pmf=1/2 parameter or with the per-network ieee80211w=1/2 parameter.
# With pmf=1/2, PMF is enabled/required by default, but can be disabled with the
# per-network ieee80211w parameter. This global default value does not apply
# for non-RSN networks (key_mgmt=NONE) since PMF is available only when using
# RSN.
#pmf=0

# Enabled SAE finite cyclic groups in preference order
# By default (if this parameter is not set), the mandatory group 19 (ECC group
# defined over a 256-bit prime order field) is preferred, but other groups are
# also enabled. If this parameter is set, the groups will be tried in the
# indicated order. The group values are listed in the IANA registry:
# http://www.iana.org/assignments/ipsec-registry/ipsec-registry.xml#ipsec-registry-9
#sae_groups=21 20 19 26 25

# Default value for DTIM period (if not overridden in network block)
#dtim_period=2

# Default value for Beacon interval (if not overridden in network block)
#beacon_int=100

# Additional vendor specific elements for Beacon and Probe Response frames
# This parameter can be used to add additional vendor specific element(s) into
# the end of the Beacon and Probe Response frames. The format for these
# element(s) is a hexdump of the raw information elements (id+len+payload for
# one or more elements). This is used in AP and P2P GO modes.
#ap_vendor_elements=dd0411223301

# Ignore scan results older than request
#
# The driver may have a cache of scan results that makes it return
# information that is older than our scan trigger. This parameter can
# be used to configure such old information to be ignored instead of
# allowing it to update the internal BSS table.
#ignore_old_scan_res=0

# scan_cur_freq: Whether to scan only the current frequency
# 0:  Scan all available frequencies. (Default)
# 1:  Scan current operating frequency if another VIF on the same radio
#     is already associated.

# MAC address policy default
# 0 = use permanent MAC address
# 1 = use random MAC address for each ESS connection
# 2 = like 1, but maintain OUI (with local admin bit set)
#
# By default, permanent MAC address is used unless policy is changed by
# the per-network mac_addr parameter. Global mac_addr=1 can be used to
# change this default behavior.
#mac_addr=0

# Lifetime of random MAC address in seconds (default: 60)
#rand_addr_lifetime=60

# MAC address policy for pre-association operations (scanning, ANQP)
# 0 = use permanent MAC address
# 1 = use random MAC address
# 2 = like 1, but maintain OUI (with local admin bit set)
#preassoc_mac_addr=0

# MAC address policy for GAS operations
# 0 = use permanent MAC address
# 1 = use random MAC address
# 2 = like 1, but maintain OUI (with local admin bit set)
#gas_rand_mac_addr=0

# Lifetime of GAS random MAC address in seconds (default: 60)
#gas_rand_addr_lifetime=60

# Interworking (IEEE 802.11u)

# Enable Interworking
# interworking=1

# Enable P2P GO advertisement of Interworking
# go_interworking=1

# P2P GO Interworking: Access Network Type
# 0 = Private network
# 1 = Private network with guest access
# 2 = Chargeable public network
# 3 = Free public network
# 4 = Personal device network
# 5 = Emergency services only network
# 14 = Test or experimental
# 15 = Wildcard
#go_access_network_type=0

# P2P GO Interworking: Whether the network provides connectivity to the Internet
# 0 = Unspecified
# 1 = Network provides connectivity to the Internet
#go_internet=1

# P2P GO Interworking: Group Venue Info (optional)
# The available values are defined in IEEE Std 802.11-2016, 9.4.1.35.
# Example values (group,type):
# 0,0 = Unspecified
# 1,7 = Convention Center
# 1,13 = Coffee Shop
# 2,0 = Unspecified Business
# 7,1  Private Residence
#go_venue_group=7
#go_venue_type=1

# Homogenous ESS identifier
# If this is set, scans will be used to request response only from BSSes
# belonging to the specified Homogeneous ESS. This is used only if interworking
# is enabled.
# hessid=00:11:22:33:44:55

# Automatic network selection behavior
# 0 = do not automatically go through Interworking network selection
#     (i.e., require explicit interworking_select command for this; default)
# 1 = perform Interworking network selection if one or more
#     credentials have been configured and scan did not find a
#     matching network block
#auto_interworking=0

# GAS Address3 field behavior
# 0 = P2P specification (Address3 = AP BSSID); default
# 1 = IEEE 802.11 standard compliant (Address3 = Wildcard BSSID when
#     sent to not-associated AP; if associated, AP BSSID)
#gas_address3=0

# Publish fine timing measurement (FTM) responder functionality in
# the Extended Capabilities element bit 70.
# Controls whether FTM responder functionality will be published by AP/STA.
# Note that actual FTM responder operation is managed outside wpa_supplicant.
# 0 = Do not publish; default
# 1 = Publish
#ftm_responder=0

# Publish fine timing measurement (FTM) initiator functionality in
# the Extended Capabilities element bit 71.
# Controls whether FTM initiator functionality will be published by AP/STA.
# Note that actual FTM initiator operation is managed outside wpa_supplicant.
# 0 = Do not publish; default
# 1 = Publish
#ftm_initiator=0

# credential block
#
# Each credential used for automatic network selection is configured as a set
# of parameters that are compared to the information advertised by the APs when
# interworking_select and interworking_connect commands are used.
#
# credential fields:
#
# temporary: Whether this credential is temporary and not to be saved
#
# priority: Priority group
#	By default, all networks and credentials get the same priority group
#	(0). This field can be used to give higher priority for credentials
#	(and similarly in struct wpa_ssid for network blocks) to change the
#	Interworking automatic networking selection behavior. The matching
#	network (based on either an enabled network block or a credential)
#	with the highest priority value will be selected.
#
# pcsc: Use PC/SC and SIM/USIM card
#
# realm: Home Realm for Interworking
#
# username: Username for Interworking network selection
#
# password: Password for Interworking network selection
#
# ca_cert: CA certificate for Interworking network selection
#
# client_cert: File path to client certificate file (PEM/DER)
#	This field is used with Interworking networking selection for a case
#	where client certificate/private key is used for authentication
#	(EAP-TLS). Full path to the file should be used since working
#	directory may change when wpa_supplicant is run in the background.
#
#	Certificates from PKCS#11 tokens can be referenced by a PKCS#11 URI.
#
#	For example: private_key="pkcs11:manufacturer=piv_II;id=%01"
#
#	Alternatively, a named configuration blob can be used by setting
#	this to blob://blob_name.
#
# private_key: File path to client private key file (PEM/DER/PFX)
#	When PKCS#12/PFX file (.p12/.pfx) is used, client_cert should be
#	commented out. Both the private key and certificate will be read
#	from the PKCS#12 file in this case. Full path to the file should be
#	used since working directory may change when wpa_supplicant is run
#	in the background.
#
#	Keys in PKCS#11 tokens can be referenced by a PKCS#11 URI.
#	For example: private_key="pkcs11:manufacturer=piv_II;id=%01"
#
#	Windows certificate store can be used by leaving client_cert out and
#	configuring private_key in one of the following formats:
#
#	cert://substring_to_match
#
#	hash://certificate_thumbprint_in_hex
#
#	For example: private_key="hash://63093aa9c47f56ae88334c7b65a4"
#
#	Note that when running wpa_supplicant as an application, the user
#	certificate store (My user account) is used, whereas computer store
#	(Computer account) is used when running wpasvc as a service.
#
#	Alternatively, a named configuration blob can be used by setting
#	this to blob://blob_name.
#
# private_key_passwd: Password for private key file
#
# imsi: IMSI in <MCC> | <MNC> | '-' | <MSIN> format
#
# milenage: Milenage parameters for SIM/USIM simulator in <Ki>:<OPc>:<SQN>
#	format
#
# domain: Home service provider FQDN(s)
#	This is used to compare against the Domain Name List to figure out
#	whether the AP is operated by the Home SP. Multiple domain entries can
#	be used to configure alternative FQDNs that will be considered home
#	networks.
#
# roaming_consortium: Roaming Consortium OI
#	If roaming_consortium_len is non-zero, this field contains the
#	Roaming Consortium OI that can be used to determine which access
#	points support authentication with this credential. This is an
#	alternative to the use of the realm parameter. When using Roaming
#	Consortium to match the network, the EAP parameters need to be
#	pre-configured with the credential since the NAI Realm information
#	may not be available or fetched.
#
# required_roaming_consortium: Required Roaming Consortium OI
#	If required_roaming_consortium_len is non-zero, this field contains the
#	Roaming Consortium OI that is required to be advertised by the AP for
#	the credential to be considered matching.
#
# roaming_consortiums: Roaming Consortium OI(s) memberships
#	This string field contains one or more comma delimited OIs (hexdump)
#	identifying the roaming consortiums of which the provider is a member.
#	The list is sorted from the most preferred one to the least preferred
#	one. A match between the Roaming Consortium OIs advertised by an AP and
#	the OIs in this list indicates that successful authentication is
#	possible.
#	(Hotspot 2.0 PerProviderSubscription/<X+>/HomeSP/RoamingConsortiumOI)
#
# eap: Pre-configured EAP method
#	This optional field can be used to specify which EAP method will be
#	used with this credential. If not set, the EAP method is selected
#	automatically based on ANQP information (e.g., NAI Realm).
#
# phase1: Pre-configure Phase 1 (outer authentication) parameters
#	This optional field is used with like the 'eap' parameter.
#
# phase2: Pre-configure Phase 2 (inner authentication) parameters
#	This optional field is used with like the 'eap' parameter.
#
# excluded_ssid: Excluded SSID
#	This optional field can be used to excluded specific SSID(s) from
#	matching with the network. Multiple entries can be used to specify more
#	than one SSID.
#
# roaming_partner: Roaming partner information
#	This optional field can be used to configure preferences between roaming
#	partners. The field is a string in following format:
#	<FQDN>,<0/1 exact match>,<priority>,<* or country code>
#	(non-exact match means any subdomain matches the entry; priority is in
#	0..255 range with 0 being the highest priority)
#
# update_identifier: PPS MO ID
#	(Hotspot 2.0 PerProviderSubscription/UpdateIdentifier)
#
# provisioning_sp: FQDN of the SP that provisioned the credential
#	This optional field can be used to keep track of the SP that provisioned
#	the credential to find the PPS MO (./Wi-Fi/<provisioning_sp>).
#
# Minimum backhaul threshold (PPS/<X+>/Policy/MinBackhauldThreshold/*)
#	These fields can be used to specify minimum download/upload backhaul
#	bandwidth that is preferred for the credential. This constraint is
#	ignored if the AP does not advertise WAN Metrics information or if the
#	limit would prevent any connection. Values are in kilobits per second.
# min_dl_bandwidth_home
# min_ul_bandwidth_home
# min_dl_bandwidth_roaming
# min_ul_bandwidth_roaming
#
# max_bss_load: Maximum BSS Load Channel Utilization (1..255)
#	(PPS/<X+>/Policy/MaximumBSSLoadValue)
#	This value is used as the maximum channel utilization for network
#	selection purposes for home networks. If the AP does not advertise
#	BSS Load or if the limit would prevent any connection, this constraint
#	will be ignored.
#
# req_conn_capab: Required connection capability
#	(PPS/<X+>/Policy/RequiredProtoPortTuple)
#	This value is used to configure set of required protocol/port pairs that
#	a roaming network shall support (include explicitly in Connection
#	Capability ANQP element). This constraint is ignored if the AP does not
#	advertise Connection Capability or if this constraint would prevent any
#	network connection. This policy is not used in home networks.
#	Format: <protocol>[:<comma-separated list of ports]
#	Multiple entries can be used to list multiple requirements.
#	For example, number of common TCP protocols:
#	req_conn_capab=6,22,80,443
#	For example, IPSec/IKE:
#	req_conn_capab=17:500
#	req_conn_capab=50
#
# ocsp: Whether to use/require OCSP to check server certificate
#	0 = do not use OCSP stapling (TLS certificate status extension)
#	1 = try to use OCSP stapling, but not require response
#	2 = require valid OCSP stapling response
#	3 = require valid OCSP stapling response for all not-trusted
#	    certificates in the server certificate chain
#
# sim_num: Identifier for which SIM to use in multi-SIM devices
#
# for example:
#
#cred={
#	realm="example.com"
#	username="user@example.com"
#	password="password"
#	ca_cert="/etc/wpa_supplicant/ca.pem"
#	domain="example.com"
#}
#
#cred={
#	imsi="310026-000000000"
#	milenage="90dca4eda45b53cf0f12d7c9c3bc6a89:cb9cccc4b9258e6dca4760379fb82"
#}
#
#cred={
#	realm="example.com"
#	username="user"
#	password="password"
#	ca_cert="/etc/wpa_supplicant/ca.pem"
#	domain="example.com"
#	roaming_consortium=223344
#	eap=TTLS
#	phase2="auth=MSCHAPV2"
#}

# Hotspot 2.0
# hs20=1

# Scheduled scan plans
#
# A space delimited list of scan plans. Each scan plan specifies the scan
# interval and number of iterations, delimited by a colon. The last scan plan
# will run infinitely and thus must specify only the interval and not the number
# of iterations.
#
# The driver advertises the maximum number of scan plans supported. If more scan
# plans than supported are configured, only the first ones are set (up to the
# maximum supported). The last scan plan that specifies only the interval is
# always set as the last plan.
#
# If the scan interval or the number of iterations for a scan plan exceeds the
# maximum supported, it will be set to the maximum supported value.
#
# Format:
# sched_scan_plans=<interval:iterations> <interval:iterations> ... <interval>
#
# Example:
# sched_scan_plans=10:100 20:200 30

# Multi Band Operation (MBO) non-preferred channels
# A space delimited list of non-preferred channels where each channel is a colon
# delimited list of values.
# Format:
# non_pref_chan=<oper_class>:<chan>:<preference>:<reason>
# Example:
# non_pref_chan=81:5:10:2 81:1:0:2 81:9:0:2

# MBO Cellular Data Capabilities
# 1 = Cellular data connection available
# 2 = Cellular data connection not available
# 3 = Not cellular capable (default)
#mbo_cell_capa=3

# Optimized Connectivity Experience (OCE)
# oce: Enable OCE features (bitmap)
# Set BIT(0) to Enable OCE in non-AP STA mode (default; disabled if the driver
#	does not indicate support for OCE in STA mode)
# Set BIT(1) to Enable OCE in STA-CFON mode
#oce=1

# network block
#
# Each network (usually AP's sharing the same SSID) is configured as a separate
# block in this configuration file. The network blocks are in preference order
# (the first match is used).
#
# network block fields:
#
# disabled:
#	0 = this network can be used (default)
#	1 = this network block is disabled (can be enabled through ctrl_iface,
#	    e.g., with wpa_cli or wpa_gui)
#
# id_str: Network identifier string for external scripts. This value is passed
#	to external action script through wpa_cli as WPA_ID_STR environment
#	variable to make it easier to do network specific configuration.
#
# ssid: SSID (mandatory); network name in one of the optional formats:
#	- an ASCII string with double quotation
#	- a hex string (two characters per octet of SSID)
#	- a printf-escaped ASCII string P"<escaped string>"
#
# scan_ssid:
#	0 = do not scan this SSID with specific Probe Request frames (default)
#	1 = scan with SSID-specific Probe Request frames (this can be used to
#	    find APs that do not accept broadcast SSID or use multiple SSIDs;
#	    this will add latency to scanning, so enable this only when needed)
#
# bssid: BSSID (optional); if set, this network block is used only when
#	associating with the AP using the configured BSSID
#
# priority: priority group (integer)
# By default, all networks will get same priority group (0). If some of the
# networks are more desirable, this field can be used to change the order in
# which wpa_supplicant goes through the networks when selecting a BSS. The
# priority groups will be iterated in decreasing priority (i.e., the larger the
# priority value, the sooner the network is matched against the scan results).
# Within each priority group, networks will be selected based on security
# policy, signal strength, etc.
# Please note that AP scanning with scan_ssid=1 and ap_scan=2 mode are not
# using this priority to select the order for scanning. Instead, they try the
# networks in the order that used in the configuration file.
#
# mode: IEEE 802.11 operation mode
# 0 = infrastructure (Managed) mode, i.e., associate with an AP (default)
# 1 = IBSS (ad-hoc, peer-to-peer)
# 2 = AP (access point)
# Note: IBSS can only be used with key_mgmt NONE (plaintext and static WEP) and
# WPA-PSK (with proto=RSN). In addition, key_mgmt=WPA-NONE (fixed group key
# TKIP/CCMP) is available for backwards compatibility, but its use is
# deprecated. WPA-None requires following network block options:
# proto=WPA, key_mgmt=WPA-NONE, pairwise=NONE, group=TKIP (or CCMP, but not
# both), and psk must also be set.
#
# frequency: Channel frequency in megahertz (MHz) for IBSS, e.g.,
# 2412 = IEEE 802.11b/g channel 1. This value is used to configure the initial
# channel for IBSS (adhoc) networks. It is ignored in the infrastructure mode.
# In addition, this value is only used by the station that creates the IBSS. If
# an IBSS network with the configured SSID is already present, the frequency of
# the network will be used instead of this configured value.
#
# pbss: Whether to use PBSS. Relevant to IEEE 802.11ad networks only.
# 0 = do not use PBSS
# 1 = use PBSS
# 2 = don't care (not allowed in AP mode)
# Used together with mode configuration. When mode is AP, it means to start a
# PCP instead of a regular AP. When mode is infrastructure it means connect
# to a PCP instead of AP. In this mode you can also specify 2 (don't care)
# which means connect to either PCP or AP.
# P2P_GO and P2P_GROUP_FORMATION modes must use PBSS in IEEE 802.11ad network.
# For more details, see IEEE Std 802.11ad-2012.
#
# scan_freq: List of frequencies to scan
# Space-separated list of frequencies in MHz to scan when searching for this
# BSS. If the subset of channels used by the network is known, this option can
# be used to optimize scanning to not occur on channels that the network does
# not use. Example: scan_freq=2412 2437 2462
#
# freq_list: Array of allowed frequencies
# Space-separated list of frequencies in MHz to allow for selecting the BSS. If
# set, scan results that do not match any of the specified frequencies are not
# considered when selecting a BSS.
#
# This can also be set on the outside of the network block. In this case,
# it limits the frequencies that will be scanned.
#
# bgscan: Background scanning
# wpa_supplicant behavior for background scanning can be specified by
# configuring a bgscan module. These modules are responsible for requesting
# background scans for the purpose of roaming within an ESS (i.e., within a
# single network block with all the APs using the same SSID). The bgscan
# parameter uses following format: "<bgscan module name>:<module parameters>"
# Following bgscan modules are available:
# simple - Periodic background scans based on signal strength
# bgscan="simple:<short bgscan interval in seconds>:<signal strength threshold>:
# <long interval>"
# bgscan="simple:30:-45:300"
# learn - Learn channels used by the network and try to avoid bgscans on other
# channels (experimental)
# bgscan="learn:<short bgscan interval in seconds>:<signal strength threshold>:
# <long interval>[:<database file name>]"
# bgscan="learn:30:-45:300:/etc/wpa_supplicant/network1.bgscan"
# Explicitly disable bgscan by setting
# bgscan=""
#
# This option can also be set outside of all network blocks for the bgscan
# parameter to apply for all the networks that have no specific bgscan
# parameter.
#
# proto: list of accepted protocols
# WPA = WPA/IEEE 802.11i/D3.0
# RSN = WPA2/IEEE 802.11i (also WPA2 can be used as an alias for RSN)
# Note that RSN is used also for WPA3.
# If not set, this defaults to: WPA RSN
#
# key_mgmt: list of accepted authenticated key management protocols
# WPA-PSK = WPA pre-shared key (this requires 'psk' field)
# WPA-EAP = WPA using EAP authentication
# IEEE8021X = IEEE 802.1X using EAP authentication and (optionally) dynamically
#	generated WEP keys
# NONE = WPA is not used; plaintext or static WEP could be used
# WPA-NONE = WPA-None for IBSS (deprecated; use proto=RSN key_mgmt=WPA-PSK
#	instead)
# FT-PSK = Fast BSS Transition (IEEE 802.11r) with pre-shared key
# FT-EAP = Fast BSS Transition (IEEE 802.11r) with EAP authentication
# FT-EAP-SHA384 = Fast BSS Transition (IEEE 802.11r) with EAP authentication
#	and using SHA384
# WPA-PSK-SHA256 = Like WPA-PSK but using stronger SHA256-based algorithms
# WPA-EAP-SHA256 = Like WPA-EAP but using stronger SHA256-based algorithms
# SAE = Simultaneous authentication of equals; pre-shared key/password -based
#	authentication with stronger security than WPA-PSK especially when using
#	not that strong password; a.k.a. WPA3-Personal
# FT-SAE = SAE with FT
# WPA-EAP-SUITE-B = Suite B 128-bit level
# WPA-EAP-SUITE-B-192 = Suite B 192-bit level
# OSEN = Hotspot 2.0 Rel 2 online signup connection
# FILS-SHA256 = Fast Initial Link Setup with SHA256
# FILS-SHA384 = Fast Initial Link Setup with SHA384
# FT-FILS-SHA256 = FT and Fast Initial Link Setup with SHA256
# FT-FILS-SHA384 = FT and Fast Initial Link Setup with SHA384
# OWE = Opportunistic Wireless Encryption (a.k.a. Enhanced Open)
# DPP = Device Provisioning Protocol
# If not set, this defaults to: WPA-PSK WPA-EAP
#
# ieee80211w: whether management frame protection is enabled
# 0 = disabled (default unless changed with the global pmf parameter)
# 1 = optional
# 2 = required
# The most common configuration options for this based on the PMF (protected
# management frames) certification program are:
# PMF enabled: ieee80211w=1 and key_mgmt=WPA-EAP WPA-EAP-SHA256
# PMF required: ieee80211w=2 and key_mgmt=WPA-EAP-SHA256
# (and similarly for WPA-PSK and WPA-WPSK-SHA256 if WPA2-Personal is used)
#
# auth_alg: list of allowed IEEE 802.11 authentication algorithms
# OPEN = Open System authentication (required for WPA/WPA2)
# SHARED = Shared Key authentication (requires static WEP keys)
# LEAP = LEAP/Network EAP (only used with LEAP)
# If not set, automatic selection is used (Open System with LEAP enabled if
# LEAP is allowed as one of the EAP methods).
#
# pairwise: list of accepted pairwise (unicast) ciphers for WPA
# CCMP = AES in Counter mode with CBC-MAC [RFC 3610, IEEE 802.11i/D7.0]
# TKIP = Temporal Key Integrity Protocol [IEEE 802.11i/D7.0]
# NONE = Use only Group Keys (deprecated, should not be included if APs support
#	pairwise keys)
# If not set, this defaults to: CCMP TKIP
#
# group: list of accepted group (broadcast/multicast) ciphers for WPA
# CCMP = AES in Counter mode with CBC-MAC [RFC 3610, IEEE 802.11i/D7.0]
# TKIP = Temporal Key Integrity Protocol [IEEE 802.11i/D7.0]
# WEP104 = WEP (Wired Equivalent Privacy) with 104-bit key
# WEP40 = WEP (Wired Equivalent Privacy) with 40-bit key [IEEE 802.11]
# If not set, this defaults to: CCMP TKIP WEP104 WEP40
#
# group_mgmt: list of accepted group management ciphers for RSN (PMF)
# AES-128-CMAC = BIP-CMAC-128
# BIP-GMAC-128
# BIP-GMAC-256
# BIP-CMAC-256
# If not set, no constraint on the cipher, i.e., accept whichever cipher the AP
# indicates.
#
# psk: WPA preshared key; 256-bit pre-shared key
# The key used in WPA-PSK mode can be entered either as 64 hex-digits, i.e.,
# 32 bytes or as an ASCII passphrase (in which case, the real PSK will be
# generated using the passphrase and SSID). ASCII passphrase must be between
# 8 and 63 characters (inclusive). ext:<name of external PSK field> format can
# be used to indicate that the PSK/passphrase is stored in external storage.
# This field is not needed, if WPA-EAP is used.
# Note: Separate tool, wpa_passphrase, can be used to generate 256-bit keys
# from ASCII passphrase. This process uses lot of CPU and wpa_supplicant
# startup and reconfiguration time can be optimized by generating the PSK only
# only when the passphrase or SSID has actually changed.
#
# mem_only_psk: Whether to keep PSK/passphrase only in memory
# 0 = allow psk/passphrase to be stored to the configuration file
# 1 = do not store psk/passphrase to the configuration file
#mem_only_psk=0
#
# sae_password: SAE password
# This parameter can be used to set a password for SAE. By default, the
# passphrase from the psk parameter is used if this separate parameter is not
# used, but psk follows the WPA-PSK constraints (8..63 characters) even though
# SAE passwords do not have such constraints.
#
# sae_password_id: SAE password identifier
# This parameter can be used to set an identifier for the SAE password. By
# default, no such identifier is used. If set, the specified identifier value
# is used by the other peer to select which password to use for authentication.
#
# eapol_flags: IEEE 802.1X/EAPOL options (bit field)
# Dynamic WEP key required for non-WPA mode
# bit0 (1): require dynamically generated unicast WEP key
# bit1 (2): require dynamically generated broadcast WEP key
# 	(3 = require both keys; default)
# Note: When using wired authentication (including MACsec drivers),
# eapol_flags must be set to 0 for the authentication to be completed
# successfully.
#
# macsec_policy: IEEE 802.1X/MACsec options
# This determines how sessions are secured with MACsec (only for MACsec
# drivers).
# 0: MACsec not in use (default)
# 1: MACsec enabled - Should secure, accept key server's advice to
#    determine whether to use a secure session or not.
#
# macsec_integ_only: IEEE 802.1X/MACsec transmit mode
# This setting applies only when MACsec is in use, i.e.,
#  - macsec_policy is enabled
#  - the key server has decided to enable MACsec
# 0: Encrypt traffic (default)
# 1: Integrity only
#
# macsec_port: IEEE 802.1X/MACsec port
# Port component of the SCI
# Range: 1-65534 (default: 1)
#
# mka_cak, mka_ckn, and mka_priority: IEEE 802.1X/MACsec pre-shared key mode
# This allows to configure MACsec with a pre-shared key using a (CAK,CKN) pair.
# In this mode, instances of wpa_supplicant can act as MACsec peers. The peer
# with lower priority will become the key server and start distributing SAKs.
# mka_cak (CAK = Secure Connectivity Association Key) takes a 16-bytes (128 bit)
# hex-string (32 hex-digits)
# mka_ckn (CKN = CAK Name) takes a 32-bytes (256 bit) hex-string (64 hex-digits)
# mka_priority (Priority of MKA Actor) is in 0..255 range with 255 being
# default priority
#
# mixed_cell: This option can be used to configure whether so called mixed
# cells, i.e., networks that use both plaintext and encryption in the same
# SSID, are allowed when selecting a BSS from scan results.
# 0 = disabled (default)
# 1 = enabled
#
# proactive_key_caching:
# Enable/disable opportunistic PMKSA caching for WPA2.
# 0 = disabled (default unless changed with the global okc parameter)
# 1 = enabled
#
# wep_key0..3: Static WEP key (ASCII in double quotation, e.g. "abcde" or
# hex without quotation, e.g., 0102030405)
# wep_tx_keyidx: Default WEP key index (TX) (0..3)
#
# wpa_ptk_rekey: Maximum lifetime for PTK in seconds. This can be used to
# enforce rekeying of PTK to mitigate some attacks against TKIP deficiencies.
#
# group_rekey: Group rekeying time in seconds. This value, if non-zero, is used
# as the dot11RSNAConfigGroupRekeyTime parameter when operating in
# Authenticator role in IBSS, or in AP and mesh modes.
#
# Following fields are only used with internal EAP implementation.
# eap: space-separated list of accepted EAP methods
#	MD5 = EAP-MD5 (insecure and does not generate keying material ->
#			cannot be used with WPA; to be used as a Phase 2 method
#			with EAP-PEAP or EAP-TTLS)
#       MSCHAPV2 = EAP-MSCHAPv2 (cannot be used separately with WPA; to be used
#		as a Phase 2 method with EAP-PEAP or EAP-TTLS)
#       OTP = EAP-OTP (cannot be used separately with WPA; to be used
#		as a Phase 2 method with EAP-PEAP or EAP-TTLS)
#       GTC = EAP-GTC (cannot be used separately with WPA; to be used
#		as a Phase 2 method with EAP-PEAP or EAP-TTLS)
#	TLS = EAP-TLS (client and server certificate)
#	PEAP = EAP-PEAP (with tunnelled EAP authentication)
#	TTLS = EAP-TTLS (with tunnelled EAP or PAP/CHAP/MSCHAP/MSCHAPV2
#			 authentication)
#	If not set, all compiled in methods are allowed.
#
# identity: Identity string for EAP
#	This field is also used to configure user NAI for
#	EAP-PSK/PAX/SAKE/GPSK.
# anonymous_identity: Anonymous identity string for EAP (to be used as the
#	unencrypted identity with EAP types that support different tunnelled
#	identity, e.g., EAP-TTLS). This field can also be used with
#	EAP-SIM/AKA/AKA' to store the pseudonym identity.
# password: Password string for EAP. This field can include either the
#	plaintext password (using ASCII or hex string) or a NtPasswordHash
#	(16-byte MD4 hash of password) in hash:<32 hex digits> format.
#	NtPasswordHash can only be used when the password is for MSCHAPv2 or
#	MSCHAP (EAP-MSCHAPv2, EAP-TTLS/MSCHAPv2, EAP-TTLS/MSCHAP, LEAP).
#	EAP-PSK (128-bit PSK), EAP-PAX (128-bit PSK), and EAP-SAKE (256-bit
#	PSK) is also configured using this field. For EAP-GPSK, this is a
#	variable length PSK. ext:<name of external password field> format can
#	be used to indicate that the password is stored in external storage.
# ca_cert: File path to CA certificate file (PEM/DER). This file can have one
#	or more trusted CA certificates. If ca_cert and ca_path are not
#	included, server certificate will not be verified. This is insecure and
#	a trusted CA certificate should always be configured when using
#	EAP-TLS/TTLS/PEAP. Full path should be used since working directory may
#	change when wpa_supplicant is run in the background.
#
#	Alternatively, this can be used to only perform matching of the server
#	certificate (SHA-256 hash of the DER encoded X.509 certificate). In
#	this case, the possible CA certificates in the server certificate chain
#	are ignored and only the server certificate is verified. This is
#	configured with the following format:
#	hash:://server/sha256/cert_hash_in_hex
#	For example: "hash://server/sha256/
#	5a1bc1296205e6fdbe3979728efe3920798885c1c4590b5f90f43222d239ca6a"
#
#	On Windows, trusted CA certificates can be loaded from the system
#	certificate store by setting this to cert_store://<name>, e.g.,
#	ca_cert="cert_store://CA" or ca_cert="cert_store://ROOT".
#	Note that when running wpa_supplicant as an application, the user
#	certificate store (My user account) is used, whereas computer store
#	(Computer account) is used when running wpasvc as a service.
# ca_path: Directory path for CA certificate files (PEM). This path may
#	contain multiple CA certificates in OpenSSL format. Common use for this
#	is to point to system trusted CA list which is often installed into
#	directory like /etc/ssl/certs. If configured, these certificates are
#	added to the list of trusted CAs. ca_cert may also be included in that
#	case, but it is not required.
# client_cert: File path to client certificate file (PEM/DER)
#	Full path should be used since working directory may change when
#	wpa_supplicant is run in the background.
#	Alternatively, a named configuration blob can be used by setting this
#	to blob://<blob name>.
# private_key: File path to client private key file (PEM/DER/PFX)
#	When PKCS#12/PFX file (.p12/.pfx) is used, client_cert should be
#	commented out. Both the private key and certificate will be read from
#	the PKCS#12 file in this case. Full path should be used since working
#	directory may change when wpa_supplicant is run in the background.
#	Windows certificate store can be used by leaving client_cert out and
#	configuring private_key in one of the following formats:
#	cert://substring_to_match
#	hash://certificate_thumbprint_in_hex
#	for example: private_key="hash://63093aa9c47f56ae88334c7b65a4"
#	Note that when running wpa_supplicant as an application, the user
#	certificate store (My user account) is used, whereas computer store
#	(Computer account) is used when running wpasvc as a service.
#	Alternatively, a named configuration blob can be used by setting this
#	to blob://<blob name>.
# private_key_passwd: Password for private key file (if left out, this will be
#	asked through control interface)
# dh_file: File path to DH/DSA parameters file (in PEM format)
#	This is an optional configuration file for setting parameters for an
#	ephemeral DH key exchange. In most cases, the default RSA
#	authentication does not use this configuration. However, it is possible
#	setup RSA to use ephemeral DH key exchange. In addition, ciphers with
#	DSA keys always use ephemeral DH keys. This can be used to achieve
#	forward secrecy. If the file is in DSA parameters format, it will be
#	automatically converted into DH params.
# subject_match: Substring to be matched against the subject of the
#	authentication server certificate. If this string is set, the server
#	certificate is only accepted if it contains this string in the subject.
#	The subject string is in following format:
#	/C=US/ST=CA/L=San Francisco/CN=Test AS/emailAddress=as@example.com
#	Note: Since this is a substring match, this cannot be used securely to
#	do a suffix match against a possible domain name in the CN entry. For
#	such a use case, domain_suffix_match or domain_match should be used
#	instead.
# altsubject_match: Semicolon separated string of entries to be matched against
#	the alternative subject name of the authentication server certificate.
#	If this string is set, the server certificate is only accepted if it
#	contains one of the entries in an alternative subject name extension.
#	altSubjectName string is in following format: TYPE:VALUE
#	Example: EMAIL:server@example.com
#	Example: DNS:server.example.com;DNS:server2.example.com
#	Following types are supported: EMAIL, DNS, URI
# domain_suffix_match: Constraint for server domain name. If set, this FQDN is
#	used as a suffix match requirement for the AAA server certificate in
#	SubjectAltName dNSName element(s). If a matching dNSName is found, this
#	constraint is met. If no dNSName values are present, this constraint is
#	matched against SubjectName CN using same suffix match comparison.
#
#	Suffix match here means that the host/domain name is compared one label
#	at a time starting from the top-level domain and all the labels in
#	domain_suffix_match shall be included in the certificate. The
#	certificate may include additional sub-level labels in addition to the
#	required labels.
#
#	For example, domain_suffix_match=example.com would match
#	test.example.com but would not match test-example.com.
# domain_match: Constraint for server domain name
#	If set, this FQDN is used as a full match requirement for the
#	server certificate in SubjectAltName dNSName element(s). If a
#	matching dNSName is found, this constraint is met. If no dNSName
#	values are present, this constraint is matched against SubjectName CN
#	using same full match comparison. This behavior is similar to
#	domain_suffix_match, but has the requirement of a full match, i.e.,
#	no subdomains or wildcard matches are allowed. Case-insensitive
#	comparison is used, so "Example.com" matches "example.com", but would
#	not match "test.Example.com".
# phase1: Phase1 (outer authentication, i.e., TLS tunnel) parameters
#	(string with field-value pairs, e.g., "peapver=0" or
#	"peapver=1 peaplabel=1")
#	'peapver' can be used to force which PEAP version (0 or 1) is used.
#	'peaplabel=1' can be used to force new label, "client PEAP encryption",
#	to be used during key derivation when PEAPv1 or newer. Most existing
#	PEAPv1 implementation seem to be using the old label, "client EAP
#	encryption", and wpa_supplicant is now using that as the default value.
#	Some servers, e.g., Radiator, may require peaplabel=1 configuration to
#	interoperate with PEAPv1; see eap_testing.txt for more details.
#	'peap_outer_success=0' can be used to terminate PEAP authentication on
#	tunneled EAP-Success. This is required with some RADIUS servers that
#	implement draft-josefsson-pppext-eap-tls-eap-05.txt (e.g.,
#	Lucent NavisRadius v4.4.0 with PEAP in "IETF Draft 5" mode)
#	include_tls_length=1 can be used to force wpa_supplicant to include
#	TLS Message Length field in all TLS messages even if they are not
#	fragmented.
#	sim_min_num_chal=3 can be used to configure EAP-SIM to require three
#	challenges (by default, it accepts 2 or 3)
#	result_ind=1 can be used to enable EAP-SIM and EAP-AKA to use
#	protected result indication.
#	'crypto_binding' option can be used to control PEAPv0 cryptobinding
#	behavior:
#	 * 0 = do not use cryptobinding (default)
#	 * 1 = use cryptobinding if server supports it
#	 * 2 = require cryptobinding
#	EAP-WSC (WPS) uses following options: pin=<Device Password> or
#	pbc=1.
#
#	For wired IEEE 802.1X authentication, "allow_canned_success=1" can be
#	used to configure a mode that allows EAP-Success (and EAP-Failure)
#	without going through authentication step. Some switches use such
#	sequence when forcing the port to be authorized/unauthorized or as a
#	fallback option if the authentication server is unreachable. By default,
#	wpa_supplicant discards such frames to protect against potential attacks
#	by rogue devices, but this option can be used to disable that protection
#	for cases where the server/authenticator does not need to be
#	authenticated.
# phase2: Phase2 (inner authentication with TLS tunnel) parameters
#	(string with field-value pairs, e.g., "auth=MSCHAPV2" for EAP-PEAP or
#	"autheap=MSCHAPV2 autheap=MD5" for EAP-TTLS). "mschapv2_retry=0" can be
#	used to disable MSCHAPv2 password retry in authentication failure cases.
#
# TLS-based methods can use the following parameters to control TLS behavior
# (these are normally in the phase1 parameter, but can be used also in the
# phase2 parameter when EAP-TLS is used within the inner tunnel):
# tls_allow_md5=1 - allow MD5-based certificate signatures (depending on the
#	TLS library, these may be disabled by default to enforce stronger
#	security)
# tls_disable_time_checks=1 - ignore certificate validity time (this requests
#	the TLS library to accept certificates even if they are not currently
#	valid, i.e., have expired or have not yet become valid; this should be
#	used only for testing purposes)
# tls_disable_session_ticket=1 - disable TLS Session Ticket extension
# tls_disable_session_ticket=0 - allow TLS Session Ticket extension to be used
#	Note: If not set, this is automatically set to 1 for EAP-TLS/PEAP/TTLS
#	as a workaround for broken authentication server implementations unless
#	EAP workarounds are disabled with eap_workaround=0.
#	For EAP-FAST, this must be set to 0 (or left unconfigured for the
#	default value to be used automatically).
# tls_disable_tlsv1_0=1 - disable use of TLSv1.0
# tls_disable_tlsv1_1=1 - disable use of TLSv1.1 (a workaround for AAA servers
#	that have issues interoperating with updated TLS version)
# tls_disable_tlsv1_2=1 - disable use of TLSv1.2 (a workaround for AAA servers
#	that have issues interoperating with updated TLS version)
# tls_disable_tlsv1_3=1 - disable use of TLSv1.3 (a workaround for AAA servers
#	that have issues interoperating with updated TLS version)
# tls_ext_cert_check=0 - No external server certificate validation (default)
# tls_ext_cert_check=1 - External server certificate validation enabled; this
#	requires an external program doing validation of server certificate
#	chain when receiving CTRL-RSP-EXT_CERT_CHECK event from the control
#	interface and report the result of the validation with
#	CTRL-RSP_EXT_CERT_CHECK.
# tls_suiteb=0 - do not apply Suite B 192-bit constraints on TLS (default)
# tls_suiteb=1 - apply Suite B 192-bit constraints on TLS; this is used in
#	particular when using Suite B with RSA keys of >= 3K (3072) bits
#
# Following certificate/private key fields are used in inner Phase2
# authentication when using EAP-TTLS or EAP-PEAP.
# ca_cert2: File path to CA certificate file. This file can have one or more
#	trusted CA certificates. If ca_cert2 and ca_path2 are not included,
#	server certificate will not be verified. This is insecure and a trusted
#	CA certificate should always be configured.
# ca_path2: Directory path for CA certificate files (PEM)
# client_cert2: File path to client certificate file
# private_key2: File path to client private key file
# private_key2_passwd: Password for private key file
# dh_file2: File path to DH/DSA parameters file (in PEM format)
# subject_match2: Substring to be matched against the subject of the
#	authentication server certificate. See subject_match for more details.
# altsubject_match2: Semicolon separated string of entries to be matched
#	against the alternative subject name of the authentication server
#	certificate. See altsubject_match documentation for more details.
# domain_suffix_match2: Constraint for server domain name. See
#	domain_suffix_match for more details.
#
# fragment_size: Maximum EAP fragment size in bytes (default 1398).
#	This value limits the fragment size for EAP methods that support
#	fragmentation (e.g., EAP-TLS and EAP-PEAP). This value should be set
#	small enough to make the EAP messages fit in MTU of the network
#	interface used for EAPOL. The default value is suitable for most
#	cases.
#
# ocsp: Whether to use/require OCSP to check server certificate
#	0 = do not use OCSP stapling (TLS certificate status extension)
#	1 = try to use OCSP stapling, but not require response
#	2 = require valid OCSP stapling response
#	3 = require valid OCSP stapling response for all not-trusted
#	    certificates in the server certificate chain
#
# openssl_ciphers: OpenSSL specific cipher configuration
#	This can be used to override the global openssl_ciphers configuration
#	parameter (see above).
#
# erp: Whether EAP Re-authentication Protocol (ERP) is enabled
#
# EAP-FAST variables:
# pac_file: File path for the PAC entries. wpa_supplicant will need to be able
#	to create this file and write updates to it when PAC is being
#	provisioned or refreshed. Full path to the file should be used since
#	working directory may change when wpa_supplicant is run in the
#	background. Alternatively, a named configuration blob can be used by
#	setting this to blob://<blob name>
# phase1: fast_provisioning option can be used to enable in-line provisioning
#         of EAP-FAST credentials (PAC):
#         0 = disabled,
#         1 = allow unauthenticated provisioning,
#         2 = allow authenticated provisioning,
#         3 = allow both unauthenticated and authenticated provisioning
#	fast_max_pac_list_len=<num> option can be used to set the maximum
#		number of PAC entries to store in a PAC list (default: 10)
#	fast_pac_format=binary option can be used to select binary format for
#		storing PAC entries in order to save some space (the default
#		text format uses about 2.5 times the size of minimal binary
#		format)
#
# wpa_supplicant supports number of "EAP workarounds" to work around
# interoperability issues with incorrectly behaving authentication servers.
# These are enabled by default because some of the issues are present in large
# number of authentication servers. Strict EAP conformance mode can be
# configured by disabling workarounds with eap_workaround=0.

# update_identifier: PPS MO ID
#	(Hotspot 2.0 PerProviderSubscription/UpdateIdentifier)
#
# roaming_consortium_selection: Roaming Consortium Selection
#	The matching Roaming Consortium OI that was used to generate this
#	network profile.

# Station inactivity limit
#
# If a station does not send anything in ap_max_inactivity seconds, an
# empty data frame is sent to it in order to verify whether it is
# still in range. If this frame is not ACKed, the station will be
# disassociated and then deauthenticated. This feature is used to
# clear station table of old entries when the STAs move out of the
# range.
#
# The station can associate again with the AP if it is still in range;
# this inactivity poll is just used as a nicer way of verifying
# inactivity; i.e., client will not report broken connection because
# disassociation frame is not sent immediately without first polling
# the STA with a data frame.
# default: 300 (i.e., 5 minutes)
#ap_max_inactivity=300

# DTIM period in Beacon intervals for AP mode (default: 2)
#dtim_period=2

# Beacon interval (default: 100 TU)
#beacon_int=100

# WPS in AP mode
# 0 = WPS enabled and configured (default)
# 1 = WPS disabled
#wps_disabled=0

# FILS DH Group
# 0 = PFS disabled with FILS shared key authentication (default)
# 1-65535 = DH Group to use for FILS PFS
#fils_dh_group=0

# MAC address policy
# 0 = use permanent MAC address
# 1 = use random MAC address for each ESS connection
# 2 = like 1, but maintain OUI (with local admin bit set)
#mac_addr=0

# disable_ht: Whether HT (802.11n) should be disabled.
# 0 = HT enabled (if AP supports it)
# 1 = HT disabled
#
# disable_ht40: Whether HT-40 (802.11n) should be disabled.
# 0 = HT-40 enabled (if AP supports it)
# 1 = HT-40 disabled
#
# disable_sgi: Whether SGI (short guard interval) should be disabled.
# 0 = SGI enabled (if AP supports it)
# 1 = SGI disabled
#
# disable_ldpc: Whether LDPC should be disabled.
# 0 = LDPC enabled (if AP supports it)
# 1 = LDPC disabled
#
# ht40_intolerant: Whether 40 MHz intolerant should be indicated.
# 0 = 40 MHz tolerant (default)
# 1 = 40 MHz intolerant
#
# ht_mcs:  Configure allowed MCS rates.
#  Parsed as an array of bytes, in base-16 (ascii-hex)
# ht_mcs=""                                   // Use all available (default)
# ht_mcs="0xff 00 00 00 00 00 00 00 00 00 "   // Use MCS 0-7 only
# ht_mcs="0xff ff 00 00 00 00 00 00 00 00 "   // Use MCS 0-15 only
#
# disable_max_amsdu:  Whether MAX_AMSDU should be disabled.
# -1 = Do not make any changes.
# 0  = Enable MAX-AMSDU if hardware supports it.
# 1  = Disable AMSDU
#
# ampdu_factor: Maximum A-MPDU Length Exponent
# Value: 0-3, see 7.3.2.56.3 in IEEE Std 802.11n-2009.
#
# ampdu_density:  Allow overriding AMPDU density configuration.
#  Treated as hint by the kernel.
# -1 = Do not make any changes.
# 0-3 = Set AMPDU density (aka factor) to specified value.

# disable_vht: Whether VHT should be disabled.
# 0 = VHT enabled (if AP supports it)
# 1 = VHT disabled
#
# vht_capa: VHT capabilities to set in the override
# vht_capa_mask: mask of VHT capabilities
#
# vht_rx_mcs_nss_1/2/3/4/5/6/7/8: override the MCS set for RX NSS 1-8
# vht_tx_mcs_nss_1/2/3/4/5/6/7/8: override the MCS set for TX NSS 1-8
#  0: MCS 0-7
#  1: MCS 0-8
#  2: MCS 0-9
#  3: not supported

##### Fast Session Transfer (FST) support #####################################
#
# The options in this section are only available when the build configuration
# option CONFIG_FST is set while compiling wpa_supplicant. They allow this
# interface to be a part of FST setup.
#
# FST is the transfer of a session from a channel to another channel, in the
# same or different frequency bands.
#
# For details, see IEEE Std 802.11ad-2012.

# Identifier of an FST Group  the interface belongs to.
#fst_group_id=bond0

# Interface priority within the FST Group.
# Announcing a higher priority for an interface means declaring it more
# preferable for FST switch.
# fst_priority is in 1..255 range with 1 being the lowest priority.
#fst_priority=100

# Default LLT value for this interface in milliseconds. The value used in case
# no value provided during session setup. Default is 50 msec.
# fst_llt is in 1..4294967 range (due to spec limitation, see 10.32.2.2
# Transitioning between states).
#fst_llt=100

# Example blocks:

# Simple case: WPA-PSK, PSK as an ASCII passphrase, allow all valid ciphers
network={
	ssid="simple"
	psk="very secret passphrase"
	priority=5
}

# Same as previous, but request SSID-specific scanning (for APs that reject
# broadcast SSID)
network={
	ssid="second ssid"
	scan_ssid=1
	psk="very secret passphrase"
	priority=2
}

# Only WPA-PSK is used. Any valid cipher combination is accepted.
network={
	ssid="example"
	proto=WPA
	key_mgmt=WPA-PSK
	pairwise=CCMP TKIP
	group=CCMP TKIP WEP104 WEP40
	psk=06b4be19da289f475aa46a33cb793029d4ab3db7a23ee92382eb0106c72ac7bb
	priority=2
}

# WPA-Personal(PSK) with TKIP and enforcement for frequent PTK rekeying
network={
	ssid="example"
	proto=WPA
	key_mgmt=WPA-PSK
	pairwise=TKIP
	group=TKIP
	psk="not so secure passphrase"
	wpa_ptk_rekey=600
}

# Only WPA-EAP is used. Both CCMP and TKIP is accepted. An AP that used WEP104
# or WEP40 as the group cipher will not be accepted.
network={
	ssid="example"
	proto=RSN
	key_mgmt=WPA-EAP
	pairwise=CCMP TKIP
	group=CCMP TKIP
	eap=TLS
	identity="user@example.com"
	ca_cert="/etc/cert/ca.pem"
	client_cert="/etc/cert/user.pem"
	private_key="/etc/cert/user.prv"
	private_key_passwd="password"
	priority=1
}

# EAP-PEAP/MSCHAPv2 configuration for RADIUS servers that use the new peaplabel
# (e.g., Radiator)
network={
	ssid="example"
	key_mgmt=WPA-EAP
	eap=PEAP
	identity="user@example.com"
	password="foobar"
	ca_cert="/etc/cert/ca.pem"
	phase1="peaplabel=1"
	phase2="auth=MSCHAPV2"
	priority=10
}

# EAP-TTLS/EAP-MD5-Challenge configuration with anonymous identity for the
# unencrypted use. Real identity is sent only within an encrypted TLS tunnel.
network={
	ssid="example"
	key_mgmt=WPA-EAP
	eap=TTLS
	identity="user@example.com"
	anonymous_identity="anonymous@example.com"
	password="foobar"
	ca_cert="/etc/cert/ca.pem"
	priority=2
}

# EAP-TTLS/MSCHAPv2 configuration with anonymous identity for the unencrypted
# use. Real identity is sent only within an encrypted TLS tunnel.
network={
	ssid="example"
	key_mgmt=WPA-EAP
	eap=TTLS
	identity="user@example.com"
	anonymous_identity="anonymous@example.com"
	password="foobar"
	ca_cert="/etc/cert/ca.pem"
	phase2="auth=MSCHAPV2"
}

# WPA-EAP, EAP-TTLS with different CA certificate used for outer and inner
# authentication.
network={
	ssid="example"
	key_mgmt=WPA-EAP
	eap=TTLS
	# Phase1 / outer authentication
	anonymous_identity="anonymous@example.com"
	ca_cert="/etc/cert/ca.pem"
	# Phase 2 / inner authentication
	phase2="autheap=TLS"
	ca_cert2="/etc/cert/ca2.pem"
	client_cert2="/etc/cer/user.pem"
	private_key2="/etc/cer/user.prv"
	private_key2_passwd="password"
	priority=2
}

# Both WPA-PSK and WPA-EAP is accepted. Only CCMP is accepted as pairwise and
# group cipher.
network={
	ssid="example"
	bssid=00:11:22:33:44:55
	proto=WPA RSN
	key_mgmt=WPA-PSK WPA-EAP
	pairwise=CCMP
	group=CCMP
	psk=06b4be19da289f475aa46a33cb793029d4ab3db7a23ee92382eb0106c72ac7bb
}

# Special characters in SSID, so use hex string. Default to WPA-PSK, WPA-EAP
# and all valid ciphers.
network={
	ssid=00010203
	psk=000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f
}


# EAP-SIM with a GSM SIM or USIM
network={
	ssid="eap-sim-test"
	key_mgmt=WPA-EAP
	eap=SIM
	pin="1234"
	pcsc=""
}


# EAP-PSK
network={
	ssid="eap-psk-test"
	key_mgmt=WPA-EAP
	eap=PSK
	anonymous_identity="eap_psk_user"
	password=06b4be19da289f475aa46a33cb793029
	identity="eap_psk_user@example.com"
}


# IEEE 802.1X/EAPOL with dynamically generated WEP keys (i.e., no WPA) using
# EAP-TLS for authentication and key generation; require both unicast and
# broadcast WEP keys.
network={
	ssid="1x-test"
	key_mgmt=IEEE8021X
	eap=TLS
	identity="user@example.com"
	ca_cert="/etc/cert/ca.pem"
	client_cert="/etc/cert/user.pem"
	private_key="/etc/cert/user.prv"
	private_key_passwd="password"
	eapol_flags=3
}


# LEAP with dynamic WEP keys
network={
	ssid="leap-example"
	key_mgmt=IEEE8021X
	eap=LEAP
	identity="user"
	password="foobar"
}

# EAP-IKEv2 using shared secrets for both server and peer authentication
network={
	ssid="ikev2-example"
	key_mgmt=WPA-EAP
	eap=IKEV2
	identity="user"
	password="foobar"
}

# EAP-FAST with WPA (WPA or WPA2)
network={
	ssid="eap-fast-test"
	key_mgmt=WPA-EAP
	eap=FAST
	anonymous_identity="FAST-000102030405"
	identity="username"
	password="password"
	phase1="fast_provisioning=1"
	pac_file="/etc/wpa_supplicant.eap-fast-pac"
}

network={
	ssid="eap-fast-test"
	key_mgmt=WPA-EAP
	eap=FAST
	anonymous_identity="FAST-000102030405"
	identity="username"
	password="password"
	phase1="fast_provisioning=1"
	pac_file="blob://eap-fast-pac"
}

# Plaintext connection (no WPA, no IEEE 802.1X)
network={
	ssid="plaintext-test"
	key_mgmt=NONE
}


# Shared WEP key connection (no WPA, no IEEE 802.1X)
network={
	ssid="static-wep-test"
	key_mgmt=NONE
	wep_key0="abcde"
	wep_key1=0102030405
	wep_key2="1234567890123"
	wep_tx_keyidx=0
	priority=5
}


# Shared WEP key connection (no WPA, no IEEE 802.1X) using Shared Key
# IEEE 802.11 authentication
network={
	ssid="static-wep-test2"
	key_mgmt=NONE
	wep_key0="abcde"
	wep_key1=0102030405
	wep_key2="1234567890123"
	wep_tx_keyidx=0
	priority=5
	auth_alg=SHARED
}


# IBSS/ad-hoc network with RSN
network={
	ssid="ibss-rsn"
	key_mgmt=WPA-PSK
	proto=RSN
	psk="12345678"
	mode=1
	frequency=2412
	pairwise=CCMP
	group=CCMP
}

# IBSS/ad-hoc network with WPA-None/TKIP (deprecated)
network={
	ssid="test adhoc"
	mode=1
	frequency=2412
	proto=WPA
	key_mgmt=WPA-NONE
	pairwise=NONE
	group=TKIP
	psk="secret passphrase"
}

# open mesh network
network={
	ssid="test mesh"
	mode=5
	frequency=2437
	key_mgmt=NONE
}

# secure (SAE + AMPE) network
network={
	ssid="secure mesh"
	mode=5
	frequency=2437
	key_mgmt=SAE
	psk="very secret passphrase"
}


# Catch all example that allows more or less all configuration modes
network={
	ssid="example"
	scan_ssid=1
	key_mgmt=WPA-EAP WPA-PSK IEEE8021X NONE
	pairwise=CCMP TKIP
	group=CCMP TKIP WEP104 WEP40
	psk="very secret passphrase"
	eap=TTLS PEAP TLS
	identity="user@example.com"
	password="foobar"
	ca_cert="/etc/cert/ca.pem"
	client_cert="/etc/cert/user.pem"
	private_key="/etc/cert/user.prv"
	private_key_passwd="password"
	phase1="peaplabel=0"
}

# Example of EAP-TLS with smartcard (openssl engine)
network={
	ssid="example"
	key_mgmt=WPA-EAP
	eap=TLS
	proto=RSN
	pairwise=CCMP TKIP
	group=CCMP TKIP
	identity="user@example.com"
	ca_cert="/etc/cert/ca.pem"

	# Certificate and/or key identified by PKCS#11 URI (RFC7512)
	client_cert="pkcs11:manufacturer=piv_II;id=%01"
	private_key="pkcs11:manufacturer=piv_II;id=%01"

	# Optional PIN configuration; this can be left out and PIN will be
	# asked through the control interface
	pin="1234"
}

# Example configuration showing how to use an inlined blob as a CA certificate
# data instead of using external file
network={
	ssid="example"
	key_mgmt=WPA-EAP
	eap=TTLS
	identity="user@example.com"
	anonymous_identity="anonymous@example.com"
	password="foobar"
	ca_cert="blob://exampleblob"
	priority=20
}

blob-base64-exampleblob={
SGVsbG8gV29ybGQhCg==
}


# Wildcard match for SSID (plaintext APs only). This example select any
# open AP regardless of its SSID.
network={
	key_mgmt=NONE
}

# Example configuration blacklisting two APs - these will be ignored
# for this network.
network={
	ssid="example"
	psk="very secret passphrase"
	bssid_blacklist=02:11:22:33:44:55 02:22:aa:44:55:66
}

# Example configuration limiting AP selection to a specific set of APs;
# any other AP not matching the masked address will be ignored.
network={
	ssid="example"
	psk="very secret passphrase"
	bssid_whitelist=02:55:ae:bc:00:00/ff:ff:ff:ff:00:00 00:00:77:66:55:44/00:00:ff:ff:ff:ff
}

# Example config file that will only scan on channel 36.
freq_list=5180
network={
	key_mgmt=NONE
}


# Example configuration using EAP-TTLS for authentication and key
# generation for MACsec
network={
	key_mgmt=IEEE8021X
	eap=TTLS
	phase2="auth=PAP"
	anonymous_identity="anonymous@example.com"
	identity="user@example.com"
	password="secretr"
	ca_cert="/etc/cert/ca.pem"
	eapol_flags=0
	macsec_policy=1
}

# Example configuration for MACsec with preshared key
network={
	key_mgmt=NONE
	eapol_flags=0
	macsec_policy=1
	mka_cak=0123456789ABCDEF0123456789ABCDEF
	mka_ckn=6162636465666768696A6B6C6D6E6F707172737475767778797A303132333435
	mka_priority=128
}
`
