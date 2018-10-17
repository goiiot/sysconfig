import React from 'react';
import PropTypes from 'prop-types';
import CardHeader from '@material-ui/core/CardHeader/CardHeader';
import CardActions from '@material-ui/core/CardActions/CardActions';
import Card from '@material-ui/core/Card/Card';
import Button from '@material-ui/core/Button/Button';
import Switch from '@material-ui/core/Switch/Switch';
import withRoot from '../../withRoot';
import {CardContent, withStyles} from '@material-ui/core';
import {getLoraInfo, getLoraStatus, startLora, stopLora, updateLoraConfig} from '../../api/ApiConfigure';
import {showNotificationDialog, showNotificationSnack} from '../../mgmt/MgmtNotification';
import IconButton from '@material-ui/core/IconButton/IconButton';
import RefreshIcon from '@material-ui/icons/Refresh';
import TextField from '@material-ui/core/TextField';
import {loraBridgeDefaultConfig} from '../../template';

const styles = (theme) => ({
  card: {
    width: 300,
    height: 380,
    overflow: 'auto',
    margin: theme.spacing.unit * 2
  }
});

class CardLoraBridgeConfig extends React.Component {
  state = {
    on: false,
    disabled: true,

    mqttServer: "",
    mqttQoS: 0,
    mqttUsername: "",
    mqttPassword: "",

    config: loraBridgeDefaultConfig,
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
    const generic = this.state.config.backend.mqtt;
    generic.server = this.state.mqttServer;
    generic.qos = parseInt(this.state.mqttQoS, 10);
    generic.username = this.state.mqttUsername;
    generic.password = this.state.mqttPassword;
    return this.state.config;
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
      mqttServer: config.backend.mqtt.server,
      mqttQoS: config.backend.mqtt.qos,
      mqttUsername: config.backend.mqtt.username,
      mqttPassword: config.backend.mqtt.password,
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

  handleQoSChange = (event) => {
    const value = event.target.value;
    this.setState({
      mqttQoS: value < 0 ? 0 : value > 2 ? 2 : value
    });
  };

  handleValueChange = (name) => (event) => {
    const value = event.target.value;
    this.setState({[name]: value});
  };

  componentDidMount() {
    this.handleDeviceInfo();
  }

  render() {
    const {classes, name} = this.props;
    const {disabled} = this.state;

    return (
      <Card className={classes.card}>
        <CardHeader
          title={name}
          action={[
            <IconButton key={`card-lora-bridge-${name}-refresh`} color="default"
                        aria-label="Refresh" onClick={this.handleDeviceInfo}>
              <RefreshIcon/>
            </IconButton>,
            <Switch key={`card-lora-bridge-${name}-switch`} checked={this.state.on}
                    disabled={this.state.disabled} onClick={this.handleDeviceToggle}/>
          ]}
          subheader="MQTT Gateway Forwarder"
        />
        <CardContent>
          <TextField required label="MQTT Server" value={this.state.mqttServer}
                     disabled={disabled} onChange={this.handleValueChange('mqttServer')}/>
          <br/>
          <br/>
          <TextField required type="number" label="MQTT QoS" value={this.state.mqttQoS}
                     disabled={disabled} onChange={this.handleQoSChange}/>
          <br/>
          <br/>
          <TextField label="MQTT Username" value={this.state.mqttUsername}
                     disabled={disabled} onChange={this.handleValueChange('mqttUsername')}/>
          <br/>
          <br/>
          <TextField label="MQTT Password" value={this.state.mqttPassword}
                     disabled={disabled} onChange={this.handleValueChange('mqttPassword')}/>
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

CardLoraBridgeConfig.propTypes = {
  classes: PropTypes.object.isRequired,
  name: PropTypes.string.isRequired
};

export default withRoot(withStyles(styles)(CardLoraBridgeConfig));
