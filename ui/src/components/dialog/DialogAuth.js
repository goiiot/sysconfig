import React from 'react';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogTitle from '@material-ui/core/DialogTitle';

import {login} from "../../api/ApiAuth";
import {AuthStream} from "../../mgmt/Streams";

export default class DialogAuth extends React.Component {
  state = {
    username: "",
    password: "",
  };

  handleAuthRequest = () =>
    login(this.state.username, this.state.password).subscribe(
      ok => AuthStream.next(true),
      err => AuthStream.next(false));

  handleUsernameInput = (event) => {
    this.setState({username: event.target.value});
  };

  handlePasswordInput = (event) => {
    this.setState({password: event.target.value});
  };

  render() {
    return (
      <Dialog
        open={this.props.open}
        aria-labelledby="form-dialog-title">
        <DialogTitle id="form-dialog-title">Login</DialogTitle>
        <DialogContent>
          <TextField autoFocus margin="dense" id="login-username" label="User Name" type="text"
                     fullWidth vlaue={this.state.username} onChange={this.handleUsernameInput}/>
          <TextField margin="dense" id="login-password" label="Password" type="password"
                     fullWidth vlaue={this.state.password} onChange={this.handlePasswordInput}/>
        </DialogContent>
        <DialogActions>
          <Button onClick={this.handleAuthRequest} color="primary">
            Login
          </Button>
        </DialogActions>
      </Dialog>
    );
  }
}