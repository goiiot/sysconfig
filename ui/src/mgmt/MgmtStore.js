import store from 'store';
import defaults from 'store/plugins/defaults';
import {PrefStream} from './Streams';

const keyPreference = "preference";
const keyMetrics = "metrics";

store.addPlugin(defaults);
store.defaults({
  preference: {
    monitoring: {
      enabled: true,
      refresh_interval: 5000,
    },
  },
  metrics: [],
});

export const storePreference = (preference) => {
  store.set(keyPreference, preference);
  PrefStream.next(preference);
};

export const getPreference = () => {
  return store.get(keyPreference);
};

export const storeMetrics = (metrics) => {
  store.set(keyMetrics, metrics);
};

export const getMetrics = () => {
  return store.get(keyMetrics);
};