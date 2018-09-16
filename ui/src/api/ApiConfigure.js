import {Get, Post, Put} from './apiBase';

const getList = (category) => Get(`/configure/${category}`);
const getInfo = (category, name) => Get(`/configure/${category}/${name}`);
const getConfig = (category, name) => Get(`/configure/${category}/${name}/conf`);
const updateConfig = (category, name, config) => Put(`/configure/${category}/${name}/conf`, null, config);
const actDevice = (category, name, action) => Post(`/configure/${category}/${name}/action`, {action: action});

// lora
export const getLoraList = () => getList('lora');
export const getLoraInfo = (name) => getInfo('lora', name);
export const getLoraConfig = (name) => getConfig('lora', name);
export const updateLoraConfig = (name, config) => updateConfig('lora', name, config);
export const getLoraStatus = (name) => actDevice('lora', name, 'status');
export const startLora = (name) => actDevice('lora', name, 'start');
export const stopLora = (name) => actDevice('lora', name, 'stop');
export const restartLora = (name) => actDevice('lora', name, 'restart');

// bus
export const getBusList = () => getList('bus');
export const getBusInfo = (name) => getInfo('bus', name);
export const getBusConfig = (name) => getConfig('bus', name);
export const updateBusConfig = (name, config) => updateConfig('bus', name, config);
export const getBusStatus = (name) => actDevice('bus', name, 'status');
export const startBus = (name) => actDevice('bus', name, 'start');
export const stopBus = (name) => actDevice('bus', name, 'stop');
export const restartBus = (name) => actDevice('bus', name, 'restart');

// periph
export const getPeriphList = () => getList('periph');
export const getPeriphInfo = (name) => getInfo('periph', name);
export const getPeriphConfig = (name) => getConfig('periph', name);
export const updatePeriphConfig = (name, config) => updateConfig('periph', name, config);
export const getPeriphStatus = (name) => actDevice('periph', name, 'status');
export const startPeriph = (name) => actDevice('periph', name, 'start');
export const stopPeriph = (name) => actDevice('periph', name, 'stop');
export const restartPeriph = (name) => actDevice('periph', name, 'restart');

// net/wifi
export const getWifiList = () => getList('net/wifi');
export const getWifiInfo = (name) => getInfo('net/wifi', name);
export const getWifiConfig = (name) => getConfig('net/wifi', name);
export const updateWifiConfig = (name, config) => updateConfig('net/wifi', name, config);
export const getWifiStatus = (name) => actDevice('net/wifi', name, 'status');
export const startWifi = (name) => actDevice('net/wifi', name, 'start');
export const stopWifi = (name) => actDevice('net/wifi', name, 'stop');
export const restartWifi = (name) => actDevice('net/wifi', name, 'restart');

// net/cell
export const getCellularList = () => getList('net/cell');
export const getCellularInfo = (name) => getInfo('net/cell', name);
export const getCellularConfig = (name) => getConfig('net/cell', name);
export const updateCellularConfig = (name, config) => updateConfig('net/cell', name, config);
export const getCellularStatus = (name) => actDevice('net/cell', name, 'status');
export const startCellular = (name) => actDevice('net/cell', name, 'start');
export const stopCellular = (name) => actDevice('net/cell', name, 'stop');
export const restartCellular = (name) => actDevice('net/cell', name, 'restart');

// net/iface
export const getInterfaceList = () => getList('net/iface');
export const getInterfaceInfo = (name) => getInfo('net/iface', name);
export const getInterfaceConfig = (name) => getConfig('net/iface', name);
export const updateInterfaceConfig = (name, config) => updateConfig('net/iface', name, config);
export const getInterfaceStatus = (name) => actDevice('net/iface', name, 'status');
export const startInterface = (name) => actDevice('net/iface', name, 'start');
export const stopInterface = (name) => actDevice('net/iface', name, 'stop');
export const restartInterface = (name) => actDevice('net/iface', name, 'restart');
