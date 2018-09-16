import React from 'react';
import Button from '@material-ui/core/Button';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogTitle from '@material-ui/core/DialogTitle';
import Input from '@material-ui/core/Input';
import InputAdornment from '@material-ui/core/InputAdornment';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableRow from '@material-ui/core/TableRow';
import Typography from '@material-ui/core/Typography';

import {getPreference, storePreference} from "../../mgmt/MgmtStore";
import {withStyles} from "@material-ui/core";
import Switch from "@material-ui/core/Switch/Switch";

const styles = {
  root: {
    width: "100%",
    height: "100%",
  }
};

const PrefItem = (props) => {
  return (
    <TableRow>
      <TableCell>{props.name}</TableCell>
      <TableCell>
        {props.component}
      </TableCell>
    </TableRow>
  )
};

class DialogPref extends React.Component {
  state = {
    monitoring: false,
    refresh_interval: "",
  };
  handleStorePreference = () => {
    const pref = getPreference();
    const m = pref.monitoring;
    m.enabled = this.state.monitoring;
    m.refresh_interval = parseInt(this.state.refresh_interval, 10);
    storePreference(pref);
    this.props.closeHandler();
  };
  toggleMonitoring = () => {
    const newState = !this.state.monitoring;
    this.setState({monitoring: newState});
  };
  handleRefreshIntervalInput = (event) => {
    this.setState({refresh_interval: event.target.value});
  };

  componentDidMount() {
    const pref = getPreference();
    this.setState({
      monitoring: pref.monitoring.enabled,
      refresh_interval: `${pref.monitoring.refresh_interval}`
    });
  }

  render() {
    const {classes} = this.props;
    return (
      <Dialog className={classes.root}
              scroll='paper'
              onClose={this.props.closeHandler}
              open={this.props.open}
              aria-labelledby="form-dialog-title">
        <DialogTitle id="form-dialog-title">Preference</DialogTitle>
        <DialogContent>
          <Table>
            <TableBody>
              <TableRow>
                <TableCell>
                  <Typography variant="subheading">
                    Monitoring
                  </Typography>
                </TableCell>
              </TableRow>
              <PrefItem name="Default Enabled" component={
                <Switch checked={this.state.monitoring}
                        onChange={this.toggleMonitoring}
                        value="Monitoring"
                        color="secondary"/>}/>
              <PrefItem name="Refresh Interval" component={
                <Input id="adornment-weight"
                       value={this.state.refresh_interval}
                       onChange={this.handleRefreshIntervalInput}
                       endAdornment={<InputAdornment position="end">ms</InputAdornment>}/>}/>
            </TableBody>
          </Table>
        </DialogContent>
        <DialogActions>
          <Button onClick={this.handleStorePreference} color="primary">
            Save
          </Button>
        </DialogActions>
      </Dialog>
    );
  }
}

export default withStyles(styles, {withTheme: true})(DialogPref);