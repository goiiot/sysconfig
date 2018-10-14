import React from 'react';
import PropTypes from 'prop-types';

import CardHeader from '@material-ui/core/CardHeader/CardHeader';
import CardActions from '@material-ui/core/CardActions/CardActions';
import Card from '@material-ui/core/Card/Card';
import Button from '@material-ui/core/Button/Button';
import Typography from '@material-ui/core/Typography/Typography';
import TextField from '@material-ui/core/TextField/TextField';
import CardContent from '@material-ui/core/CardContent/CardContent';
import IconButton from '@material-ui/core/IconButton/IconButton';
import Switch from '@material-ui/core/Switch/Switch';
import RefreshIcon from '@material-ui/icons/Refresh';
import FormControlLabel from "@material-ui/core/FormControlLabel/FormControlLabel";

import withRoot from '../../withRoot';
import {withStyles} from '@material-ui/core';
import {getLoraInfo, getLoraStatus, startLora, stopLora, updateLoraConfig} from '../../api/ApiConfigure';
import {showNotificationDialog, showNotificationSnack} from '../../mgmt/MgmtNotification';

import {loraDefaultConfig} from '../../template';

const styles = (theme) => ({
  card: {
    width: 300,
    height: 450,
    overflow: 'auto',
    margin: theme.spacing.unit * 2
  }
});

class CardLoraConfig extends React.Component {
  state = {
    on: false,
    disabled: true,

    gatewayID: "",
    loraWANPublic: false,

    config: loraDefaultConfig,
    status: {},
  };

  handleStatus = () => {
    const {name} = this.props;
    this.setState({disabled: true});

    getLoraStatus(name).subscribe(
      (status) => this.updateDisplayStatus(status),
      (err) => showNotificationSnack(`Failed to get status of ${name} ${err}`, 'Retry', this.handleStatus),
      () => this.setState({disabled: false})
    );
  };

  getDisplayConfig = () => {
    const config = this.state.config;
    config.gateway_conf.gateway_ID = this.state.gatewayID;
    config.SX1301_conf.lorawan_public = this.state.loraWANPublic;
    return config;
  };

  requestUpdateConfig = () => {
    const {name} = this.props;
    this.setState({disabled: true});

    updateLoraConfig(name, this.getDisplayConfig()).subscribe(
      r => {
        // get update result via get device info
        this.handleDeviceInfo();
        showNotificationDialog(`Update ${name} config success`);
      },
      err =>
        showNotificationSnack(`Failed to update config of ${name} ${err}`, 'Retry', this.requestUpdateConfig),
      () => this.setState({disabled: false})
    );
  };

  handleDeviceToggle = () => {
    const on = !this.state.on;
    const {name} = this.props;
    this.setState({disabled: true});

    if (on) {
      startLora(name).subscribe(
        (resp) => {
          this.handleStatus();
          showNotificationDialog(`Start ${name} success`, resp);
        },
        (err) => showNotificationSnack(`Failed to start ${name} ${err}`, 'Retry', this.handleStart),
        () => this.setState({disabled: false})
      );
    } else {
      stopLora(name).subscribe(
        (resp) => {
          this.handleStatus();
          showNotificationDialog(`Stop ${name} success`, resp);
        },
        (err) => showNotificationSnack(`Failed to stop ${name} ${err}`, 'Retry', this.handleStop),
        () => this.setState({disabled: false})
      );
    }
  };

  updateDisplayConfig = (config) => {
    this.setState({
      config: config,
      gatewayID: config.gateway_conf.gateway_ID,
      loraWANPublic: config.SX1301_conf.lorawan_public,
    });
  };

  updateDisplayStatus = (status) => {
    this.setState({on: status.on, status: status});
  };

  handleDeviceInfo = () => {
    const {name} = this.props;
    this.setState({disabled: true});

    getLoraInfo(name).subscribe(
      (info) => {
        this.updateDisplayStatus(info.status);
        this.updateDisplayConfig(info.config);
        this.setState({disabled: false});
      },
      (err) => {
        this.setState({disabled: true});
        showNotificationSnack(`Failed to get config of ${name} ${err}`);
      }
    );
  };

  handleValueChange = (name) => (event) => {
    const value = event.target.value;
    this.setState({[name]: value});
  };

  handleToggleChange = (name) => (event) => {
    this.setState({
      [name]: event.target.checked
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
        <CardHeader
          title={name}
          action={[
            <IconButton key={`card-lora-pkt-fwd-${name}-refresh`} color="default"
                        aria-label="Refresh" onClick={this.handleDeviceInfo}>
              <RefreshIcon/>
            </IconButton>,
            <Switch key={`card-lora-pkt-fwd-${name}-switch`} checked={this.state.on}
                    disabled={this.state.disabled} onClick={this.handleDeviceToggle}/>
          ]}
          subheader='MQTT Packet Forwarder'
        />
        <CardContent>
          <FormControlLabel control={<Switch checked={this.state.loraWANPublic}
                                             onChange={this.handleToggleChange("loraWANPublic")}/>}
                            label="LoRaWAN Public"/>
          <TextField label="Gateway Id" value={this.state.gatewayID}
                     onChange={this.handleValueChange("gatewayID")}/>
          <br/>
          <br/>
          <Typography variant="title">Radio Information</Typography>
          <br/>
          <Typography variant="body1" color="textPrimary">
            Radio 0 ({config.SX1301_conf.radio_0.tx_enable ? 'Enabled' : 'Disabled'})
          </Typography>
          <Typography variant="body2" color="textPrimary">
            Type: {config.SX1301_conf.radio_0.type} ({config.SX1301_conf.radio_0.enable ? 'Enabled' : 'Disabled'})
          </Typography>
          <Typography variant="body2" color="textPrimary">
            TX Frequency
            Range: {config.SX1301_conf.radio_0.tx_freq_min / 1000000} - {config.SX1301_conf.radio_0.tx_freq_max / 1000000} MHz
          </Typography>
          <Typography variant="body2" color="textPrimary">
            Frequency: {config.SX1301_conf.radio_0.freq / 1000000} MHz
          </Typography>
          <br/>
          <Typography variant="body1" color="textPrimary">
            Radio 1 ({config.SX1301_conf.radio_1.enable ? 'Enabled' : 'Disabled'})
          </Typography>
          <Typography variant="body2" color="textPrimary">
            Type: {config.SX1301_conf.radio_1.type}
          </Typography>
          <Typography variant="body2" color="textPrimary">
            Frequency: {config.SX1301_conf.radio_1.freq / 1000000} MHz
          </Typography>
        </CardContent>
        <CardActions>
          <Button variant="raised" component="span" color="primary" className={classes.button}
                  disabled={disabled} onClick={this.requestUpdateConfig}>
            Save
          </Button>
        </CardActions>
      </Card>
    );
  }
}

CardLoraConfig.propTypes = {
  classes: PropTypes.object.isRequired,
  name: PropTypes.string.isRequired
};

export default withRoot(withStyles(styles)(CardLoraConfig));
