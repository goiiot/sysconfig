import * as moment from 'moment';
import {DialogNotificationStream, SnackNotificationStream} from "./Streams";

export const showNotificationDialog = (title, msg, ok, onOk, cancel, onCancel) => {
  DialogNotificationStream.next({
    id: moment().unix(),
    title: title,
    msg: msg,
    ok: ok,
    onOk: onOk,
    cancel: cancel,
    onCancel: onCancel,
  });
};

export const showNotificationSnack = (msg, action, onAction) => {
  SnackNotificationStream.next({
    id: moment().unix(),
    msg: msg,
    action: action,
    onAction: onAction,
  });
};