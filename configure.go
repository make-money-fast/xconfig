package xconfig

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/ghodss/yaml"
)

func ParseFromFile(file string, out interface{}) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	return parseData(data, out)
}

func ParseFromData(data []byte, out interface{}) error {
	return parseData(data, out)
}

func parseData(b []byte, out interface{}) error {
	if err := yaml.Unmarshal(b, out); err != nil {
		return err
	}
	val := reflect.Indirect(reflect.ValueOf(out))
	typ := val.Type()

	return assignDefaultValueToField(val, typ)
}

func assignDefaultValueToField(val reflect.Value, typ reflect.Type) error {
	for i := 0; i < typ.NumField(); i++ {
		field := val.Field(i)
		typ := typ.Field(i)
		ignore := typ.Tag.Get("json") == "-"
		if ignore {
			continue
		}
		fieldValue := reflect.Indirect(field)
		if field.Kind() == reflect.Ptr && field.IsNil() {
			fieldValue = reflect.New(field.Type().Elem())
			field.Set(fieldValue)
			fieldValue = fieldValue.Elem()
		}
		if fieldValue.Kind() == reflect.Struct {
			if err := assignDefaultValueToField(fieldValue, fieldValue.Type()); err != nil {
				return err
			}
			continue
		}
		if !fieldValue.CanSet() {
			continue
		}
		if !fieldValue.IsZero() {
			continue
		}
		value := strings.TrimSpace(typ.Tag.Get("default"))
		if value == "" {
			continue
		}
		envKey := typ.Tag.Get("env")
		if envKey != "" {
			v := os.Getenv(envKey)
			if v != "" {
				value = v
			}
		}

		switch fieldValue.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return fmt.Errorf("%s: can not parse: %v to an integer number", fieldValue.Type().Name(), value)
			}
			fieldValue.SetInt(v)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return fmt.Errorf("%s: can not parse: %v to an integer number", fieldValue.Type().Name(), value)
			}
			if v < 0 {
				return fmt.Errorf("%s: can not set a negative value: %s to uint value", fieldValue.Type().Name(), value)
			}
			fieldValue.SetUint(uint64(v))
		case reflect.Float64, reflect.Float32:
			v, err := strconv.ParseFloat(value, 10)
			if err != nil {
				return fmt.Errorf("%s: can not parse: %v to an float number", fieldValue.Type().Name(), value)
			}
			fieldValue.SetFloat(v)
		case reflect.String:
			fieldValue.SetString(value)
		case reflect.Slice:
			// 将 JSON 数据解码到空切片中
			if err := json.Unmarshal([]byte(value), fieldValue.Addr().Interface()); err != nil {
				return err
			}
		case reflect.Bool:
			ok, err := strconv.ParseBool(value)
			if err != nil {
				return err
			}
			fieldValue.SetBool(ok)
		default:
		}
	}

	return nil
}
