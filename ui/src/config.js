// icons
import WifiIcon from '@material-ui/icons/Wifi';
import SettingsEthernetIcon from '@material-ui/icons/SettingsEthernet';
import BuildIcon from '@material-ui/icons/Build';
import LoraIcon from './icons/LoraIcon';
import SerialIcon from './icons/SerialIcon';
import DashboardIcon from './icons/DashboardIcon';
// pages
import PageTools from "./pages/PageTools";
import PageBus from "./pages/PageBus";
import PageConfigLoraRadio from "./pages/PageLora";
import PagePeriph from "./pages/PagePeriph";
import PageNetwork from "./pages/PageNetwork";
import PageDashboard from "./pages/PageDashboard";

let debug = process.env.NODE_ENV !== 'production';
const debugApiSrv = "localhost:8080";

export const apiConfig = {
  apiServer: debug ? debugApiSrv : window.location.host,
  maxRedirects: 10,
  apiVersion: "v1",
};

export const Pictures = {
  feature_img: "./static/pic/feature.png",
};

export const navItems = [
  {title: "Dashboard", icon: DashboardIcon, page: PageDashboard},
  {title: "Lora", icon: LoraIcon, page: PageConfigLoraRadio},
  {title: "Network", icon: WifiIcon, page: PageNetwork},
  {title: "Bus", icon: SettingsEthernetIcon, page: PageBus},
  {title: "Periph", icon: SerialIcon, page: PagePeriph},
  {title: "Tools", icon: BuildIcon, page: PageTools},
];

export const supportedFormats = [
  "zip", "tar", "tar.gz", "tar.bz2", "tar.xz"
];