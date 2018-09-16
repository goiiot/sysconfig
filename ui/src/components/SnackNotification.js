import React from 'react';
import PropTypes from 'prop-types';
import {withStyles} from '@material-ui/core/styles';
import Snackbar from '@material-ui/core/Snackbar';
import IconButton from '@material-ui/core/IconButton';
import CloseIcon from '@material-ui/icons/Close';
import Button from '@material-ui/core/Button/Button';
import withRoot from "../withRoot";

const styles = theme => ({
  close: {
    width: theme.spacing.unit * 4,
    height: theme.spacing.unit * 4,
  },
});

class SnackbarNotification extends React.Component {
  render() {
    const {classes} = this.props;
    return (
      <Snackbar anchorOrigin={{vertical: 'top', horizontal: 'right'}}
                onClose={this.props.closeHandler} autoHideDuration={5000} open={true}
                ContentProps={{'aria-describedby': 'message-id'}}
                message={<span id="message-id">{this.props.msg}</span>}
                action={[
                  this.props.action &&
                  <Button onClick={this.props.onAction} color="secondary" variant="text">
                    {this.props.action}
                  </Button>,
                  <IconButton key="close" aria-label="Close" color="inherit"
                              className={classes.close} onClick={this.props.closeHandler}>
                    <CloseIcon/>
                  </IconButton>
                ]}/>
    );
  }
}

SnackbarNotification.propTypes = {
  classes: PropTypes.object.isRequired,
  action: PropTypes.string,
  onAction: PropTypes.function,
  msg: PropTypes.string.isRequired,
};

export default withRoot(withStyles(styles)(SnackbarNotification));