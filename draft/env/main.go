package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type config struct {
	Environment string
	Port        int
	DisableLogs bool
}

func main() {
	// load .env into environment variables, it does trim strings
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	deployEnvironment, err := parseEnv("DEPLOY_PORT", "local", true)
	if err != nil {
		log.Fatalf("Error in configuration: %v", err)
	}
	// fmt.Printf("%T %v\n", deployEnvironment, deployEnvironment)

	deployPort, err := parseEnv("DEPLOY_PORT", 8080, false)
	if err != nil {
		log.Fatalf("Error in configuration: %v", err)
	}

	disableLog := mustParseEnv("DISABLE_LOG", false, false)

	cfg := config{
		Environment: deployEnvironment,
		Port:        deployPort,
		DisableLogs: disableLog,
	}
	fmt.Printf("%#v", cfg)
}

type SupportedEnvTypes interface {
	string | int | bool
}

// parseEnv is very opinionated about required and empty behaviour
// The type is defined by defaultValue
func parseEnv[T SupportedEnvTypes](name string, defaultValue T, isRequired bool) (T, error) {
	var ret T

	val, present := os.LookupEnv(name)
	if !present && isRequired {
		return ret, fmt.Errorf("required environment variable %s is not set", name)
	}
	if !present {
		return defaultValue, nil
	}
	val = strings.TrimSpace(val)
	if val == "" {
		return defaultValue, nil
	}

	switch any(defaultValue).(type) {
	case int:
		i, err := strconv.Atoi(val)
		if err != nil {
			return ret, fmt.Errorf("parse error for env var\"%s\" is not an int", name)
		}
		ret = reflect.ValueOf(i).Interface().(T)
	case string:
		ret = reflect.ValueOf(val).Interface().(T)
	case bool:
		if val == "True" || val == "true" || val == "1" {
			ret = reflect.ValueOf(true).Interface().(T)
		} else if val == "False" || val == "false" || val == "0" {
			ret = reflect.ValueOf(false).Interface().(T)
		} else {
			return ret, fmt.Errorf("parse error for env var\"%s\" is not a bool", name)
		}
	default:
		// TODO: check dynamically basedo n type
		return ret, fmt.Errorf("type %T is not supported", defaultValue)
	}

	return ret, nil
}

func mustParseEnv[T SupportedEnvTypes](name string, defaultValue T, isRequired bool) T {
	ret, err := parseEnv(name, defaultValue, isRequired)
	if err != nil {
		log.Fatalf("Error in configuration: %v", err)
	}
	return ret
}
