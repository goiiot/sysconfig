import React from 'react';
import PropTypes from 'prop-types';
import {withStyles} from '@material-ui/core/styles';
import DialogContent from "@material-ui/core/DialogContent/DialogContent";
import DialogActions from "@material-ui/core/DialogActions/DialogActions";
import Button from "@material-ui/core/Button/Button";
import Dialog from "@material-ui/core/Dialog/Dialog";
import Tabs from "@material-ui/core/Tabs/Tabs";
import Tab from "@material-ui/core/Tab/Tab";
import AppBar from "@material-ui/core/AppBar/AppBar";
import AddIcon from '@material-ui/icons/Add';

import withRoot from '../../withRoot';
import WebTerminal from "../WebTerminal";
import {createAndConnectTerminal, destroyTerminal, refreshTermList} from "../../mgmt/MgmtTerminals";
import {TermStream} from "../../mgmt/Streams";

const styles = {
  root: {
    width: '100%',
    height: '100%',
    zIndex: 1,
    overflow: 'hidden',
  },
  terminal: {
    width: '100%',
    height: '100%',
    position: 'relative',
  },
};

class DialogTerminal extends React.Component {
  state = {
    selectedTab: 0,
    terms: [],
  };

  timeoutHandle = null;

  handleChange = (event, selectedTab) => {
    this.setState({selectedTab: selectedTab});
    if (this.state.terms.length - 1 < selectedTab) {
      // open new terminal
      createAndConnectTerminal();
    }
  };

  handleTerminalDialogClose = () => {
    if (this.state.selectedTab < this.state.terms.length) {
      let goto = this.state.selectedTab - 1;
      this.setState({selectedTab: (goto < 0) ? 0 : goto});

      destroyTerminal(this.state.terms[this.state.selectedTab].id);
    }
  };

  componentDidMount = () => {
    refreshTermList();
  };

  componentWillMount() {
    TermStream.asObservable().subscribe(
      terms => this.setState({terms: terms}),
      err => console.log(err),
      () => console.log("TermStream Complete"));
  }

  render() {
    const {classes} = this.props;
    return (
      <Dialog fullScreen open={this.props.open} onClose={this.props.closeHandler}
              aria-labelledby="terminal-dialog-title">
        <AppBar position="static" color="default">
          <Tabs value={this.state.selectedTab} onChange={this.handleChange}
                indicatorColor="primary" textColor="primary" color='default'
                scrollable scrollButtons="auto">
            {this.state.terms.map((v, i) => {
              return (
                <Tab label={`Terminal # ${v.id}`} key={`terminal-tab-${v.id}`}
                     selected={i !== this.state.terms.length && i === this.state.selectedTab}/>
              );
            })}
            <Tab icon={<AddIcon/>}/>
          </Tabs>
        </AppBar>
        <DialogContent>
          <WebTerminal className={classes.terminal}
                       id={(this.state.selectedTab === this.state.terms.length) ? null : (this.state.terms[this.state.selectedTab] ? this.state.terms[this.state.selectedTab].id : null)}
                       term={(this.state.selectedTab === this.state.terms.length) ? null : (this.state.terms[this.state.selectedTab] ? this.state.terms[this.state.selectedTab].term : null)}
                       ws={(this.state.selectedTab === this.state.terms.length) ? null : (this.state.terms[this.state.selectedTab] ? this.state.terms[this.state.selectedTab].ws : null)}/>
        </DialogContent>
        <DialogActions>
          <Button onClick={this.handleTerminalDialogClose} color="primary">
            Close
          </Button>
          <Button onClick={this.props.closeHandler} color="primary">
            Hide
          </Button>
        </DialogActions>
      </Dialog>
    );
  }
}

DialogTerminal.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withRoot(withStyles(styles)(DialogTerminal));
