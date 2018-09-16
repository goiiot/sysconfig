import {Post} from "./apiBase";

export const requestReboot = (time) => Post("/power/reboot", {time: time});
export const requestShutdown = (time) => Post("/power/shutdown", {time: time});