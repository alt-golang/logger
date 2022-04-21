package logger

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"
)

type JSONFormatter struct {
	Formatter
}

func (jsonFormatter JSONFormatter) Format(t time.Time, category string, level string, message string, meta interface{}) string {

	target := map[string]interface{}{
		"level":     level,
		"message":   message,
		"timestamp": t,
		"Category":  category,
	}
	if meta != nil {

		val := reflect.ValueOf(meta)
		fmt.Println("VALUE = ", val)
		fmt.Println("KIND = ", val.Kind())

		if val.Kind() == reflect.Map {
			for _, e := range val.MapKeys() {
				v := val.MapIndex(e)
				target[fmt.Sprint(e)] = v.Interface()
			}
		} else if val.Kind() == reflect.Pointer {
			stype := reflect.ValueOf(meta).Elem()
			for i := 0; i < stype.NumField(); i++ {
				varName := stype.Type().Field(i).Name
				if varName[:1] == strings.ToUpper(varName[:1]) {
					varValue := stype.Field(i).Interface()
					target[varName] = varValue
				}
			}

		} else {
			target["meta"] = meta
		}

	}
	result, _ := json.Marshal(target)

	return string(result)
}
