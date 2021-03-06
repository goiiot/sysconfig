package model

type Status struct {
	On bool `json:"on" confl:"on"`
}

type LoraStatus struct {
	Status
}

type WifiStatus struct {
	Status
	SSIDList []string `json:"ssid_list" confl:"ssid_list"`
}
