const WIDGET_NONE = "none";
const WIDGET_INPUT = "input";
const WIDGET_SWITCH = "switch";
const WIDGET_SECTION = "section";

const getFieldWidgetType = (field) => {
  if (!field) return "";
  switch (typeof field) {
    case "boolean":
      return WIDGET_SWITCH;
    case "string":
    case "number":
      return WIDGET_INPUT;
    case "object":
      return WIDGET_SECTION;
    default:
      return WIDGET_NONE;
  }
};

const convertFieldToWidget = (conf) => {
  const keys = Object.keys(conf);
  console.log(keys);
};