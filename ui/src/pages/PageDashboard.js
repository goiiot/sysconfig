import React from 'react';
import PropTypes from 'prop-types';
import withWidth, {isWidthUp} from '@material-ui/core/withWidth';
import {withStyles} from '@material-ui/core/styles';
import Grid from '@material-ui/core/Grid';
import Paper from '@material-ui/core/Paper';

import withRoot from '../withRoot';
import {getBufferedMetrics} from "../mgmt/MgmtMetrics";
import {MetricsStream} from "../mgmt/Streams";
import {Pictures} from "../config";
import {Area, AreaChart, CartesianGrid, Label, Legend, ResponsiveContainer, Tooltip, XAxis, YAxis} from "recharts";

const styles = theme => ({
  root: {
    width: '100%',
    height: '100%',
    display: 'flex',
    flexWrap: 'wrap',
    overflow: "auto",
    justifyContent: 'center',
    flexDirection: 'row',
    alignContent: 'start',
  },
  card: {
    display: "flex",
    justifyContent: "center",
    maxHeight: 300,
    backgroundColor: theme.palette.background.default,
  },
  img: {
    maxWidth: "100%",
    objectFit: "contain",
  }
});

// convert number to string
const getUnitStr = (unit) => {
  switch (unit) {
    case 1:
      return "";
    case 1 << 10:
      return "K";
    case 1 << 20:
      return "M";
    case 1 << 30:
      return "G";
    case 1 << 40:
      return "T";
    case 1 << 50:
      return "P";
    default:
      return "";
  }
};

class PageDashboard extends React.Component {
  state = {
    metrics: [],
    netPktUnit: "",
    netBytesUnit: "B",
  };

  updateMetrics = (metrics) => {
    const newMetrics = [...metrics];

    const threshold = 100;
    let bytesUnit = 1, pktUnit = 1;
    // get suitable unit for pkt and bytes
    newMetrics.forEach((v) => {
      if (!v || !v.net_bytes_recv) {
        return;
      }
      while (v.net_pkt_recv / pktUnit > threshold) {
        pktUnit <<= 10;
      }

      while (v.net_pkt_sent / pktUnit > threshold) {
        pktUnit <<= 10;
      }

      while (v.net_bytes_recv / bytesUnit > threshold) {
        bytesUnit <<= 10;
      }

      while (v.net_bytes_sent / bytesUnit > threshold) {
        bytesUnit <<= 10;
      }
    });

    // apply unit
    newMetrics.forEach((v, i) => {
      if (!v) {
        return
      }

      if (v.scaled) {
        // scaled but the unit get bigger
        if (v.net_bytes_recv > threshold) {
          newMetrics[i].net_bytes_recv = parseFloat((v.net_bytes_recv / 1024).toFixed(2));
          newMetrics[i].net_bytes_sent = parseFloat((v.net_bytes_sent / 1024).toFixed(2));
        }

        if (v.net_bytes_recv > threshold) {
          newMetrics[i].net_pkt_recv = parseFloat((v.net_pkt_recv / 1024).toFixed(2));
          newMetrics[i].net_pkt_sent = parseFloat((v.net_pkt_sent / 1024).toFixed(2));
        }

        return;
      }

      if (v.net_bytes_recv) {
        newMetrics[i].net_bytes_recv = parseFloat((v.net_bytes_recv / bytesUnit).toFixed(2));
        newMetrics[i].net_bytes_sent = parseFloat((v.net_bytes_sent / bytesUnit).toFixed(2));
      }
      if (v.net_pkt_recv) {
        newMetrics[i].net_pkt_recv = parseFloat((v.net_pkt_recv / pktUnit).toFixed(2));
        newMetrics[i].net_pkt_sent = parseFloat((v.net_pkt_sent / pktUnit).toFixed(2));
      }

      newMetrics[i].scaled = true;
    });

    this.setState({
      metrics: newMetrics,
      netBytesUnit: getUnitStr(bytesUnit) + "B",
      netPktUnit: getUnitStr(pktUnit) + "P"
    });
  };

  componentWillMount() {
    MetricsStream.asObservable().subscribe(
      metrics => this.updateMetrics(metrics),
      err => console.log(err),
      () => console.log("metrics stream complete"));
  }

  componentDidMount() {
    this.setState({metrics: getBufferedMetrics()});
  }

