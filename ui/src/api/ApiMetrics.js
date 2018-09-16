import {Get} from "./apiBase";

export const getAllMetrics = () => Get("/metrics");
export const getDynamicMetrics = () => Get("/metrics/aggregated/dynamic");
export const getStaticMetrics = () => Get("/metrics/aggregated/static");
export const getCpuMetrics = () => Get("/metrics/cpu");
export const getCpuInfoMetrics = () => Get("/metrics/cpu/info");
export const getCpuUsageMetrics = () => Get("/metrics/cpu/usage");
export const getMemMetrics = () => Get("/metrics/mem");
export const getDiskMetrics = () => Get("/metrics/disk");
export const getDiskInfoMetrics = () => Get("/metrics/disk/info");
export const getDiskUsageMetrics = () => Get("/metrics/disk/usage");
export const getNetMetrics = () => Get("/metrics/net");
export const getNetInfoMetrics = () => Get("/metrics/net/info");
export const getNetUsageMetrics = () => Get("/metrics/net/usage");
