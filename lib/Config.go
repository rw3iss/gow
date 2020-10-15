package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

type JsonMap map[string]interface{}

//type StringType string
//type NumType float64
//type BoolType bool

// Config structure of file and data
type Config struct {
	path string
	db   JsonMap
}

// config global instance
var config Config

// NewConfig loads the given config file into a dictionary.
func NewConfig(filename string) (*Config, error) {
	config = Config{
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
			//fmt.Printf("Item %q is a string, containing %q\n", k, c)
			config.db[k] = v.(string)
		case bool:
			//fmt.Printf("Item %q is a bool, containing %t\n", k, c)
			config.db[k] = v.(bool)
		case int:
			//fmt.Printf("Looks like item %q is a number, specifically %f\n", k, c)
			config.db[k] = v.(int)
		case float64:
			//fmt.Printf("Looks like item %q is a number, specifically %f\n", k, c)
			config.db[k] = v.(float64)
		default:
			config.db[k] = JsonMap(v.(map[string]interface{}))
		}
	}

	return &config, nil
}

// GetVal returns the underlying anonymous interface{} value.
// TODO: Parse the incoming key for periods, and recurse into the subvalues.
func (c *Config) GetVal(key string, def ...interface{}) interface{} {
	var _v interface{}

	if len(def) > 0 {
		//Log("set default: %+v, len:%d, type:%T\n", def[0], len(def), def)
		//Log("type: %T\n", def[0])
		_v = def[0]
	} else {
		//Log("No default.\n")
	}

	// If no existing config key exists, try to use the default value
	var v = config.db[key]
	if v == nil {
		v = _v
	}

	switch v.(type) {
	// case NumType:
	// 	v = int(v.(NumType))
	// case StringType:
	// 	v = string(v.(StringType))
	// case BoolType:
	// 	v = bool(v.(BoolType))
	case int:
		v = v.(int)
	case float64:
		v = v.(float64)
	case string:
		v = v.(string)
	case bool:
		v = v.(bool)
	case nil:
		panic("Config value '" + key + "' does not exist, and no default was provided.")
	default:
		Log("Unknown type parsing config value '%s'. Value: %v, Type: %T", key, v, v)
	}

	return v
}

// Get returns a string value for given config key.
func (c *Config) Get(key string, def ...interface{}) string {
	var v interface{} = config.GetVal(key, def...)
	// return formatted string
	return fmt.Sprintf("%v", v)
}

// GetInt returns a number value for given config key.
func (c *Config) GetInt(key string, def ...interface{}) int {
	var v interface{} = config.GetVal(key, def...)
	switch v.(type) {
	case int:
		v = v.(int)
	case float64:
		v = int(v.(float64))
	case string:
		v, _ = strconv.Atoi(v.(string))
	}
	return v.(int)
}

// GetBool returns a bool value for given config key.
func (c *Config) GetBool(key string, def ...interface{}) bool {
	var v interface{} = config.GetVal(key, def...)
	var b bool

	switch c := v.(type) {
	case string:
		b, _ = strconv.ParseBool(v.(string))
	case bool:
		b = v.(bool)
	default:
		panic(fmt.Sprintf("Unknown type for '%s' bool given: %s\n", key, c))
	}

	return b
}
