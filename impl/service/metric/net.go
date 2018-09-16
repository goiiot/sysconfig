package metric

import "github.com/shirou/gopsutil/net"

type InterfaceMetric struct {
	Name        string   `json:"name"`
	MTU         int      `json:"mtu"`
	HwAddr      string   `json:"hw_addr"`
	Addrs       []string `json:"addrs"`
	Flags       []string `json:"flags"`
	BytesSent   uint64   `json:"bytes_sent"`
	BytesRecv   uint64   `json:"bytes_recv"`
	PacketsSent uint64   `json:"packets_sent"`
	PacketsRecv uint64   `json:"packets_recv"`
}

type NetMetric struct {
	Interfaces []*InterfaceMetric `json:"interfaces"`
}

// done
func setNetUsageMetrics(m *Metrics) {
	if m == nil {
		return
	}

	cs, err := net.IOCounters(true)
	if err != nil {
		return
	}

	if m.Net == nil {
		m.Net = &NetMetric{Interfaces: nil}
	}
	if m.Net.Interfaces == nil {
		m.Net.Interfaces = make([]*InterfaceMetric, len(cs))
		for _, v := range cs {
			f := &InterfaceMetric{
				Name:        v.Name,
				BytesSent:   v.BytesSent,
				BytesRecv:   v.BytesRecv,
				PacketsSent: v.PacketsSent,
				PacketsRecv: v.PacketsRecv,
			}
			m.Net.Interfaces = append(m.Net.Interfaces, f)
		}
	} else {
		for _, f := range cs {
			for _, fa := range m.Net.Interfaces {
				if f.Name == fa.Name {
					fa.BytesSent = f.BytesSent
					fa.BytesRecv = f.BytesRecv
					fa.PacketsSent = f.PacketsSent
					fa.PacketsRecv = f.PacketsRecv
				}
			}
		}
	}
}

// done
func setNetInterfaceMetrics(m *Metrics) {
	if m == nil {
		return
	}

	if m.Net == nil {
		m.Net = &NetMetric{
			Interfaces: []*InterfaceMetric{},
		}
	}

	ifaces, err := net.Interfaces()
	if err != nil {
		return
	}

	for i := range ifaces {
		f := &InterfaceMetric{
			Name:   ifaces[i].Name,
			MTU:    ifaces[i].MTU,
			Flags:  ifaces[i].Flags,
			HwAddr: ifaces[i].HardwareAddr,
		}

		for j := range ifaces[i].Addrs {
			f.Addrs = append(f.Addrs, ifaces[i].Addrs[j].Addr)
		}
		m.Net.Interfaces = append(m.Net.Interfaces, f)
	}
}

// todo
// func setNetConnectionMetrics(m *Metrics) {
// 	if m == nil {
// 		return
// 	}
// }
