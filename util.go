package env_config

import "os"

func Env(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func PreferEnv(fallback string, vars ...string) string {
	for _, env := range vars {
		res := os.Getenv(env)
		if res != "" {
			return res
		}
	}
	return fallback
}
