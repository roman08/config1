package env_config

import (
	"os"

	"github.com/inconshreveable/log15"
)

type Validator struct {
	logger log15.Logger
}

func NewValidator(logger log15.Logger) *Validator {
	return &Validator{logger}
}

func (*Validator) CriticalIfEmpty(value string, name string, logger log15.Logger) {
	if len(value) == 0 {
		logger.Error("config not found. Check env variables", "envKey", name)
		os.Exit(1)
	}
}

func (*Validator) WarningIfEmpty(value string, name string, logger log15.Logger) {
	if len(value) == 0 {
		logger.Warn("config not found. Check env variables", "envKey", name)
	}
}

func (*Validator) ValidateCors(corsConfig *Cors, logger log15.Logger) {
	if len(corsConfig.Headers) < 1 {
		logger.Warn("corsConfig not found. Check env variables", "envKey", "Cors.Headers")
	}
	if len(corsConfig.Methods) < 1 {
		logger.Warn("corsConfig not found. Check env variables", "envKey", "Cors.Methods")
	}
	if len(corsConfig.Origins) < 1 {
		logger.Warn("corsConfig not found. Check env variables", "envKey", "Cors.Origins")
	}
}

func (v *Validator) ValidateDb(config *Db, logger log15.Logger) {
	v.CriticalIfEmpty(config.Host, "DB.Host", logger)
	v.CriticalIfEmpty(config.User, "DB.User", logger)
	v.CriticalIfEmpty(config.Password, "DB.Password", logger)
}
