package env

import "os"

func GetVar(key string) string {
	return os.Getenv(key)
}

func GetVarWithDefault(key string, def string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	return v
}
