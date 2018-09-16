import {Put, RawGet} from "./apiBase";
import {Subject} from "rxjs";

export const uploadFiles = (dst, files) => {
  const data = new FormData();
  data.append("files", files);
  return Put("/file", {dst: dst}, data);
};

export const downloadFiles = (src, format) => {
  const sub = new Subject();
  RawGet("/file", {src: src, format: format}, {responseType: 'blob'})
    .then(resp => {
      let filename = "";
      // find filename with best effort
      const disposition = resp.headers['content-disposition'];
      if (disposition) {
        filename = decodeURI(disposition.match(/filename="(.*)"/)[1]);
      } else {
        const parts = src.split("/");
        if (parts.length > 0) {
          filename = `${parts[parts.length - 1]}.${format}`;
        } else {
          filename = `file_download.${format}`
        }
      }

      const url = window.URL.createObjectURL(new Blob([resp.data]));
      const link = document.createElement('a');
      link.href = url;
      link.setAttribute('download', filename);
      document.body.appendChild(link);
      link.click();
      document.body.removeChild(link);

      sub.complete();
    })
    .catch(err => sub.error(err));

  return sub.asObservable();
};