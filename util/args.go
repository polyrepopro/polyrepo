package util

import (
	"reflect"

	"github.com/mateothegreat/go-multilog/multilog"
	"github.com/spf13/cobra"
)

func GetArg[T any](cmd *cobra.Command, name string) T {
	var result T
	resultType := reflect.TypeOf(result)
	switch resultType.Kind() {
	case reflect.Int:
		v, err := cmd.Flags().GetInt(name)
		if err != nil {
			multilog.Fatal("util.GetArg", "failed to get arg", map[string]interface{}{
				"error": err,
				"kind":  resultType.Kind(),
			})
		}
		result = any(v).(T)
	case reflect.Float64:
		v, err := cmd.Flags().GetFloat64(name)
		if err != nil {
			multilog.Fatal("util.GetArg", "failed to get arg", map[string]interface{}{
				"error": err,
				"kind":  resultType.Kind(),
			})
		}
		result = any(v).(T)
	case reflect.String:
		v, err := cmd.Flags().GetString(name)
		if err != nil {
			multilog.Fatal("util.GetArg", "failed to get arg", map[string]interface{}{
				"error": err,
				"kind":  resultType.Kind(),
			})
		}
		result = any(v).(T)
	case reflect.Bool:
		v, err := cmd.Flags().GetBool(name)
		if err != nil {
			multilog.Fatal("util.GetArg", "failed to get arg", map[string]interface{}{
				"error": err,
				"kind":  resultType.Kind(),
			})
		}
		result = any(v).(T)
	case reflect.Slice:
		if resultType.Elem().Kind() == reflect.String {
			v, err := cmd.Flags().GetStringSlice(name)
			if err != nil {
				multilog.Fatal("util.GetArg", "failed to get arg", map[string]interface{}{
					"error": err,
					"kind":  resultType.Kind(),
				})
			}
			result = any(v).(T)
		} else {
			multilog.Fatal("util.GetArg", "unsupported slice element type", map[string]interface{}{
				"type": resultType.Elem().Kind(),
			})
		}
	default:
		multilog.Fatal("util.GetArg", "unsupported type", map[string]interface{}{
			"type": resultType.Kind(),
		})
	}

	return result
}
