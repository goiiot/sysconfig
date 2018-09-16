import {Terminal} from 'xterm';
import * as fit from 'xterm/lib/addons/fit/fit';
import * as attach from 'xterm/lib/addons/attach/attach';
import * as search from 'xterm/lib/addons/search/search';
import * as webLinks from 'xterm/lib/addons/webLinks/webLinks';
import * as winptyCompat from 'xterm/lib/addons/winptyCompat/winptyCompat';
import {closeShell, connectShell, createShell, listShells} from "../api/ApiShell";
import {TermStream} from "./Streams";

const terms = [];

export const refreshTermList = () => {
  listShells().subscribe(
    shs => {
      // remove non-exist ones
      terms.forEach((v, i) => {
        let found = false;
        shs.forEach(value => {
          if (v.id === value) {
            found = true;
          }
        });

        if (!found) {
          terms.splice(i, 1);
        }
      });

      // add usable ones
      shs.forEach(v => {
        let found = false;
        terms.forEach(value => {
          if (value.id === v) {
            found = true;
          }
        });

        if (!found) {
          const term = createTerminal();
          const socket = connectShell(v);
          socket.onopen = () => term.attach(socket);
          terms.push({id: v, term: term, ws: socket})
        }
      });

      terms.sort((a, b) => {
        return parseInt(a.id, 10) - parseInt(b.id, 10);
      });
      TermStream.next(terms);
    },
    err => console.log(err));
};

const createTerminal = () => {
  Terminal.applyAddon(attach);
  Terminal.applyAddon(webLinks);
  Terminal.applyAddon(winptyCompat);
  Terminal.applyAddon(search);
  Terminal.applyAddon(fit);

  return new Terminal({
    cursorBlink: false,
    allowTransparency: false,
    cursorStyle: "block",
    fontSize: 12,
  });
};

export const createAndConnectTerminal = () => {
  const term = createTerminal();

  createShell(term.cols, term.rows).subscribe(
    t => {
      const socket = connectShell(t.id);
      socket.onopen = () => term.attach(socket);
      terms.push({id: t.id, term: term, ws: socket});
      TermStream.next(terms);
    },
    err => console.log(err));
};

export const destroyTerminal = (id) => {
  let todelIdx = null, todel = null;
  terms.forEach((v, i) => {
    if (v.id === id) {
      todelIdx = i;
      todel = v;
    }
  });

  if (todel) {
    terms.splice(todelIdx, 1);
    closeShell(id); // best effort
    todel.term.dispose();
    todel.ws.close();
  }

  TermStream.next(terms);
};