package metric

import "github.com/shirou/gopsutil/process"

type ProcessMetric struct {
	Count int `json:"count"`
}

// done
func setProcessMetrics(m *Metrics) {
	if m == nil {
		return
	}

	pids, err := process.Pids()
	if err == nil {
		m.Processes = &ProcessMetric{
			Count: len(pids),
		}
	}
}
