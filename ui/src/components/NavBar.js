import React from 'react';
import PropTypes from 'prop-types';
import classNames from 'classnames';
import {withStyles} from '@material-ui/core/styles';
import SwipeableDrawer from '@material-ui/core/SwipeableDrawer';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import MenuItem from '@material-ui/core/MenuItem';
import Menu from '@material-ui/core/Menu';
import IconButton from '@material-ui/core/IconButton';
import Divider from '@material-ui/core/Divider';
import List from "@material-ui/core/List/List";
import ListItemIcon from "@material-ui/core/ListItemIcon/ListItemIcon";
import ListItemText from "@material-ui/core/ListItemText/ListItemText";
import ListItem from "@material-ui/core/ListItem/ListItem";
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Switch from '@material-ui/core/Switch';
import Typography from '@material-ui/core/Typography';
// icon
import MenuIcon from '@material-ui/icons/Menu';
import AccountCircle from '@material-ui/icons/AccountCircle';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import ChevronRightIcon from '@material-ui/icons/ChevronRight';
// pages
import {navItems} from "../config";
import {startCollectMetrics, stopCollectMetrics} from "../mgmt/MgmtMetrics";
import {getPreference} from "../mgmt/MgmtStore";
import {PrefStream, VersionStream} from "../mgmt/Streams";
import DialogPref from "./dialog/DialogPref";
import Paper from "@material-ui/core/Paper/Paper";
import PageDashboard from "../pages/PageDashboard";

const drawerWidth = 240;

const styles = theme => ({
  root: {
    width: '100%',
    height: '100%',
    flexGrow: 1,
    zIndex: 1,
    overflow: 'hidden',
    position: 'relative',
    display: 'flex',
  },
  appFrame: {
    position: 'relative',
    display: 'flex',
    width: '100%',
    height: '100%',
  },
  appBar: {
    zIndex: theme.zIndex.drawer + 1,
    transition: theme.transitions.create(['width', 'margin'], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen,
    }),
  },
  appBarShift: {
    marginLeft: drawerWidth,
    width: `calc(100% - ${drawerWidth}px)`,
    transition: theme.transitions.create(['width', 'margin'], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen,
    }),
  },
  flex: {
    flexGrow: 1,
  },
  menuButton: {
    marginLeft: 12,
    marginRight: 36,
  },
  hide: {
    display: 'none',
  },
  drawerPaper: {
    position: 'relative',
    whiteSpace: 'nowrap',
    width: drawerWidth,
    transition: theme.transitions.create('width', {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen,
    }),
  },
  drawerPaperClose: {
    overflowX: 'hidden',
    transition: theme.transitions.create('width', {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen,
    }),
    width: theme.spacing.unit * 7,
    [theme.breakpoints.up('sm')]: {
      width: theme.spacing.unit * 9,
    },
  },
  versionBar: {
    height: 20,
    position: 'absolute',
    display: "flex",
    // flex: "center",
    justifyContent: "center",
    // padding: 5,
    bottom: 0,
    left: 0,
    right: 0,
  },
  toolbar: {
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'flex-end',
    padding: '0 8px',
    ...theme.mixins.toolbar,
  },
  toolbarText: {
    color: "#FFFFFF",
  },
  content: {
    height: `calc(100% - 56px)`,
    [theme.breakpoints.up('sm')]: {
      height: `calc(100% - 64px)`,
    },
    flexGrow: 1,
    backgroundColor: theme.palette.background.default,
    padding: 0,
  },
});

class NavBar extends React.Component {
  state = {
    collectMetrics: false,
    refresh_interval: Infinity,

    profileAnchorEl: null,
    prefDialogOpen: false,

    drawerOpen: false,
    page: PageDashboard,

    version: "",
    buildTime: "",
    commit: "",
    goVersion: "",
  };

  handleNavItemClick = (page) => () => this.setState({page: page});
  handleProfileMenuToggle = event => {
    const target = this.state.profileAnchorEl == null ? event.target : null;
    this.setState({profileAnchorEl: target});
  };

  handlePrefDialogToggle = () => {
    const open = !this.state.prefDialogOpen;
    this.setState({profileAnchorEl: null, prefDialogOpen: open});
  };

  handleDrawerToggle = () => {
    const open = !this.state.drawerOpen;
    this.setState({drawerOpen: open});
  };

