import React from "react";
import CardHeader from "@material-ui/core/CardHeader/CardHeader";
import CardContent from "@material-ui/core/CardContent/CardContent";
import TextField from "@material-ui/core/TextField/TextField";
import CardActions from "@material-ui/core/CardActions/CardActions";
import Button from "@material-ui/core/Button/Button";
import CircularProgress from "@material-ui/core/CircularProgress/CircularProgress";
import Card from "@material-ui/core/Card/Card";
import withRoot from "../../withRoot";
import {withStyles} from "@material-ui/core";
import Input from "@material-ui/core/Input/Input";
import {uploadFiles} from "../../api/ApiFile";
import {showNotificationSnack} from "../../mgmt/MgmtNotification";
import green from "@material-ui/core/colors/green";
import CloudUploadIcon from '@material-ui/icons/CloudUpload';

const styles = theme => ({
  card: {
    width: 300,
    height: 280,
    overflow: "auto",
    margin: theme.spacing.unit * 2,
  },
  inputField: {
    maxWidth: "50%",
    minWidth: "30%",
  },
  wrapper: {
    margin: theme.spacing.unit,
    position: 'relative',
  },
  buttonProgress: {
    color: green[500],
    position: 'absolute',
    top: '50%',
    left: '50%',
    marginTop: -12,
    marginLeft: -12,
  },
});

class CardUpload extends React.Component {
  state = {
    dst: "",
    file: null,
    uploadDisabled: false,
  };

  handleDestinationChange = (event) => {
    this.setState({dst: event.target.value});
  };

  handleFileUpload = () => {
    this.setState({uploadDisabled: true});
    uploadFiles(this.state.dst, this.state.file).subscribe(
      p => {
      },
      err => {
        showNotificationSnack(`File upload failed`);
        this.setState({uploadDisabled: false});
      },
      () => {
        showNotificationSnack("File upload success");
        this.setState({uploadDisabled: false});
      },
    );
  };

  handleFileChange = (event) => {
    let files = event.target.files || event.dataTransfer.files;
    if (!files.length) {
      this.setState({file: null});
      return;
    }
    this.setState({file: files[0]});
  };

  render() {
    const {classes} = this.props;

    return (
      <Card className={classes.card}>
        <CardHeader title="File Upload" subheader="Upload Files to Your Device"/>
        <CardContent>
          <Input type="file" className={classes.input} id="upload-file-input"
                 disabled={this.state.uploadDisabled} onChange={this.handleFileChange}
                 multiple fullWidth/>
          <TextField margin="dense" label="To (device side)" helperText="absolute path in device"
                     disabled={this.state.uploadDisabled} value={this.state.dst} fullWidth
                     onChange={this.handleDestinationChange}/>
        </CardContent>
        <CardActions>
          <div className={classes.wrapper}>
            <Button variant="raised" component="span" color='primary' className={classes.button}
                    disabled={this.state.uploadDisabled}
                    onClick={this.handleFileUpload}>
              <CloudUploadIcon/>
              &nbsp;&nbsp;upload
            </Button>
            {this.state.uploadDisabled && <CircularProgress size={24} className={classes.buttonProgress}/>}
          </div>
        </CardActions>
      </Card>
    )
  }
}

export default withRoot(withStyles(styles)(CardUpload));