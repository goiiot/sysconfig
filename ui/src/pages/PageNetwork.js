import React from 'react';
import PropTypes from 'prop-types';
import {withStyles} from '@material-ui/core/styles';

import withRoot from '../withRoot';
import {getCellularList, getInterfaceList, getWifiList} from "../api/ApiConfigure";
import {showNotificationSnack} from "../mgmt/MgmtNotification";
import CardInterfaceConfig from "../components/card/CardInterfaceConfig";
import CardCellularConfig from "../components/card/CardCellularConfig";
import CardWifiConfig from "../components/card/CardWifiConfig";

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

class PageNetwork extends React.Component {
  state = {
    ifaceList: [],
    cellList: [],
    wifiList: [],
  };

  componentDidMount() {
    getInterfaceList().subscribe(
      list => this.setState({ifaceList: [...list]}),
      err => showNotificationSnack(`Failed to get interface list ${err}`));

    getCellularList().subscribe(
      list => this.setState({cellList: [...list]}),
      err => showNotificationSnack(`Failed to get cellular list ${err}`));

    getWifiList().subscribe(
      list => this.setState({wifiList: [...list]}),
      err => showNotificationSnack(`Failed to get wifi list ${err}`));
  }

  render() {
    const {classes} = this.props;
    const {ifaceList, cellList, wifiList} = this.state;
    return (
      <div className={classes.root}>
        {ifaceList.map((v) => {
          return (
            <CardInterfaceConfig key={`iface-item-${v.name}`} name={v.name}/>
          );
        })}
        {cellList.map((v) => {
          return (
            <CardCellularConfig key={`cell-item-${v.name}`} name={v.name}/>
          )
        })}
        {wifiList.map((v) => {
          return (
            <CardWifiConfig key={`wifi-item-${v.name}`} name={v.name}/>
          )
        })}
      </div>
    );
  }
};

PageNetwork.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withRoot(withStyles(styles)(PageNetwork));
