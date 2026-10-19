package env

import (
	"os"
	"strconv"
	"time"
)

func GetString(name, defaults string) string{
	value, ok := os.LookupEnv(name)
	if !ok{
		return defaults
	}

	return value
}

func GetInt(name, defaults string) int{
	value, ok := os.LookupEnv(name)
	if !ok{
		return toInt(defaults)
	}

	return toInt(value)
}

func toInt(v string) int{
	val, _:= strconv.Atoi(v)
	return val
}

func GetTime(name, defaults string) time.Duration{
	value, ok := os.LookupEnv(name)
	if !ok{
		value = defaults
	}
	timeout, err := toTime(value)
	if err!=nil{
		return time.Second * 5
	}
	return timeout
}

func toTime(value string) (time.Duration, error){
	timeout, err := time.ParseDuration(value)
	if err!=nil{
		return 0, err
	}
	return timeout, nil
}