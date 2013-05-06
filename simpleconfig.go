package simpleconfig
//a package for handling simple config files in key=value format.
//the code treats # as the comment character.

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	COMMENT   = "#"
	SEPARATOR = "="
)

//struct representing a config file
// which are simple key=value entries
type Config struct {
	Path    string
	Entries map[string]string
}

//Creates a new config struct for the given path
// which is assumed to be a file on the file system
func NewConfig(path string) (c *Config, err error) {
	cfg := new(Config)
	cfg.Path = path
	cfg.Entries = make(map[string]string)
	parseError := cfg.parse()
	return cfg, parseError
}

//Parses the configuration file and stores all key-value
//pairs in the Entries map
func (c Config) parse() error {
	contents, err := ioutil.ReadFile(c.Path)
	if err != nil {
		return err
	}
	for _, line := range strings.Split(string(contents), "\n") {
		line = strings.TrimSpace(line)
		if len(line) > 2 && !strings.HasPrefix(line, COMMENT) {
			row := strings.SplitN(line, SEPARATOR, 2)
			key, value := strings.TrimSpace(string(row[0])),
				strings.TrimSpace(string(row[1]))
			c.Entries[key] = value
		}
	}
	return nil
}

// Returns the config entry belonging to the given key, if present,
// otherwise err will be not nil.
func (c Config) GetString(key string) (entry string, err error) {
	value, ok := c.Entries[key]
	var e error
	if !ok {
		msg := fmt.Sprintf("unknown configuration entry '%s'", key)
		e = errors.New(msg)
	} else {
		e = nil
	}
	return value, e
}

//Returns the value for the given key, if it is present, otherwise
//the val is returned.
func (c Config) GetStringDefault(key string, val string) string {
	entry, e := c.GetString(key)
	if e == nil {
		return entry
	}
	return val
}

//Returns the value stored under the key as an int if present, otherwise
//err is set to non nil
func (c Config) GetInt(key string) (val int, err error) {
	entry, e := c.GetString(key)
	if e == nil {
		return strconv.Atoi(entry)
	}
	msg := fmt.Sprintf("unknown configuration entry '%s'", key)
	return -1, errors.New(msg)
}

// Returns the value stored under key as an int, if no such entry exists
// it will return the given val.
func (c Config) GetIntDefault(key string, val int) int {
	entry, e := c.GetInt(key)
	if e == nil {
		return entry
	}
	return val
}

// Returns the boolean value for the given key. In case the key does not exist
// or the parsing of the boolean value fails err will be set to a non nil value
func (c Config) GetBool(key string) (b bool, err error){
    entry, e := c.GetString(key)
	if e == nil {
		return strconv.ParseBool(entry)
    }
	msg := fmt.Sprintf("unknown configuration entry '%s'", key)
	return false, errors.New(msg)
}


// Returns the boolean value for the given key. In case the key does not exist
// the given default value is returned.
func (c Config) GetBoolDefault(key string, value bool) (b bool){
    entry, e := c.GetBool(key)
    if e == nil{
        return entry
    }
    return value;
}
