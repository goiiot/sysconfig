import React from 'react';
import PropTypes from 'prop-types';
import {withStyles} from '@material-ui/core/styles';

import withRoot from '../withRoot';
import {getPeriphList} from "../api/ApiConfigure";
import {showNotificationSnack} from "../mgmt/MgmtNotification";
import CardPeriphConfig from "../components/card/CardPeriphConfig";

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

class PagePeriph extends React.Component {
  state = {
    devices: [],
  };

  componentDidMount() {
    getPeriphList().subscribe(
      list => this.setState({devices: [...list]}),
      err => showNotificationSnack(`Failed to get periph list ${err}`));
  }

  render() {
    const {classes} = this.props;
    const {devices} = this.state;
    return (
      <div className={classes.root}>
        {devices.map((v) => {
          return (
            <CardPeriphConfig key={`periph-item-${v.name}`} name={v.name}/>
          );
        })}
      </div>
    );
  }
};

PagePeriph.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withRoot(withStyles(styles)(PagePeriph));
