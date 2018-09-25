import React from 'react';
import PropTypes from 'prop-types';
import CardHeader from '@material-ui/core/CardHeader/CardHeader';
import CardActions from '@material-ui/core/CardActions/CardActions';
import Card from '@material-ui/core/Card/Card';
import Button from '@material-ui/core/Button/Button';
import Switch from '@material-ui/core/Switch/Switch';
import withRoot from '../../withRoot';
import { withStyles, CardContent } from '@material-ui/core';
import {
	getPeriphConfig,
	getPeriphInfo,
	getPeriphStatus,
	restartPeriph,
	startPeriph,
	stopPeriph,
	updatePeriphConfig
} from '../../api/ApiConfigure';
import { showNotificationDialog, showNotificationSnack } from '../../mgmt/MgmtNotification';
import IconButton from '@material-ui/core/IconButton/IconButton';
import RefreshIcon from '@material-ui/icons/Refresh';
import TextField from '@material-ui/core/TextField';

const styles = (theme) => ({
	card: {
		height: 'auto',
		width: 400,
		overflow: 'auto',
		margin: theme.spacing.unit * 2
	}
});

class CardPeriphConfig extends React.Component {
	state = {
		on: false,
		disabled: true,

		inputError: false,

		mqServer: '',
		mqPort: 0,
		qos: 0,

		config: {}, // TODO split config to multi parts
		status: {} // TODO split status to multi parts
	};

	handleRestart = () => {
		const { name } = this.props;
		this.setState({ disabled: true });

		restartPeriph(name).subscribe(
			(resp) => {
				this.handleStatus();
				showNotificationDialog(`Restart ${name} success`, resp);
			},
			(err) => showNotificationSnack(`Failed to restart ${name} ${err}`, 'Retry', this.handleRestart)
		);
	};

	handleStatus = () => {
		const { name } = this.props;
		this.setState({ disabled: true });

		getPeriphStatus(name).subscribe(
			(status) => this.updateDisplayStatus(status),
			(err) => showNotificationSnack(`Failed to get status of ${name} ${err}`, 'Retry', this.handleStatus),
			() => this.setState({ disabled: false })
		);
	};

	getDisplayConfig = (cfg) => {
		// TODO make displayed config wholesome
		return this.state.config;
	};

	requestUpdateConfig = () => {
		const { name } = this.props;
		this.setState({ disabled: true });

		updatePeriphConfig(name, this.getDisplayConfig()).subscribe(
			(r) => {
				// get update result via get device info
				this.handleDeviceInfo();
				showNotificationDialog(`Update ${name} config success`);
			},
			(err) =>
				showNotificationSnack(`Failed to update config of ${name} ${err}`, 'Retry', this.requestUpdateConfig),
			() => this.setState({ disabled: false })
		);
	};

	handleDeviceToggle = () => {
		const on = !this.state.on;
		const { name } = this.props;
		this.setState({ disabled: true });

		if (on) {
			startPeriph(name).subscribe(
				(resp) => {
					this.handleStatus();
					showNotificationDialog(`Start ${name} success`, resp);
				},
				(err) => showNotificationSnack(`Failed to start ${name} ${err}`, 'Retry', this.handleStart),
				() => this.setState({ disabled: false })
			);
		} else {
			stopPeriph(name).subscribe(
				(resp) => {
					this.handleStatus();
					showNotificationDialog(`Stop ${name} success`, resp);
				},
				(err) => showNotificationSnack(`Failed to stop ${name} ${err}`, 'Retry', this.handleStop),
				() => this.setState({ disabled: false })
			);
		}
	};

	updateDisplayConfig = (config) => {
		this.setState({ config: config });
	};

	updateDisplayStatus = (status) => {
		this.setState({ on: status.on, status: status });
	};

	handleDeviceInfo = () => {
		const { name } = this.props;
		this.setState({ disabled: true });

		getPeriphInfo(name).subscribe(
			(info) => {
				this.updateDisplayStatus(info.status);
				this.updateDisplayConfig(info.config);
				this.setState({ disabled: false });
			},
			(err) => {
				this.setState({ disabled: true });
				showNotificationSnack(`Failed to get config of ${name} ${err}`);
			}
		);
	};

	componentDidMount() {
		this.handleDeviceInfo();
	}

	handleQoSChange = (event) => {
		const value = event.target.value;
		this.setState({
			qos: value < 0 ? 0 : value > 2 ? 2 : value
		});
	};

	handleInputChange = (name) => (event) => {
		const value = event.target.value;
		this.setState({
			[name]: value
		});
		if (name === 'mqServer') {
			var testRegExp = /^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$/;
			testRegExp.test(value) ? this.setState({ inputError: false }) : this.setState({ inputError: true });
		}
	};

	render() {
		const { classes, name } = this.props;
		const { disabled, config } = this.state;

		return (
			<Card className={classes.card}>
				<CardHeader
					title={name}
					action={[
						<IconButton
							key={`card-periph-${name}-refresh`}
							color="default"
							aria-label="Refresh"
							onClick={this.handleDeviceInfo}
						>
							<RefreshIcon />
						</IconButton>,
						<Switch
							key={`card-periph-${name}-switch`}
							checked={this.state.on}
							disabled={this.state.disabled}
							onClick={this.handleDeviceToggle}
						/>
					]}
				/>
				<CardContent>
					<TextField
						error={this.state.inputError}
						required
						label="mqtt server"
						value={this.state.mqServer}
						onChange={this.handleInputChange('mqServer')}
					/>
					<br />
					<TextField
						required
						type="number"
						label="port"
						value={this.state.mqPort}
						onChange={this.handleInputChange('mqPort')}
					/>
					<br />
					<TextField
						required
						type="number"
						label="QoS"
						value={this.state.qos}
						onChange={this.handleQoSChange}
					/>
				</CardContent>
				<CardActions>
					<Button
						variant="raised"
						component="span"
						color="primary"
						disabled={disabled}
						className={classes.button}
						onClick={this.requestUpdateConfig}
					>
						Save
					</Button>
				</CardActions>
			</Card>
		);
	}
}

CardPeriphConfig.propTypes = {
	classes: PropTypes.object.isRequired,
	name: PropTypes.string.isRequired
};

export default withRoot(withStyles(styles)(CardPeriphConfig));
