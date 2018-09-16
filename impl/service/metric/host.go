package metric

import "github.com/shirou/gopsutil/host"

type HostMetric struct {
	Hostname      string `json:"hostname"`
	UpTime        uint64 `json:"up_time"`
	BootTime      uint64 `json:"boot_time"`
	CpuCount      uint64 `json:"cpu_count"`
	OS            string `json:"os"`
	Platform      string `json:"platform"`
	KernelVersion string `json:"kernel_version"`
	HostID        string `json:"host_id"`
}

// done
func setHostMetrics(m *Metrics) {
	if m == nil {
		return
	}

	hinfo, err := host.Info()
	if err == nil {
		m.Host = &HostMetric{
			Hostname:      hinfo.Hostname,
			UpTime:        hinfo.Uptime,
			BootTime:      hinfo.BootTime,
			CpuCount:      hinfo.Procs,
			OS:            hinfo.OS,
			Platform:      hinfo.Platform,
			KernelVersion: hinfo.KernelVersion,
			HostID:        hinfo.HostID,
		}
	}
}
