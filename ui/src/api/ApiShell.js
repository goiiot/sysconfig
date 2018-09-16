import autobahn from "autobahn";
import {apiConfig as cfg} from '../config';
import {Delete, Get, Post} from './apiBase';

export const listShells = () => Get(`/shell`);
export const createShell = (cols, rows, shell) => Post(`/shell`, {cols: cols, rows: rows, shell: shell});
export const closeShell = (id) => Delete(`/shell/${id}`);
export const resizeShell = (id, cols, rows) => Post(`/shell/${id}/size`, {cols: cols, rows: rows});
export const connectShell = (id) => {
  const protocol = (window.location.protocol === 'https:') ? 'wss' : 'ws';
  const wsURL = `${protocol}://${cfg.apiServer}/api/${cfg.apiVersion}/shell/${id}`;

  if ('WebSocket' in window) {
    return new WebSocket(wsURL);
  } else {
    return new autobahn.Connection({url: wsURL});
  }
};
