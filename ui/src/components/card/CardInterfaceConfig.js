import React from 'react';
import PropTypes from 'prop-types';
import CardHeader from '@material-ui/core/CardHeader/CardHeader';
import CardActions from '@material-ui/core/CardActions/CardActions';
import Card from '@material-ui/core/Card/Card';
import Button from '@material-ui/core/Button/Button';
import withRoot from '../../withRoot';
import {withStyles} from '@material-ui/core';
import {requestReboot} from '../../api/ApiPower';
import {
  getInterfaceConfig,
  getInterfaceStatus,
  restartInterface,
  startInterface,
  stopInterface,
  updateInterfaceConfig
} from '../../api/ApiConfigure';
import {showNotificationDialog, showNotificationSnack} from '../../mgmt/MgmtNotification';
import Typography from '@material-ui/core/Typography/Typography';
import TextField from '@material-ui/core/TextField/TextField';
import CardContent from '@material-ui/core/CardContent/CardContent';
import Switch from '@material-ui/core/Switch/Switch';

const styles = (theme) => ({
  card: {
    height: 'auto',
    width: 400,
    overflow: 'auto',
    margin: theme.spacing.unit * 2
  }
});

class CardInterfaceConfig extends React.Component {
  state = {
    disabled: true,
    config: {},

    ipError: false,
    netmaskError: false,
    gatewayError: false,

    useDHCP: true,
    isRunning: true,
    ip: '192.168.1.1',
    netmask: '255.255.255.0',
    gateway: '192.168.1.1'
  };

  handleStart = () => {
    const {name} = this.props;
    startInterface(name).subscribe(
      (resp) => showNotificationDialog(`Start ${name} success`, resp),
      (err) => showNotificationSnack(`Failed to start ${name} ${err}`, 'Retry', this.handleStart)
    );
    requestReboot();
  };

  handleStop = () => {
    const {name} = this.props;
    stopInterface(name).subscribe(
      (resp) => showNotificationDialog(`Stop ${name} success`, resp),
      (err) => showNotificationSnack(`Failed to stop ${name} ${err}`, 'Retry', this.handleStop)
    );
  };

  handleRestart = () => {
    const {name} = this.props;
    restartInterface(name).subscribe(
      (resp) => showNotificationDialog(`Restart ${name} success`, resp),
      (err) => showNotificationSnack(`Failed to restart ${name} ${err}`, 'Retry', this.handleRestart)
    );
  };

  handleStatus = () => {
    const {name} = this.props;
    getInterfaceStatus(name).subscribe(
      (resp) => showNotificationDialog(`Get ${name} Status success`, resp),
      (err) => showNotificationSnack(`Failed to get status of ${name} ${err}`, 'Retry', this.handleStatus)
    );
  };

  handleUpdateConfig = () => {
    const {name} = this.props;
    updateInterfaceConfig(name, this.state.config).subscribe(
      (resp) => showNotificationDialog(`Update ${name} success`, resp),
      (err) =>
        showNotificationSnack(`Failed to update config of ${name} ${err}`, 'Retry', this.handleUpdateConfig)
    );
  };
  handleToggleChange = (event) => {
    const name = event.target.name;
    this.setState({
      [name]: event.target.checked
    });
  };
  handleInputChange = (event) => {
    const name = event.target.name;
    const value = event.target.value;
    this.setState({
      [name]: value
    });

    if (name === 'ip') {
      const testRegExp = /^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$/;
      testRegExp.test(value) ? this.setState({ipError: false}) : this.setState({ipError: true});
    }
    if (name === 'netmask') {
      const testRegExp = /^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$/;
      testRegExp.test(value) ? this.setState({netmaskError: false}) : this.setState({netmaskError: true});
    }
    if (name === 'gateway') {
      const testRegExp = /^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$/;
      testRegExp.test(value) ? this.setState({gatewayError: false}) : this.setState({gatewayError: true});
    }
  };
  handleClickDHCP = () => {
    this.setState({useDHCP: true});
  };
  handleClickManual = () => {
    this.setState({useDHCP: false});
  };
  handleSet = () => {
  };
  handleDHCPC = () => {
  };

  componentDidMount() {
    const {name} = this.props;
    getInterfaceConfig(name).subscribe(
      (c) => this.setState({disabled: false, config: c}),
      (err) => showNotificationSnack(`Failed to get config of ${name} ${err}`)
    );
  }

  render() {
    const {classes, name} = this.props;
    const {disabled, config} = this.state;

    return (
      <Card className={classes.card}>
        <CardHeader
          title={name}
          action={
            <Switch checked={this.state.isRunning} onChange={this.handleToggleChange} name="isRunning"/>
          }
        />
        <CardContent>
          <Button variant="text" color="secondary" onClick={this.handleClickDHCP}>
            DHCP
          </Button>
          <Button variant="text" color="secondary" onClick={this.handleClickManual}>
            Manual
          </Button>
          <br/>
          <Typography variant="title">{this.state.useDHCP ? 'DHCP' : 'Manual'}</Typography>
          <br/>
          {!this.state.useDHCP && (
            <div>
              <Typography variant="body2">
                <TextField
                  error={this.state.ipError}
                  required
                  name="ip"
                  label="IP Addr"
                  value={this.state.ip}
                  onChange={this.handleInputChange}
                />
              </Typography>
              <Typography variant="body2">
                <TextField
                  error={this.state.netmaskError}
                  required
                  name="netmask"
                  label="Netmask"
                  value={this.state.netmask}
                  onChange={this.handleInputChange}
                />
              </Typography>
              <Typography variant="body2">
                <TextField
                  error={this.state.gatewayError}
                  required
                  name="gateway"
                  label="Gateway"
                  value={this.state.gateway}
                  onChange={this.handleInputChange}
                />
              </Typography>
              <Button variant="contained" color="primary" onClick={this.handleSet}>
                SET
              </Button>
            </div>
          )}
          {this.state.useDHCP && (
            <div>
              <Typography variant="body2">
                <TextField label="IP Addr" value={this.state.ip} InputProps={{readOnly: true}}/>
              </Typography>
              <Typography variant="body2">
                <TextField label="Netmask" value={this.state.netmask} InputProps={{readOnly: true}}/>
              </Typography>
              <Typography variant="body2">
                <TextField label="Gateway" value={this.state.gateway} InputProps={{readOnly: true}}/>
              </Typography>
              <Button variant="contained" color="primary" onClick={this.handleDHCPC}>
                DHCPC
              </Button>
            </div>
          )}
          <CardActions>
            <Button
              variant="raised"
              component="span"
              color="primary"
              disabled={disabled}
              className={classes.button}
              onClick={this.handleStatus}
            >
              Status
            </Button>
            <Button
              variant="raised"
              component="span"
              color="primary"
              disabled={disabled}
              className={classes.button}
              onClick={this.handleStart}
            >
              Start
            </Button>
            <Button
              variant="raised"
              component="span"
              color="primary"
              disabled={disabled}
              className={classes.button}
              onClick={this.handleStop}
            >
              Stop
            </Button>
            <Button
              variant="raised"
              component="span"
              color="primary"
              disabled={disabled}
              className={classes.button}
              onClick={this.handleRestart}
            >
              Restart
            </Button>
          </CardActions>
        </CardContent>
      </Card>
    );
  }
}

CardInterfaceConfig.propTypes = {
  classes: PropTypes.object.isRequired,
  name: PropTypes.string.isRequired
};

export default withRoot(withStyles(styles)(CardInterfaceConfig));
