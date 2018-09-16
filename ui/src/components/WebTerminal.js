import * as React from 'react';
// xterm
import {Terminal} from 'xterm';
import * as xtermCSS from 'xterm/lib/xterm.css';
import * as fit from 'xterm/lib/addons/fit/fit';
import * as attach from 'xterm/lib/addons/attach/attach';
import * as search from 'xterm/lib/addons/search/search';
import * as webLinks from 'xterm/lib/addons/webLinks/webLinks';
import * as winptyCompat from 'xterm/lib/addons/winptyCompat/winptyCompat';
// api
import withRoot from "../withRoot";
import {withStyles} from "@material-ui/core";
import {resizeShell} from "../api/ApiShell";

const styles = theme => ({
  content: {
    position: 'absolute',
    left: 0,
    right: 0,
    top: theme.mixins.toolbar.minHeight,
    bottom: theme.mixins.toolbar.minHeight,
  }
});

const getTermColsRows = (term, container) => {
  const
    cellHeight = term._core.renderer.dimensions.actualCellHeight,
    cellWidth = term._core.renderer.dimensions.actualCellWidth;
  const rows = Math.round(container.offsetHeight / cellHeight);
  const cols = Math.round(container.offsetWidth / cellWidth);
  return {rows, cols};
};

let onResizeFunc = null;
let timeoutHandle = null;

const onResize = (term, container, id) => {
  return (event) => {
    if (!term) {
      return
    }
    clearTimeout(timeoutHandle);

    const {cols, rows} = getTermColsRows(term, container);
    timeoutHandle = setTimeout(() => {
      term.fit();
      resizeShell(id, cols, rows); // best effort
      timeoutHandle = null;
    }, 500);
  };
};

class WebTerminal extends React.PureComponent {
  term = null;
  ws = null;
  id = null;

  state = {
    cmd: "",
  };

  constructor(props) {
    super(props);

    Terminal.applyAddon(attach);
    Terminal.applyAddon(webLinks);
    Terminal.applyAddon(winptyCompat);
    Terminal.applyAddon(search);
    Terminal.applyAddon(fit);
  }

  prepareTerm() {
    if (!this.props.term || !this.props.ws) {
      return
    }

    while (this.container.firstChild) {
      this.container.removeChild(this.container.firstChild);
    }

    // required
    this.id = this.props.id;
    this.term = this.props.term;
    this.ws = this.props.ws;

    this.term.open(this.container);
    this.term.winptyCompatInit();
    this.term.webLinksInit();
    this.term.fit();

    const {cols, rows} = getTermColsRows(this.term, this.container);
    resizeShell(this.id, cols, rows).subscribe(
      ok => console.log("resize success"),
      err => console.log(err));

    this.term.refresh(0, this.term.rows - 1);
  }

  getSnapshotBeforeUpdate(prevProps, prevState) {
    this.prepareTerm();
    return null;
  }

  componentDidUpdate(prevProps, prevState, snapShot) {
  }

  componentDidMount() {
    this.prepareTerm();
    onResizeFunc = onResize(this.term, this.container, this.id);
    window.addEventListener('resize', onResizeFunc, true);
  }

  componentWillUnmount() {
    window.removeEventListener('resize', onResizeFunc, true);
    this.term = null;
    this.ws = null;
    this.id = null;
  }

  render() {
    const {classes} = this.props;
    return (
      <div ref={ref => (this.container = ref)} className={classes.content}/>
    );
  }
}

export default withRoot(withStyles(styles)(withStyles(xtermCSS)(WebTerminal)));