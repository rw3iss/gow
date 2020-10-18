package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

type JsonMap map[string]interface{}

//type Array []interface{}
//type StringType string
//type NumType float64
//type BoolType bool

// Config structure of file and data
type Config struct {
	path string
	db   JsonMap
}

// NewConfig loads the given config file into a dictionary.
func NewConfig(filename string) (*Config, error) {
	config := &Config{
		path: filename,
		db:   make(JsonMap),
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
	for k, v := range config.db {
		switch v.(type) {
		case string:
			config.db[k] = v.(string)
		case bool:
			config.db[k] = v.(bool)
		case int:
			config.db[k] = v.(int)
		case float64:
			config.db[k] = v.(float64)
		case []interface{}: // array of something
			config.db[k] = v.([]interface{})
		default:
			Log("Unknown config type: %T", v)
			config.db[k] = JsonMap(v.(map[string]interface{}))
		}
	}

	return config, nil
}

// GetVal returns the underlying anonymous interface{} value.
// TODO: Parse the incoming key for periods, and recurse into the subvalues.
func (c *Config) GetVal(key string, def ...interface{}) interface{} {
	var _v interface{}

	// store default value if given
	if len(def) > 0 {
		_v = def[0]
	}

	// If no existing config key exists, try to use the default value
	var v = c.db[key]
	if v == nil {
		v = _v
	}

	switch c := v.(type) {
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
	case nil:
		panic("Config value '" + key + "' does not exist, and no default was provided.")
	default:
		Log("Unknown type parsing config key '%s'. Value was: %v, of type: %s", key, v, c)
	}

	return v
}

// Get returns a string value for given config key.
func (c *Config) Get(key string, def ...interface{}) string {
	var v interface{} = c.GetVal(key, def...)
	// return formatted string
	return fmt.Sprintf("%v", v)
}

// GetArray returns an array value for given config key.
func (c *Config) GetArray(key string, def ...interface{}) []interface{} {
	var v interface{} = c.GetVal(key, def...)
	switch c := v.(type) {
	case ([]interface{}):
		v = v.([]interface{})
		return v.([]interface{})
	default:
		panic(fmt.Sprintf("Unknown type for '%s' GetArray given: %s\n", key, c))
	}
}

// GetInt returns a number value for given config key.
func (c *Config) GetInt(key string, def ...interface{}) int {
	var v interface{} = c.GetVal(key, def...)
	switch v.(type) {
	case int:
		return v.(int)
	case float64:
		return int(v.(float64))
	case string:
		if v, err := strconv.Atoi(v.(string)); err != nil {
			panic(fmt.Sprintf("Unknown type for '%s' int given: %s\n", key, c))
		} else {
			return v
		}
	}
	return v.(int)
}

// GetBool returns a bool value for given config key.
func (c *Config) GetBool(key string, def ...interface{}) bool {
	var v interface{} = c.GetVal(key, def...)
	switch c := v.(type) {
	case string:
		if b, err := strconv.ParseBool(v.(string)); err != nil {
			panic(fmt.Sprintf("Unknown type for '%s' GetBool given: %s\n", key, c))
		} else {
			return b
		}
	case bool:
		return v.(bool)
	default:
		panic(fmt.Sprintf("Unknown type for '%s' GetBool given: %s\n", key, c))
	}

}
