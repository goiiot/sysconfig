import {getDynamicMetrics} from "../api/ApiMetrics";
import moment from "moment";
import {MetricsStream} from "./Streams";

const metricsBuffer = [];
const maxBufferSize = 50;
let intervalHandle = null;

export const getBufferedMetrics = () => {
  return metricsBuffer;
};

// typ can be one of ["all", "cpu", "mem", "net", "disk"]
export const startCollectMetrics = (interval) => {
  if (intervalHandle != null) return;

  intervalHandle = setInterval(() => {
    getDynamicMetrics().subscribe(
      m => {
        const time = moment().format("hh:mm:ss");
        const actualMetric = {
          time: time,
          cpu: 0, mem: 0, disk: 0,
          net_bytes_sent: 0, net_bytes_recv: 0,
          net_pkt_sent: 0, net_pkt_recv: 0,
        };

        if (m.cpu && m.cpu.percent) {
          actualMetric.cpu = parseFloat(m.cpu.percent.toFixed(2));
        }

        if (m.virtual_mem && m.virtual_mem.used_percent) {
          actualMetric.mem = parseFloat(m.virtual_mem.used_percent.toFixed(2));
        }

        if (m.disk && m.disk.partitions) {
          const parts = m.disk.partitions;
          // let free = 0;
          let used = 0;
          let total = 0;
          parts.forEach((v) => {
            if (!v) {
              return;
            }

            if (v.free && v.used && v.total) {
              // free += v.free;
              used += v.used;
              total += v.total;
            }
          });
          actualMetric.disk = parseFloat(((used * 100) / total).toFixed(2));
        }

        if (m.net && m.net.interfaces) {
          const ifaces = m.net.interfaces;
          ifaces.forEach((iface) => {
            if (!iface) {
              return;
            }

            if (!iface.bytes_recv || !iface.packets_recv) {
              return;
            }
            actualMetric.net_pkt_sent += iface.packets_sent;
            actualMetric.net_pkt_recv += iface.packets_recv;
            actualMetric.net_bytes_sent += iface.bytes_sent;
            actualMetric.net_bytes_recv += iface.bytes_recv;
          });
        }

        metricsBuffer.push(actualMetric);
        while (metricsBuffer.length > maxBufferSize) {
          metricsBuffer.splice(0, 1);
        }

        MetricsStream.next(metricsBuffer);
      },
      err => console.log(err));
  }, interval);
};

export const stopCollectMetrics = () => {
  clearInterval(intervalHandle);
  intervalHandle = null;
};

// store to web storage for future use
// const persistMetrics = () => {
//
// };
//
// const clearPersistedMetrics = () => {
//
// };