package metric

import "github.com/shirou/gopsutil/mem"

type MemMetric struct {
	Available   uint64  `json:"avail"`
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"used_percent"`
}

// done
func setMemMetrics(m *Metrics) {
	if m == nil {
		return
	}

	vms, err := mem.VirtualMemory()
	if err == nil {
		m.VirtualMem = &MemMetric{
			Total:       vms.Total,
			Available:   vms.Available,
			Used:        vms.Used,
			UsedPercent: vms.UsedPercent,
			Free:        vms.Free,
		}
	}
	sms, err := mem.SwapMemory()
	if err == nil {
		m.SwapMem = &MemMetric{
			Total:       sms.Total,
			Used:        sms.Used,
			UsedPercent: sms.UsedPercent,
		}
	}
}
