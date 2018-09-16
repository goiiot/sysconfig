import React from 'react';
import PropTypes from 'prop-types';
import Button from '@material-ui/core/Button';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';
import withRoot from "../../withRoot";
import {withStyles} from "@material-ui/core";

const styles = {};

class DialogNotification extends React.Component {

  render() {
    const {closeHandler, title, msg, cancel, onCancel, ok, onOk} = this.props;
    return (
      <Dialog open={true}
              onClose={closeHandler}
              aria-labelledby="alert-dialog-title" aria-describedby="alert-dialog-description">
        <DialogTitle id="alert-dialog-title">{`${title}`}</DialogTitle>
        <DialogContent>
          <DialogContentText id="alert-dialog-description">
            {`${msg}`}
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          {cancel &&
          <Button onClick={onCancel} color="primary" autoFocus={ok === null}>
            {cancel}
          </Button>}
          {ok &&
          <Button onClick={onOk} color="primary" autoFocus>
            {ok}
          </Button>}
          {!ok && !cancel &&
          <Button onClick={closeHandler} color="primary" autoFocus>
            OK
          </Button>}
        </DialogActions>
      </Dialog>
    );
  }
}

DialogNotification.propTypes = {
  classes: PropTypes.object.isRequired,
  title: PropTypes.string.isRequired,
  msg: PropTypes.string.isRequired,
};

export default withRoot(withStyles(styles)(DialogNotification));