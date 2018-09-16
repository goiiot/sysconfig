import React from "react";
import CardHeader from "@material-ui/core/CardHeader/CardHeader";
import CardActions from "@material-ui/core/CardActions/CardActions";
import Button from "@material-ui/core/Button/Button";
import Card from "@material-ui/core/Card/Card";
import withRoot from "../../withRoot";
import {withStyles} from "@material-ui/core";
import {requestReboot, requestShutdown} from "../../api/ApiPower";

const styles = theme => ({
  card: {
    height: 150,
    width: "auto",
    overflow: "auto",
    margin: theme.spacing.unit * 2,
  },
});

class CardPower extends React.Component {

  handleReboot = () => {
    requestReboot(); // best effort for now
  };

  handleShutdown = () => {
    requestShutdown(); // best effort for now
  };

  render() {
    const {classes} = this.props;

    return (
      <Card className={classes.card}>
        <CardHeader title="Power" subheader="Device Power Management"/>
        <CardActions>
          <Button variant="raised" component="span" color='primary'
                  className={classes.button} onClick={this.handleReboot}>
            Reboot
          </Button>
          <Button variant="raised" component="span" color='secondary'
                  className={classes.button} onClick={this.handleShutdown}>
            Shutdown
          </Button>
        </CardActions>
      </Card>
    );
  }
}

export default withRoot(withStyles(styles)(CardPower));