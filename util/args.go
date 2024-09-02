package util

import (
	"strconv"

	"github.com/mateothegreat/go-multilog/multilog"
	"github.com/spf13/cobra"
)

func GetArg[T any](cmd *cobra.Command, name string) T {
	v, err := cmd.Flags().GetString(name)
	if err != nil {
		multilog.Fatal("util.GetArg", "failed to get arg", map[string]interface{}{
			"error": err,
		})
	}

	var result T
	switch any(result).(type) {
	case int:
		parsed, _ := strconv.Atoi(v)
		result = any(parsed).(T)
	case float64:
		parsed, _ := strconv.ParseFloat(v, 64)
		result = any(parsed).(T)
	case string:
		result = any(v).(T)
	case bool:
		parsed, _ := strconv.ParseBool(v)
		result = any(parsed).(T)
	default:
		multilog.Fatal("util.getarg", "unsupported type", map[string]interface{}{
			"type": result,
		})
	}

	return result
}
