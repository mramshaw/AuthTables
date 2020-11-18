package main

import (
	"os"
	"strconv"
)

const (
	defaultHost      = "redis"
	defaultPort      = "6379"
	defaultPassword  = ""
	defaultLogLevel  = "debug"
	defaultBloomSize = 1000000
	defaultShard     = ""
)

//Configuration holds data from ENV variables, or defaults
type Configuration struct {
	Host      string
	Port      string
	Password  string
	LogLevel  string
	BloomSize uint
	Shard     string
}

//Global access to configuration variables
var c = getConfig()

func getConfig() (c Configuration) {

	configuration := Configuration{}

	// set to defaults
	configuration.Host = defaultHost
	configuration.Port = defaultPort
	configuration.Password = defaultPassword
	configuration.LogLevel = defaultLogLevel
	configuration.BloomSize = defaultBloomSize
	configuration.Shard = defaultShard

	// override with ENV values, if set
	if envVal, envSet := os.LookupEnv("AUTHTABLES_HOST"); envSet {
		configuration.Host = envVal
	}
	if envVal, envSet := os.LookupEnv("AUTHTABLES_PORT"); envSet {
		configuration.Port = envVal
	}
	if envVal, envSet := os.LookupEnv("AUTHTABLES_PW"); envSet {
		configuration.Password = envVal
	}
	if envVal, envSet := os.LookupEnv("AUTHTABLES_LOGLEVEL"); envSet {
		configuration.LogLevel = envVal
	}
	if envVal, envSet := os.LookupEnv("AUTHTABLES_BLOOMSIZE"); envSet {
		bs, _ := strconv.ParseUint(envVal, 0, 32)
		configuration.BloomSize = uint(bs)
	}
	if envVal, envSet := os.LookupEnv("AUTHTABLES_SHARD"); envSet {
		configuration.Shard = envVal
	}

	return configuration
}
