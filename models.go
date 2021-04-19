package env_config

type Db struct {
	Driver string
	Host   string
	Port   string
	Schema string

	User     string
	Password string

	IsDebugMode bool
}

type Cors struct {
	Methods []string
	Origins []string
	Headers []string
}
