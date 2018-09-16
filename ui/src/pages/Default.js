import React from 'react';
import PropTypes from 'prop-types';
import {withStyles} from '@material-ui/core/styles';
import withRoot from '../withRoot';
// components
import NavBar from '../components/NavBar';
import {AuthStream, DialogNotificationStream, SnackNotificationStream, VersionStream} from "../mgmt/Streams";
import DialogTerminal from "../components/dialog/DialogTerminal";
import DialogAuth from "../components/dialog/DialogAuth";
import TerminalIcon from "../icons/TerminalIcon";
import Button from "@material-ui/core/Button/Button";
import {getVersion} from "../api/ApiVersion";
import DialogNotification from "../components/dialog/DialogNotification";
import SnackNotification from "../components/SnackNotification";

const styles = theme => ({
  root: {
    width: '100%',
    height: '100%',
    overflow: 'hidden',
  },
  fab: {
    zIndex: 1001,
    position: 'absolute',
    bottom: theme.spacing.unit * 4,
    right: theme.spacing.unit * 4,
  },
  noteDialog: {
    zIndex: 10000,
  },
  snack: {
    zIndex: 10000,
  }
});

class Default extends React.Component {
  state = {
    termDialogOpen: false,
    authDialogOpen: false,

    noteTitle: "",
    noteMsg: "",
    notes: [],
    noteDialogOpen: false,

    snackMsg: "",
    snacks: [],
    snackOpen: false,
  };

  handleTermDialogToggle = () => {
    const open = !this.state.termDialogOpen;
    this.setState({termDialogOpen: open});
  };

  handleNoteDialogClose = (id) => () => {
    const copy = this.state.notes.slice(0);
    for (let i = 0; i < copy.length; i++) {
      if (copy[i].id === id) {
        copy.splice(i, 1);
        break;
      }
    }
    this.setState({notes: copy})
  };

  handleSnackClose = (id) => () => {
    const copy = this.state.snacks.slice(0);
    for (let i = 0; i < copy.length; i++) {
      if (copy[i].id === id) {
        copy.splice(i, 1);
        break;
      }
    }

    this.setState({snacks: copy})
  };

  showDialogNotification = (n) => {
    const copy = [...this.state.notes, n];
    this.setState({notes: copy});
  };

  showSnackNotification = (n) => {
    const copy = [...this.state.snacks, n];
    this.setState({snacks: copy});
  };

  componentDidMount() {
    AuthStream.asObservable().subscribe(
      ok => this.setState({authDialogOpen: !ok}),
      err => console.log(err));

    DialogNotificationStream.asObservable().subscribe(
      n => this.showDialogNotification(n),
      err => console.log(err));

    SnackNotificationStream.asObservable().subscribe(
      n => this.showSnackNotification(n),
      err => console.log(err));

    getVersion().subscribe(
      v => VersionStream.next(v),
      err => console.log(err));
  }

  render() {
    const {classes} = this.props;
    return (
      <div className={classes.root}>
        <NavBar/>
        <Button variant="fab" className={classes.fab} color='primary' onClick={this.handleTermDialogToggle}>
          <TerminalIcon/>
        </Button>

        <DialogTerminal open={this.state.termDialogOpen} closeHandler={this.handleTermDialogToggle}/>
        <DialogAuth open={this.state.authDialogOpen}/>
        {this.state.notes.map((v) => {
          return (
            <DialogNotification key={`dialog-note-${v.id}`}
                                className={classes.noteDialog}
                                title={v.title} msg={v.msg}
                                ok={v.ok} cancel={v.cancel}
                                onOk={() => {
                                  v.onOk && v.onOk();
                                  this.handleNoteDialogClose(v.id);
                                }}
                                onCancel={() => {
                                  v.onCancel && v.onCancel();
                                  this.handleNoteDialogClose(v.id);
                                }}
                                closeHandler={this.handleNoteDialogClose(v.id)}/>
          );
        })}
        {this.state.snacks.map((v) => {
          return (
            <SnackNotification key={`snack-note-${v.id}`}
                               open={true}
                               className={classes.snack}
                               msg={v.msg} action={v.action}
                               onAction={() => {
                                 v.onAction();
                                 this.handleSnackClose(v.id);
                               }}
                               closeHandler={this.handleSnackClose(v.id)}/>
          );
        })}
      </div>
    );
  }
}

Default.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withRoot(withStyles(styles)(Default));
