package main

import (
	"os"
	"testing"
)

var envHost string = os.Getenv("AUTHTABLES_HOST")
var envPort string = os.Getenv("AUTHTABLES_PORT")
var envPassword string = os.Getenv("AUTHTABLES_PW")
var envLoglevel string = os.Getenv("AUTHTABLES_LOGLEVEL")
var envBloomSize string = os.Getenv("AUTHTABLES_BLOOMSIZE")
var envShard string = os.Getenv("AUTHTABLES_SHARD")

func TestConfigurationDefaults(t *testing.T) {

	os.Unsetenv("AUTHTABLES_HOST")
	os.Unsetenv("AUTHTABLES_PORT")
	os.Unsetenv("AUTHTABLES_PW")
	os.Unsetenv("AUTHTABLES_LOGLEVEL")
	os.Unsetenv("AUTHTABLES_BLOOMSIZE")
	os.Unsetenv("AUTHTABLES_SHARD")

	config := getConfig()

	if config.Host != defaultHost {
		t.Errorf("Expected default Host value '%s' - Got '%s'", defaultHost, config.Host)
	}
	if config.Port != defaultPort {
		t.Errorf("Expected default Port value '%s' - Got '%s'", defaultPort, config.Port)
	}
	if config.Password != defaultPassword {
		t.Errorf("Expected default Password value '%s' - Got '%s'", defaultPassword, config.Password)
	}
	if config.LogLevel != defaultLogLevel {
		t.Errorf("Expected default LogLevel value '%s' - Got '%s'", defaultLogLevel, config.LogLevel)
	}
	if config.BloomSize != defaultBloomSize {
		t.Errorf("Expected default BloomSize value '%d' - Got '%d'", defaultBloomSize, config.BloomSize)
	}
	if config.Shard != defaultShard {
		t.Errorf("Expected default Shard value '%s' - Got '%s'", defaultShard, config.Shard)
	}

	restoreEnv()
}

func restoreEnv() {
	os.Setenv("AUTHTABLES_HOST", envHost)
	os.Setenv("AUTHTABLES_PORT", envPort)
	os.Setenv("AUTHTABLES_PW", envPassword)
	os.Setenv("AUTHTABLES_LOGLEVEL", envLoglevel)
	os.Setenv("AUTHTABLES_BLOOMSIZE", envBloomSize)
	os.Setenv("AUTHTABLES_SHARD", envShard)
}

func TestConfigurationEnvVars(t *testing.T) {

	os.Setenv("AUTHTABLES_HOST", "test")
	os.Setenv("AUTHTABLES_PORT", "test")
	os.Setenv("AUTHTABLES_PW", "test")
	os.Setenv("AUTHTABLES_LOGLEVEL", "test")
	os.Setenv("AUTHTABLES_BLOOMSIZE", "9999")
	os.Setenv("AUTHTABLES_SHARD", "test")

	config := getConfig()

	if config.Host != "test" {
		t.Errorf("Expected default Host value '%s' - Got '%s'", "test", config.Host)
	}
	if config.Port != "test" {
		t.Errorf("Expected default Port value '%s' - Got '%s'", "test", config.Port)
	}
	if config.Password != "test" {
		t.Errorf("Expected default Password value '%s' - Got '%s'", "test", config.Password)
	}
	if config.LogLevel != "test" {
		t.Errorf("Expected default LogLevel value '%s' - Got '%s'", "test", config.LogLevel)
	}
	if config.BloomSize != 9999 {
		t.Errorf("Expected default BloomSize value '%d' - Got '%d'", 9999, config.BloomSize)
	}
	if config.Shard != "test" {
		t.Errorf("Expected default Shard value '%s' - Got '%s'", "test", config.Shard)
	}

	restoreEnv()
}
