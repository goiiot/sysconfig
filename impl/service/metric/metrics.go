package metric

import (
	"github.com/shirou/gopsutil/disk"
)

type Metrics struct {
	Cpu        *CpuMetric     `json:"cpu"`
	VirtualMem *MemMetric     `json:"virtual_mem"`
	SwapMem    *MemMetric     `json:"swap_mem"`
	Disk       *DiskMetric    `json:"disk"`
	Host       *HostMetric    `json:"host"`
	Net        *NetMetric     `json:"net"`
	Processes  *ProcessMetric `json:"process"`
}

func getCurrentMetrics() *Metrics {
	m := new(Metrics)
	disk.Partitions(true)
	// disk.Usage()
	return m
}

func setAllStaticMetrics(m *Metrics) *Metrics {

	setHostMetrics(m)
	setCpuInfoMetrics(m)
	setDiskInfoMetrics(m)
	setNetInterfaceMetrics(m)

	return m
}

func setAllDynamicMetrics(m *Metrics) *Metrics {

	setCpuUsageMetrics(m)
	setProcessMetrics(m)
	setMemMetrics(m)
	setNetUsageMetrics(m)
	setDiskUsageMetrics(m)

	return m
}

func setAllMetrics(m *Metrics) {
	setAllDynamicMetrics(setAllStaticMetrics(m))
}
