package env

import (
	"os"
	"strconv"
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