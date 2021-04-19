package env_config

import (
	"os"
	"strings"
)

type Reader struct {
	serviceName      string
	serviceNameUpper string
}

func NewReader(serviceName string) *Reader {
	serviceNameUpper := strings.ToUpper(serviceName)
	return &Reader{serviceName, serviceNameUpper}
}

// reads cors configs from ENV variables
func (r *Reader) ReadCorsConfig() *Cors {
	cors := Cors{
		Methods: strings.Split(Env("VELMIE_WALLET_"+r.serviceNameUpper+"_CORS_METHODS", "GET,POST,PUT,PATCH,DELETE,OPTIONS"), ","),
		Origins: strings.Split(Env("VELMIE_WALLET_"+r.serviceNameUpper+"_CORS_ORIGINS", "*"), ","),
		Headers: strings.Split(Env("VELMIE_WALLET_"+r.serviceNameUpper+"_CORS_HEADERS", "*"), ","),
	}

	return &cors
}

// reads DB configs from ENV variables
func (r *Reader) ReadDbConfig() *Db {
	dbConfig := Db{
		Driver:      os.Getenv("VELMIE_WALLET_" + r.serviceNameUpper + "_DB_DRIV"),
		Host:        os.Getenv("VELMIE_WALLET_" + r.serviceNameUpper + "_DB_HOST"),
		Port:        Env("VELMIE_WALLET_"+r.serviceNameUpper+"_DB_PORT", "3306"),
		Schema:      Env("VELMIE_WALLET_"+r.serviceNameUpper+"_DB_NAME", "velmie-wallet-"+r.serviceName),
		User:        os.Getenv("VELMIE_WALLET_" + r.serviceNameUpper + "_DB_USER"),
		Password:    os.Getenv("VELMIE_WALLET_" + r.serviceNameUpper + "_DB_PASS"),
		IsDebugMode: Env("VELMIE_WALLET_"+r.serviceNameUpper+"_DB_IS_DEBUG_MODE", "false") == "true",
	}
	return &dbConfig
}
