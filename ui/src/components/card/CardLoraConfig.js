import React from 'react';
import PropTypes from 'prop-types';
import classnames from 'classnames';
import CardHeader from '@material-ui/core/CardHeader/CardHeader';
import CardActions from '@material-ui/core/CardActions/CardActions';
import Card from '@material-ui/core/Card/Card';
import Button from '@material-ui/core/Button/Button';
import withRoot from '../../withRoot';
import {withStyles} from '@material-ui/core';
import {requestReboot} from '../../api/ApiPower';
import {getLoraConfig, getLoraStatus, restartLora, startLora, stopLora, updateLoraConfig} from '../../api/ApiConfigure';
import {showNotificationDialog, showNotificationSnack} from '../../mgmt/MgmtNotification';
import Typography from '@material-ui/core/Typography/Typography';
import TextField from '@material-ui/core/TextField/TextField';
import Checkbox from '@material-ui/core/Checkbox/Checkbox';
import CardContent from '@material-ui/core/CardContent/CardContent';
import Collapse from '@material-ui/core/Collapse/Collapse';
import IconButton from '@material-ui/core/IconButton/IconButton';
import Divider from '@material-ui/core/Divider/Divider';
import Switch from '@material-ui/core/Switch/Switch';

import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import RefreshIcon from '@material-ui/icons/Refresh';
import CachedIcon from '@material-ui/icons/Cached';
import {loraDefaultConfig} from '../../template';

const styles = (theme) => ({
  card: {
    height: 'auto',
    width: 400,
    overflow: 'auto',
    margin: theme.spacing.unit * 2
  }
});

class CardLoraConfig extends React.Component {
  state = {
    disabled: true,
    config: loraDefaultConfig,

    isRunning: true,
    gwid: '00049fFFFE011ba9',
    isLoRaWanPublic: true
  };

  handleStart = () => {
    const {name} = this.props;
    startLora(name).subscribe(
      (resp) => showNotificationDialog(`Start ${name} success`, resp),
      (err) => showNotificationSnack(`Failed to start ${name} ${err}`, 'Retry', this.handleStart)
    );
    requestReboot();
  };

  handleStop = () => {
    const {name} = this.props;
    stopLora(name).subscribe(
      (resp) => showNotificationDialog(`Stop ${name} success`, resp),
      (err) => showNotificationSnack(`Failed to stop ${name} ${err}`, 'Retry', this.handleStop)
    );
  };

  handleRestart = () => {
    const {name} = this.props;
    restartLora(name).subscribe(
      (resp) => showNotificationDialog(`Restart ${name} success`, resp),
      (err) => showNotificationSnack(`Failed to restart ${name} ${err}`, 'Retry', this.handleRestart)
    );
  };

  handleStatus = () => {
    const {name} = this.props;
    getLoraStatus(name).subscribe(
      (resp) => showNotificationDialog(`Get ${name} Status success`, resp),
      (err) => showNotificationSnack(`Failed to get status of ${name} ${err}`, 'Retry', this.handleStatus)
    );
  };

  handleUpdateConfig = () => {
    const {name} = this.props;
    updateLoraConfig(name, this.state.config).subscribe(
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

  handleResetChip = () => {
    if (this.state.isRunning) {
      // alert('please stop');
      // return;
    }
    // alert('reset ok');
  };

  handleUpdateGWID = () => {
    this.setState({gwid: 'aaaa'});
  };

  handleExpandClick = () => {
    this.setState((state) => ({expanded: !state.expanded}));
  };

  componentDidMount() {
    const {name} = this.props;
    getLoraConfig(name).subscribe(
      (c) => {
        console.log(c);
        this.setState({disabled: false, config: c});
      },
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
          subheader={'forward packet to mqtt broker'}
        />
        <CardContent>
          <Typography variant="body2">
            SX1301
            <Button color="primary" onClick={this.handleResetChip}>
              RESET
              <RefreshIcon/>
            </Button>
          </Typography>
          <Typography variant="body2">
            Update ID
            <IconButton aria-label="gwid_update" onClick={this.handleUpdateGWID}>
              <CachedIcon color="primary"/>
            </IconButton>
          </Typography>
          <br/>
          <TextField label="Gateway Id" value={this.state.gwid} InputProps={{readOnly: true}}/>
          <Divider/>
          <Typography variant="title">Radio Information</Typography>
          <Typography variant="body1" color="textPrimary">
            Radio 0 ({config.SX1301_conf.radio_0.freq / 1000000} MHz)
          </Typography>
          <Typography variant="body2" color="textPrimary">
            Type: {config.SX1301_conf.radio_0.type}, EN:{config.SX1301_conf.radio_0.enable ? '√' : '×'}
          </Typography>
          <Typography variant="body2" color="textPrimary">
            TX: {config.SX1301_conf.radio_0.tx_enable ? '√' : '×'}
            {config.SX1301_conf.radio_0.tx_enable && (
              <React.Fragment>
                {' |'}Range:{config.SX1301_conf.radio_0.tx_freq_min / 1000000} - {' '}
                {config.SX1301_conf.radio_0.tx_freq_max / 1000000} MHz
              </React.Fragment>
            )}
          </Typography>
          <Divider/>
          <Typography variant="body1" color="textPrimary">
            Radio 1 ({config.SX1301_conf.radio_1.freq / 1000000} MHz)
          </Typography>
          <Typography variant="body2" color="textPrimary">
            Type: {config.SX1301_conf.radio_1.type}, EN:{config.SX1301_conf.radio_1.enable ? '√' : '×'}
          </Typography>
          <Typography variant="body2" color="textPrimary">
            TX: {config.SX1301_conf.radio_1.tx_enable ? '√' : '×'}
            {config.SX1301_conf.radio_1.tx_enable && (
              <React.Fragment>
                {' |'}Range: {config.SX1301_conf.radio_1.tx_freq_min / 1000000} - {' '}
                {config.SX1301_conf.radio_1.tx_freq_max / 1000000} MHz
              </React.Fragment>
            )}
          </Typography>
          <Divider/>
          <IconButton
            classes={classnames(classes.expand, {[classes.expandOpen]: this.state.expanded})}
            onClick={this.handleExpandClick}
            aria-expanded={this.state.expanded}
            aria-label="Show more"
          >
            <ExpandMoreIcon/>
          </IconButton>
          <Collapse in={this.state.expanded} timeout="auto" unmountOnExit>
            <Typography variant="body2">
              LoRaWAN Public
              <Checkbox
                color="secondary"
                checked={this.state.isLoRaWanPublic}
                onChange={this.handleToggleChange}
              />
            </Typography>
          </Collapse>
        </CardContent>
        <CardActions>
          <Button
            variant="raised"
            component="span"
            color="primary"
            disabled={disabled}
            className={classes.button}
            onClick={this.handleUpdateConfig}
          >
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
