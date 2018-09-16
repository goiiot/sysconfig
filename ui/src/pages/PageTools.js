import React from 'react';
import PropTypes from 'prop-types';
import {withStyles} from '@material-ui/core/styles';
import withRoot from '../withRoot';
import CardUpload from "../components/card/CardUpload";
import CardDownload from "../components/card/CardDownload";
import CardPower from "../components/card/CardPower";

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

class PageTools extends React.Component {

  render() {
    const {classes} = this.props;
    return (
      <div className={classes.root}>
        <CardUpload/>
        <CardDownload/>
        <CardPower/>
      </div>
    );
  }
}

PageTools.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withRoot(withStyles(styles)(PageTools));