  toggleMetricCollection = () => {
    const collecting = !this.state.collectMetrics;
    if (collecting) {
      startCollectMetrics(this.state.refresh_interval);
    } else {
      stopCollectMetrics();
    }

    this.setState({collectMetrics: collecting});
  };
  updatePrefRelated = (pref) => {
    this.setState({refresh_interval: pref.monitoring.refresh_interval});

    if (this.state.collectMetrics) {
      stopCollectMetrics();
      startCollectMetrics(pref.monitoring.refresh_interval);
    }
  };

  updateVersionInfo = (v) => this.setState({
    version: v.version,
    buildTime: v.build_time,
    commit: v.commit,
    goVersion: v.go_version
  });

  componentDidMount() {
    // apply preference first
    this.updatePrefRelated(getPreference());

    PrefStream.asObservable().subscribe(
      pref => this.updatePrefRelated(pref),
      err => console.log(err),
      () => console.log("PrefStream Complete"));

    VersionStream.asObservable().subscribe(
      v => this.updateVersionInfo(v),
      err => console.log(err),
      () => console.log("VersionStream Complete"));
  }

  render() {
    const {classes, theme} = this.props;
    const {profileAnchorEl} = this.state;
    const profileOpen = Boolean(profileAnchorEl);
    return (
      <div className={classes.root}>
        <div className={classes.appFrame}>
          <AppBar
            position="absolute"
            className={classNames(classes.appBar, this.state.drawerOpen && classes.appBarShift)}>
            <Toolbar disableGutters={!this.state.drawerOpen}>
              <IconButton
                color="inherit"
                aria-label="Open drawer"
                onClick={this.handleDrawerToggle}
                className={classNames(classes.menuButton, this.state.drawerOpen && classes.hide)}>
                <MenuIcon/>
              </IconButton>
              <Typography variant="title" color="inherit" noWrap className={classes.flex}>
                Configuration
              </Typography>
              <div>
                <FormControlLabel labelPlacement="start"
                                  control={<Switch checked={this.state.collectMetrics}
                                                   onChange={this.toggleMetricCollection}
                                                   value="Monitoring" color="secondary"/>}
                                  label={(<Typography
                                    className={classes.toolbarText}>{this.state.collectMetrics ? "Monitoring" : "Not Monitoring"}</Typography>)}/>
                <IconButton aria-owns={profileOpen ? 'menu-appbar' : null}
                            aria-haspopup="true" color="inherit" onClick={this.handleProfileMenuToggle}>
                  <AccountCircle/>
                </IconButton>
                <Menu id="menu-appbar" open={profileOpen} anchorEl={profileAnchorEl}
                      anchorOrigin={{vertical: 'top', horizontal: 'right'}}
                      transformOrigin={{vertical: 'top', horizontal: 'right'}}
                      onClose={this.handleProfileMenuToggle}>
                  <MenuItem onClick={this.handlePrefDialogToggle}>Preference</MenuItem>
                  <MenuItem onClick={this.handleProfileMenuToggle}>Sign Out</MenuItem>
                </Menu>
              </div>
            </Toolbar>
          </AppBar>
          <SwipeableDrawer variant="permanent" onClose={this.handleDrawerToggle} onOpen={this.handleDrawerToggle}
                           open={this.state.drawerOpen}
                           classes={{paper: classNames(classes.drawerPaper, !this.state.drawerOpen && classes.drawerPaperClose)}}>
            <div className={classes.toolbar}>
              <IconButton onClick={this.handleDrawerToggle}>
                {theme.direction === 'rtl' ? <ChevronRightIcon/> : <ChevronLeftIcon/>}
              </IconButton>
            </div>
            <Divider/>
            <List>
              {navItems.map(v => {
                return (
                  <ListItem button onClick={this.handleNavItemClick(v.page)} key={`nav-item-${v.title}`}>
                    <ListItemIcon>
                      <v.icon/>
                    </ListItemIcon>
                    <ListItemText primary={v.title}/>
                  </ListItem>
                )
              })}
            </List>
            <Divider/>
            <Paper square={true} className={classes.versionBar}>
              <Typography variant="body1">
                {this.state.drawerOpen ? `${this.state.version} ${this.state.commit} ${this.state.goVersion}` : this.state.version}
              </Typography>
            </Paper>
          </SwipeableDrawer>
          <main className={classes.content}>
            <div className={classes.toolbar}/>
            {(<this.state.page/>)}
          </main>
        </div>
        <DialogPref open={this.state.prefDialogOpen} closeHandler={this.handlePrefDialogToggle}/>
      </div>
    );
  }
}

NavBar.propTypes = {
  classes: PropTypes.object.isRequired,
  theme: PropTypes.object.isRequired,
};

export default withStyles(styles, {withTheme: true})(NavBar);