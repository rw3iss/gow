package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

//type JsonMap map[string]interface{}

//type Array []interface{}
//type StringType string
//type NumType float64
//type BoolType bool

// Config structure of file and data
type Config struct {
	path string
	db   map[string]interface{}
}

// NewConfig loads the given config file into a dictionary.
func NewConfig(filename string) (*Config, error) {
	config := &Config{
		path: filename,
		db:   make(map[string]interface{}),
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &config.db)
	if err != nil {
		return nil, err
	}

	// Convert to all "real" types on load
	config.db = _parseChildProperties(config.db)

	return config, nil
}

func _parseChildProperties(config map[string]interface{}) map[string]interface{} {
	var values = make(map[string]interface{})

	for k, v := range config {
		switch v.(type) {
		case string:
			values[k] = v.(string)
		case bool:
			values[k] = v.(bool)
		case int:
			values[k] = v.(int)
		case float64:
			values[k] = v.(float64)
		case []interface{}: // array of something
			values[k] = v.([]interface{})
		case map[string]interface{}: // sub-properties
			values[k] = _parseChildProperties(v.(map[string]interface{}))
		default:
			Log("Unknown config value type: %T for key: %s with value: %s. Trying to parse anyway...", v, k, v)
			values[k] = v.(map[string]interface{})
		}
	}

	return values
}

// GetVal returns the underlying anonymous interface{} value.
// Supports sub-properties with a period delimeter, ie. SomeConfig.SomeChild.Key
func (c *Config) GetVal(key string, _default ...interface{}) interface{} {
	var _v interface{}

	// store default value if given
	if len(_default) > 0 {
		_v = _default[0]
	}

	var keys = strings.Split(key, ".")
	var v = c.db[keys[0]]

	if len(keys) > 1 {
		for i := 1; i < len(keys); i++ {
			switch v.(type) {
			case map[string]interface{}:
				v = v.(map[string]interface{})[keys[i]]
			default:
				Log("Child of nested config key '%s' was not a string map", key)
			}
		}
	}

	// If no existing config key exists, try to use the default value
	if v == nil {
		v = _v
	}

	switch v.(type) {
	case int:
		return v.(int)
	case float64:
		return v.(float64)
	case string:
		return v.(string)
	case bool:
		return v.(bool)
	case []interface{}:
		return v.([]interface{})
	case map[string]interface{}:
		return v.(map[string]interface{})
	case nil:
		panic("Config value '" + key + "' does not exist, and no default was provided.")
	default:
		Log("Unknown type parsing config key '%s'. Value was: %v, of type: %T", key, v, v)
	}

	return v
}

// Get returns a string value for given config key.
func (c *Config) Get(key string, _default ...interface{}) string {
	var v interface{} = c.GetVal(key, _default...)
	// return formatted string
	return fmt.Sprintf("%v", v)
}

// GetString returns a string value for given config key.
func (c *Config) GetString(key string, _default ...interface{}) string {
	return c.Get(key, _default...)
}

// GetArray returns an array value for given config key.
func (c *Config) GetArray(key string, _default ...interface{}) []interface{} {
	var v interface{} = c.GetVal(key, _default...)
	switch v.(type) {
	case ([]interface{}):
		v = v.([]interface{})
		return v.([]interface{})
	default:
		panic(fmt.Sprintf("Unknown type for GetArray, key: %s, expected array, given: %T\n", key, v))
	}
}

// GetStringArray returns an array of strings for a given config key.
func (c *Config) GetStringArray(key string, _default ...interface{}) []string {
	return ArrayInterfaceToString(c.GetArray(key, _default...))
}

// GetInt returns a number value for given config key.
func (c *Config) GetInt(key string, _default ...interface{}) int {
	var v interface{} = c.GetVal(key, _default...)
	switch v.(type) {
	case int:
		return v.(int)
	case float64:
		return int(v.(float64))
	case string:
		if v, err := strconv.Atoi(v.(string)); err != nil {
			panic(fmt.Sprintf("Unknown type for int, key: %s, expected int, given: %T\n", key, v))
		} else {
			return v
		}
	}
	return v.(int)
}

// GetBool returns a bool value for given config key.
func (c *Config) GetBool(key string, _default ...interface{}) bool {
	var v interface{} = c.GetVal(key, _default...)
	switch v.(type) {
	case string:
		if b, err := strconv.ParseBool(v.(string)); err != nil {
			panic(fmt.Sprintf("Unknown type for GetBool, key: %s, expected bool, given: %T\n", key, v))
		} else {
			return b
		}
	case bool:
		return v.(bool)
	default:
		panic(fmt.Sprintf("Unknown type for GetBool, key: %s, expected bool, given: %T\n", key, v))
	}

}
