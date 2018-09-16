package metric

import "github.com/shirou/gopsutil/disk"

type DiskPartitionMetric struct {
	Device            string  `json:"dev"`
	MountPoint        string  `json:"mount_point"`
	FsType            string  `json:"fs"`
	Opts              string  `json:"opts"`
	Total             uint64  `json:"total"`
	Free              uint64  `json:"free"`
	Used              uint64  `json:"used"`
	UsedPercent       float64 `json:"used_percent"`
	INodesTotal       uint64  `json:"inodes_total"`
	INodesUsed        uint64  `json:"inodes_used"`
	INodesFree        uint64  `json:"inodes_free"`
	INodesUsedPercent float64 `json:"inodes_used_percent"`
}

type DiskMetric struct {
	Partitions []*DiskPartitionMetric `json:"partitions"`
}

// done
func setDiskInfoMetrics(m *Metrics) {
	if m == nil {
		return
	}

	ps, err := disk.Partitions(true)
	if err != nil {
		return
	}

	if m.Disk == nil {
		m.Disk = &DiskMetric{Partitions: []*DiskPartitionMetric{}}
	}

	for _, v := range ps {
		p := &DiskPartitionMetric{
			Device:     v.Device,
			MountPoint: v.Mountpoint,
			FsType:     v.Fstype,
			Opts:       v.Opts,
		}
		m.Disk.Partitions = append(m.Disk.Partitions, p)
	}
}

// done
func setDiskUsageMetrics(m *Metrics) {
	if m == nil {
		return
	}

	if m.Disk == nil || m.Disk.Partitions == nil {
		setDiskInfoMetrics(m)
	}

	for _, v := range m.Disk.Partitions {
		u, err := disk.Usage(v.MountPoint)
		if err == nil {
			v.Total = u.Total
			v.Free = u.Free
			v.Used = u.Used
			v.UsedPercent = u.UsedPercent
			v.INodesTotal = u.InodesTotal
			v.INodesUsed = u.InodesUsed
			v.INodesFree = u.InodesFree
			v.INodesUsedPercent = u.InodesUsedPercent
		}
	}
}
