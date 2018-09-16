import React from "react";
import CardHeader from "@material-ui/core/CardHeader/CardHeader";
import CardContent from "@material-ui/core/CardContent/CardContent";
import Grid from "@material-ui/core/Grid/Grid";
import {supportedFormats} from "../../config";
import FormControlLabel from "@material-ui/core/FormControlLabel/FormControlLabel";
import Radio from "@material-ui/core/Radio/Radio";
import TextField from "@material-ui/core/TextField/TextField";
import CardActions from "@material-ui/core/CardActions/CardActions";
import Button from "@material-ui/core/Button/Button";
import CircularProgress from "@material-ui/core/CircularProgress/CircularProgress";
import Card from "@material-ui/core/Card/Card";
import withRoot from "../../withRoot";
import {withStyles} from "@material-ui/core";
import {downloadFiles} from "../../api/ApiFile";
import {showNotificationSnack} from "../../mgmt/MgmtNotification";
import green from "@material-ui/core/colors/green";
import SaveIcon from '@material-ui/icons/Save';

const styles = theme => ({
  card: {
    width: 300,
    height: 350,
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

class CardDownload extends React.Component {
  state = {
    src: "",
    fmt: "zip",
    downloadDisabled: false,
  };

  handleSourceChange = (event) => {
    this.setState({src: event.target.value});
  };

  handleFormatChange = (event) => {
    this.setState({fmt: event.target.value})
  };

  handleFileDownload = () => {
    this.setState({downloadDisabled: true});
    downloadFiles(this.state.src, this.state.fmt).subscribe(
      p => {
      },
      err => {
        showNotificationSnack("File download failed");
        this.setState({downloadDisabled: false});
      },
      () => {
        showNotificationSnack("File download success");
        this.setState({downloadDisabled: false});
      });
  };

  render() {
    const {classes} = this.props;

    return (
      <Card className={classes.card}>
        <CardHeader title="File Download" subheader="Download Files From Your Device"/>
        <CardContent>
          <Grid container>
            {supportedFormats.map(v => {
              return (
                <Grid item key={`format-radio-${v}`}>
                  <FormControlLabel value={v} control={
                    <Radio checked={this.state.fmt === v}
                           disabled={this.state.downloadDisabled}
                           onChange={this.handleFormatChange}
                           value={v} name='format-radio-group'
                    />} label={v}/>
                </Grid>
              )
            })}
          </Grid>
          <TextField margin="dense" label="From (device side)" helperText="absolute path in device"
                     value={this.state.src} disabled={this.state.downloadDisabled}
                     fullWidth onChange={this.handleSourceChange}/>
        </CardContent>
        <CardActions>
          <div className={classes.wrapper}>
            <Button variant="raised" component="span" color='primary' className={classes.button}
                    disabled={this.state.downloadDisabled} onClick={this.handleFileDownload}>
              <SaveIcon/>
              &nbsp;&nbsp;download
            </Button>
            {this.state.downloadDisabled && <CircularProgress size={24} className={classes.buttonProgress}/>}
          </div>
        </CardActions>
      </Card>
    );
  }
}

export default withRoot(withStyles(styles)(CardDownload));