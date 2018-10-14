import React from "react";
import PropTypes from 'prop-types';
import CardHeader from "@material-ui/core/CardHeader/CardHeader";
import CardActions from "@material-ui/core/CardActions/CardActions";
import Card from "@material-ui/core/Card/Card";
import Button from "@material-ui/core/Button/Button";
import Switch from "@material-ui/core/Switch/Switch";
import withRoot from "../../withRoot";
import {withStyles} from "@material-ui/core";
import {getBusInfo, getBusStatus, startBus, stopBus, updateBusConfig} from "../../api/ApiConfigure";
import {showNotificationDialog, showNotificationSnack} from "../../mgmt/MgmtNotification";

const styles = theme => ({
  card: {
    width: 300,
    height: 350,
    overflow: "auto",
    margin: theme.spacing.unit * 2,
  },
});

class CardDevice extends React.Component {
  state = {
    on: false,
    disabled: true,

    config: {},
  };

  handleStatus = () => {
    const {name} = this.props;
    getBusStatus(name).subscribe(
      resp => showNotificationDialog(`Get ${name} Status success`, resp),
      err => showNotificationSnack(`Failed to get status of ${name} ${err}`, "Retry", this.handleStatus));
  };
  handleUpdateConfig = () => {
    const {name} = this.props;
    updateBusConfig(name, this.state.config).subscribe(
      resp => showNotificationDialog(`Update ${name} success`, resp),
      err => showNotificationSnack(`Failed to update config of ${name} ${err}`, "Retry", this.handleUpdateConfig));
  };
  handleDeviceToggle = () => {
    const on = !this.state.on;
    const {name} = this.props;

    this.setState({disabled: true});
    if (on) {
      startBus(name).subscribe(
        resp => {
          this.setState({on: true});
          showNotificationDialog(`Start ${name} success`, resp);
        },
        err => showNotificationSnack(`Failed to start ${name} ${err}`, "Retry", this.handleStart),
        () => this.setState({disabled: false}));
    } else {
      stopBus(name).subscribe(
        resp => {
          this.setState({on: false});
          showNotificationDialog(`Stop ${name} success`, resp);
        },
        err => showNotificationSnack(`Failed to stop ${name} ${err}`, "Retry", this.handleStop),
        () => this.setState({disabled: false}));
    }
  };

  componentWillMount() {
    const {cmdStart, cmdStop, cmdRestart, cmdStatus, getInfo, updateConfig, getConfig} = this.props;
    this.handleStart = cmdStart;
    this.handleStop = cmdStop;
    this.handleRestart = cmdRestart;
    this.handleStatus = cmdStatus;
  }

  componentDidMount() {
    const {name} = this.props;
    getBusInfo(name).subscribe(
      info => this.setState({disabled: false, config: info.config, status: info.status}),
      err => {
        this.setState({disabled: true});
        showNotificationSnack(`Failed to get config of ${name} ${err}`);
      });
  }

  render() {
    const {classes, name} = this.props;
    const {disabled, config} = this.state;

    return (
      <Card className={classes.card}>
        <CardHeader title={name} action={
          <Switch checked={this.state.on}
                  onClick={this.handleDeviceToggle}/>}/>
        <CardActions>
          <Button variant="raised" component="span" color='secondary'
                  disabled={disabled} className={classes.button}
                  onClick={this.handleStatus}>
            Status
          </Button>
          <Button variant="raised" component="span" color='secondary'
                  disabled={disabled} className={classes.button}
                  onClick={this.handleStart}>
            Start
          </Button>
          <Button variant="raised" component="span" color='secondary'
                  disabled={disabled} className={classes.button}
                  onClick={this.handleStop}>
            Stop
          </Button>
          <Button variant="raised" component="span" color='primary'
                  disabled={disabled} className={classes.button}
                  onClick={this.handleRestart}>
            Restart
          </Button>
        </CardActions>
      </Card>
    );
  }
}

CardDevice.propTypes = {
  classes: PropTypes.object.isRequired,
  name: PropTypes.string.isRequired,
  cmdStart: PropTypes.function,
  cmdStop: PropTypes.function,
  cmdRestart: PropTypes.function,
  cmdStatus: PropTypes.function,
  getInfo: PropTypes.function,
  updateConfig: PropTypes.function,
  getConfig: PropTypes.function,
};

export default withRoot(withStyles(styles)(CardDevice));