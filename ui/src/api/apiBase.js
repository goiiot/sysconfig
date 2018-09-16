/* eslint-disable flowtype/require-valid-file-annotation */

import axios from 'axios';
import {apiConfig as cfg} from '../config';
import {AuthStream} from "../mgmt/Streams";
import {Subject} from "rxjs";

axios.interceptors.response.use(response => {
  return response;
}, error => {
  if (error && error.response) {
    if (error.response.status === 401) {
      AuthStream.next(false);
      return "";
    }
  }

  return Promise.reject(error.response);
});

// wrap AxiosPromise to Observable
const wrapObservable = (promise) => {
  const sub = new Subject();
  promise
    .then(resp => {
      if (resp.data.ok) {
        sub.next(resp.data.data);
      } else {
        sub.error(resp.data.reason);
      }

      sub.complete();
    })
    .catch(err => {
      sub.error(err);
      sub.complete();
    });
  return sub.asObservable();
};

const rawReq = (method, path, config) => {
  if (!config) {
    config = {};
  }

  config.method = method;
  config.url = buildURL(path);
  config.maxRedirects = cfg.maxRedirects;
  config.withCredentials = true;
  return axios.request(config);
};

export function RawGet(path, params, config) {
  if (!config) {
    config = {};
  }

  config.params = params;
  return rawReq('GET', path, config);
}

const req = (method, path, config) => wrapObservable(rawReq(method, path, config));

export function Get(path, params, config) {
  if (!config) {
    config = {};
  }
  config.params = params;
  return req('GET', path, config);
}

export function Post(path, params, body, config) {
  if (!config) {
    config = {};
  }
  config.params = params;
  config.data = body;
  return req('POST', path, config);
}

export function Put(path, params, body, config) {
  if (!config) {
    config = {};
  }
  config.params = params;
  config.data = body;
  return req('PUT', path, config);
}

export function Delete(path, params, body, config) {
  if (!config) {
    config = {};
  }
  config.params = params;
  config.data = body;
  return req('DELETE', path, config);
}

const buildURL = (path) => {
  if (path.charAt(0) !== '/') {
    path = "/" + path;
  }
  return `//${cfg.apiServer}/api/${cfg.apiVersion}${path}`;
};