package metric

import (
	"github.com/shirou/gopsutil/cpu"
)

type CpuCoreMetric struct {
	UsedPercent float64 `json:"used_percent"`
	Freq        float64 `json:"freq"`
	Model       string  `json:"model"`
	Family      string  `json:"family"`
	VendorID    string  `json:"vendorId"`
}

type CpuMetric struct {
	Percent float64          `json:"percent"`
	Cores   []*CpuCoreMetric `json:"cores"`
}

// done
func setCpuInfoMetrics(m *Metrics) {
	if m == nil {
		return
	}

	info, err := cpu.Info()
	if err != nil {
		return
	}

	if m.Cpu == nil {
		m.Cpu = &CpuMetric{
			Cores: make([]*CpuCoreMetric, len(info)),
		}
	}

	for i := range info {
		cc := &CpuCoreMetric{
			Freq:     info[i].Mhz,
			Model:    info[i].Model,
			Family:   info[i].Model,
			VendorID: info[i].VendorID,
		}
		m.Cpu.Cores = append(m.Cpu.Cores, cc)
	}
}

// done
func setCpuUsageMetrics(m *Metrics) {
	if m == nil {
		return
	}

	if m.Cpu == nil {
		m.Cpu = &CpuMetric{Cores: nil}
	}

	p, err := cpu.Percent(0, false)
	if err == nil {
		m.Cpu.Percent = p[0]
	}

	ps, err := cpu.Percent(0, true)
	if err != nil {
		return
	}

	if m.Cpu.Cores == nil {
		m.Cpu.Cores = make([]*CpuCoreMetric, len(ps))
	}

	if len(ps) > len(m.Cpu.Cores) {
		m.Cpu.Cores = append(m.Cpu.Cores, make([]*CpuCoreMetric, len(ps)-len(m.Cpu.Cores))...)
	}
	for i := range ps {
		if m.Cpu.Cores[i] == nil {
			m.Cpu.Cores[i] = &CpuCoreMetric{}
		}

		m.Cpu.Cores[i].UsedPercent = ps[i]
	}
}
