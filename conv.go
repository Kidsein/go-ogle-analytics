package ga

import (
	"fmt"
	"strconv"
)

func bool2str(val bool) string {
	if val {
		return "1"
	} else {
		return "0"
	}
}

func int2str(val int64) string {
	return fmt.Sprintf("%d", val)
}

func currency2str(val float64) string {
	return fmt.Sprintf("%.6f", val)
}

func number2str(val float64) string {
	return strconv.FormatFloat(val, 'f', -1, 64)
}