  render() {
    const {classes} = this.props;
    const CpuMem = <Grid item xs={10} sm={6} md={4} xl={3}>
      <ResponsiveContainer width="100%" height={300}>
        <AreaChart data={this.state.metrics} syncId="metrics-sync"
                   margin={{top: 10, right: 30, left: 0, bottom: 0}}>
          <Legend verticalAlign="bottom" height={48}/>
          <XAxis dataKey="time" tickCount={20}>
            <Label offset={0} position="insideBottom"/>
          </XAxis>
          <YAxis unit="%"
                 domain={[0, 100]}/>
          <CartesianGrid strokeDasharray="3 3"/>
          <Tooltip/>
          <Area isAnimationActive={false} type="monotone" dataKey="mem"
                stroke="#2196F3" fillOpacity={0.3} fill="#2196F3"/>
          <Area isAnimationActive={false} type="monotone" dataKey="cpu"
                stroke="#FF5722" fillOpacity={0.3} fill="#FF5722"/>
        </AreaChart>
      </ResponsiveContainer>
    </Grid>;
    const FeatureImg = <Grid item sm={12} md={4}>
      <Paper className={classes.card} elevation={0}>
        <img className={classes.img} src={Pictures.feature_img} alt=""/>
      </Paper>
    </Grid>;
    const Disk = <Grid item xs={10} sm={6} md={4} xl={3}>
      <ResponsiveContainer width="100%" height={300}>
        <AreaChart data={this.state.metrics} syncId="metrics-sync"
                   margin={{top: 10, right: 30, left: 0, bottom: 0}}>
          <Legend verticalAlign="bottom" height={48}/>
          <XAxis dataKey="time" tickCount={20}>
            <Label offset={0} position="insideBottom"/>
          </XAxis>
          <YAxis unit="%"
                 domain={[0, 100]}/>
          <CartesianGrid strokeDasharray="3 3"/>
          <Tooltip/>
          <Area isAnimationActive={false} type="monotone" dataKey="disk"
                stroke="#81C784" fillOpacity={0.3} fill="#81C784"/>
        </AreaChart>
      </ResponsiveContainer>
    </Grid>;
    const NetPkt = <Grid item xs={10} sm={6} md={4} xl={3}>
      <ResponsiveContainer width="100%" height={300}>
        <AreaChart data={this.state.metrics}
                   syncId="metrics-sync"
                   margin={{top: 10, right: 30, left: 0, bottom: 0}}>
          <Legend verticalAlign="bottom" height={48}/>
          <XAxis dataKey="time" tickCount={20}>
            <Label offset={0} position="insideBottom"/>
          </XAxis>
          <YAxis unit={this.state.netPktUnit}/>
          <CartesianGrid strokeDasharray="3 3"/>
          <Tooltip/>
          <Area isAnimationActive={false} type="monotone" dataKey="net_pkt_sent"
                stroke="#81C784" fillOpacity={0.3} fill="#81C784"/>
          <Area isAnimationActive={false} type="monotone" dataKey="net_pkt_recv"
                stroke="#2196F3" fillOpacity={0.3} fill="#2196F3"/>
        </AreaChart>
      </ResponsiveContainer>
    </Grid>;
    const NetBytes = <Grid item xs={10} sm={6} md={4} xl={3}>
      <ResponsiveContainer width="100%" height={300}>
        <AreaChart data={this.state.metrics} syncId="metrics-sync"
                   margin={{top: 10, right: 30, left: 0, bottom: 0}}>
          <Legend verticalAlign="bottom" height={48}/>
          <XAxis dataKey="time" tickCount={20}>
            <Label offset={0} position="insideBottom"/>
          </XAxis>
          <YAxis unit={this.state.netBytesUnit}/>
          <CartesianGrid strokeDasharray="3 3"/>
          <Tooltip/>
          <Area isAnimationActive={false} type="monotone" dataKey="net_bytes_sent"
                stroke="#81C784" fillOpacity={0.3} fill="#81C784"/>
          <Area isAnimationActive={false} type="monotone" dataKey="net_bytes_recv"
                stroke="#2196F3" fillOpacity={0.3} fill="#2196F3"/>
        </AreaChart>
      </ResponsiveContainer>
    </Grid>;

    if (isWidthUp('md', this.props.width)) {
      return (
        <div className={classes.root}>
          <Grid container direction="row" alignItems="center" justify="center">
            {CpuMem}
            {FeatureImg}
            {Disk}
            {NetPkt}
            {FeatureImg}
            {NetBytes}
          </Grid>
        </div>
      )
    }

    return (
      <div className={classes.root}>
        <Grid container direction="row" alignItems="center" justify="center">
          {FeatureImg}
          {CpuMem}
          {Disk}
          {NetPkt}
          {NetBytes}
        </Grid>
      </div>
    );
  }
}

PageDashboard.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withWidth()(withRoot(withStyles(styles)(PageDashboard)));
