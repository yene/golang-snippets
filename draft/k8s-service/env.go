package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type SupportedEnvTypes interface {
	string | int | bool
}

// parseEnv is very opinionated about required and empty behaviour
// The type is defined by defaultValue
// isRequired: the environment must be provided from the outside, useful for secrets
func parseEnv[T SupportedEnvTypes](name string, defaultValue T, isRequired bool) (T, error) {
	var emptyRet T
	var ret T

	val, present := os.LookupEnv(name)
	if !present && isRequired {
		return emptyRet, fmt.Errorf("required environment variable %s is not set", name)
	}
	if !present {
		return defaultValue, nil
	}
	val = strings.TrimSpace(val)
	if val == "" && isRequired {
		return emptyRet, fmt.Errorf("required environment variable %s has no value", name)
	}

	if val == "" {
		return defaultValue, nil
	}

	switch any(defaultValue).(type) {
	case int:
		i, err := strconv.Atoi(val)
		if err != nil {
			return emptyRet, fmt.Errorf("parse error for env var\"%s\" is not an int", name)
		}
		ret = any(i).(T)
	case string:
		ret = any(val).(T)
	case bool:
		if val == "True" || val == "true" || val == "1" {
			ret = any(true).(T)
		} else if val == "False" || val == "false" || val == "0" {
			ret = any(false).(T)
		} else {
			return emptyRet, fmt.Errorf("parse error for env var\"%s\" is not a bool", name)
		}
	default:
		return emptyRet, fmt.Errorf("type %T is not supported", defaultValue)
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
