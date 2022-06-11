package config

type Config struct {
	DB     *DatabaseConnections
	Fields *Fields
}

func GetConfig() *Config {

	config := &Config{}

	populateFields(config)
	populateDBConnections(config)

	return config
}
