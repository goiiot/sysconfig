package wifi
q
var funcMap = map[string]interface{}{
	"HasValue": hasValue,
}

func hasValue(o interface{}) bool {
	switch o.(type) {
	case bool:
		return o.(bool)
	case []uint:
		return o != nil || len(o.([]uint)) != 0
	case []uint8:
		return o != nil || len(o.([]uint8)) != 0
	case []uint16:
		return o != nil || len(o.([]uint16)) != 0
	case []uint32:
		return o != nil || len(o.([]uint32)) != 0
	case []uint64:
		return o != nil || len(o.([]uint64)) != 0
	case []int:
		return o != nil || len(o.([]int)) != 0
	case []int8:
		return o != nil || len(o.([]int8)) != 0
	case []int16:
		return o != nil || len(o.([]int16)) != 0
	case []int32:
		return o != nil || len(o.([]int32)) != 0
	case []int64:
		return o != nil || len(o.([]int64)) != 0
	case []string:
		return o != nil || len(o.([]string)) != 0
	case uint:
		return o.(uint) != 0
	case uint8:
		return o.(uint8) != 0
	case uint16:
		return o.(uint16) != 0
	case uint32:
		return o.(uint32) != 0
	case uint64:
		return o.(uint64) != 0
	case int:
		return o.(int) != 0
	case int8:
		return o.(int8) != 0
	case int16:
		return o.(int16) != 0
	case int32:
		return o.(int32) != 0
	case int64:
		return o.(int64) != 0
	case float32:
		return o.(float32) != 0
	case float64:
		return o.(float32) != 0
	case string:
		return len(o.(string)) != 0
	default:
		return o != nil
	}
}
