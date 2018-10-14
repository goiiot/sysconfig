import React from 'react';
import PropTypes from 'prop-types';
import {withStyles} from '@material-ui/core/styles';

import withRoot from '../withRoot';
import {getLoraList} from "../api/ApiConfigure";
import {showNotificationSnack} from "../mgmt/MgmtNotification";
import CardLoraConfig from "../components/card/CardLoraConfig";
import CardLoraBridgeConfig from "../components/card/CardLoraBridgeConfig";

const styles = {
  root: {
    width: '100%',
    height: '100%',
    display: 'flex',
    flexWrap: 'wrap',
    overflow: "auto",
    justifyContent: 'center',
    flexDirection: 'row',
  },
};

class PageLora extends React.Component {
  state = {
    devices: [],
  };

  componentDidMount() {
    getLoraList().subscribe(
      list => this.setState({devices: [...list]}),
      err => showNotificationSnack(`Failed to get periph list ${err}`));
  }

  render() {
    const {classes} = this.props;
    const {devices} = this.state;
    return (
      <div className={classes.root}>
        {devices.map((v) => {
          if (v.type === 'pkt_forwarder') {
            return (
              <CardLoraConfig key={`lora-item-${v.name}`} name={v.name}/>
            );
          } else if (v.type === 'brocaar/lora-gateway-bridge') {
            return (
              <CardLoraBridgeConfig key={`lora-item-${v.name}`} name={v.name}/>
            )
          } else {
            return <div/>;
          }
        })}
      </div>
    );
  }
}

PageLora.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withRoot(withStyles(styles)(PageLora));
