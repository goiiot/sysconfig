import React from "react";
import PropTypes from 'prop-types';
import CardHeader from "@material-ui/core/CardHeader/CardHeader";
import CardActions from "@material-ui/core/CardActions/CardActions";
import Card from "@material-ui/core/Card/Card";
import Button from "@material-ui/core/Button/Button";
import Switch from "@material-ui/core/Switch/Switch";
import withRoot from "../../withRoot";
import {withStyles} from "@material-ui/core";
import {getWifiInfo, getWifiStatus, startWifi, stopWifi, updateWifiConfig} from "../../api/ApiConfigure";
import {showNotificationDialog, showNotificationSnack} from "../../mgmt/MgmtNotification";
import IconButton from "@material-ui/core/IconButton/IconButton";
import RefreshIcon from "@material-ui/icons/Refresh";
import TextField from "@material-ui/core/TextField/TextField";
import CardContent from "@material-ui/core/CardContent/CardContent";

const styles = theme => ({
  card: {
    height: 'auto',
    width: 400,
    overflow: "auto",
    margin: theme.spacing.unit * 2,
  },
});

class CardWifiConfig extends React.Component {
  state = {
    on: false,
    disabled: true,

    config: {}, // TODO split config to multi parts
    status: {}, // TODO split status to multi parts
  };

  handleStatus = () => {
    const {name} = this.props;
    this.setState({disabled: true});

    getWifiStatus(name).subscribe(
      status => this.updateDisplayStatus(status),
      err => showNotificationSnack(`Failed to get status of ${name} ${err}`, "Retry", this.handleStatus),
      () => this.setState({disabled: false}));
  };

  getDisplayConfig = (cfg) => {
    // TODO make displayed config wholesome
    return this.state.config;
  };

  requestUpdateConfig = () => {
    const {name} = this.props;
    this.setState({disabled: true});

    updateWifiConfig(name, this.getDisplayConfig()).subscribe(
      r => {
        // get update result via get device info
        this.handleDeviceInfo();
        showNotificationDialog(`Update ${name} config success`);
      },
      err => showNotificationSnack(`Failed to update config of ${name} ${err}`, "Retry", this.requestUpdateConfig),
      () => this.setState({disabled: false}));
  };

  handleDeviceToggle = () => {
    const on = !this.state.on;
    const {name} = this.props;
    this.setState({disabled: true});

    if (on) {
      startWifi(name).subscribe(
        resp => {
          this.handleStatus();
          showNotificationDialog(`Start ${name} success`, resp);
        },
        err => showNotificationSnack(`Failed to start ${name} ${err}`, "Retry", this.handleStart),
        () => this.setState({disabled: false}));
    } else {
      stopWifi(name).subscribe(
        resp => {
          this.handleStatus();
          showNotificationDialog(`Stop ${name} success`, resp);
        },
        err => showNotificationSnack(`Failed to stop ${name} ${err}`, "Retry", this.handleStop),
        () => this.setState({disabled: false}));
    }
  };

  updateDisplayConfig = (config) => {
    this.setState({config: config});
  };

  updateDisplayStatus = (status) => {
    this.setState({on: status.on, status: status});
  };

  handleDeviceInfo = () => {
    const {name} = this.props;
    this.setState({disabled: true});

    getWifiInfo(name).subscribe(info => {
      this.updateDisplayStatus(info.status);
      this.updateDisplayConfig(info.config);
      this.setState({disabled: false});
    }, err => {
      this.setState({disabled: true});
      showNotificationSnack(`Failed to get config of ${name} ${err}`);
    });
  };

  handleInputChange = (event) => {
    const name = event.target.name;
    this.setState({
      [name]: event.target.value
    });
  };

  componentDidMount() {
    this.handleDeviceInfo();
  }

  render() {
    const {classes, name} = this.props;
    const {disabled, config} = this.state;

    return (
      <Card className={classes.card}>
        <CardHeader title={name} action={[
          <IconButton key={`card-wifi-${name}-refresh`}
                      color="default" aria-label="Refresh"
                      onClick={this.handleDeviceInfo}>
            <RefreshIcon/>
          </IconButton>,
          <Switch key={`card-wifi-${name}-switch`}
                  checked={this.state.on}
                  disabled={this.state.disabled}
                  onClick={this.handleDeviceToggle}/>]}/>
        <CardContent>
          <TextField required name="ssid"
                     label="SSID" margin="dense"
                     disabled={this.state.disabled}
                     value={config.ssid}
                     onChange={this.handleInputChange}/>
          <br/>
          <TextField name="password" type="password"
                     label="Password" margin="dense"
                     disabled={this.state.disabled}
                     value={config.password}
                     onChange={this.handleInputChange}/>
        </CardContent>
        <CardActions>
          <Button variant="raised" component="span" color='primary'
                  disabled={disabled} className={classes.button}
                  onClick={this.requestUpdateConfig}>
            Save
          </Button>
        </CardActions>
      </Card>
    );
  }
}

CardWifiConfig.propTypes = {
  classes: PropTypes.object.isRequired,
  name: PropTypes.string.isRequired,
};

export default withRoot(withStyles(styles)(CardWifiConfig));