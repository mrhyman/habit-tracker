package config

type DatabaseConfig struct {
	Connection string `yaml:"connection" validate:"required" env:"DB_ENDPOINT"`
	dbName     string `yaml:"db_name" validate:"required"`
	userName   string `yaml:"user_name" validate:"required"`
}

func (dc DatabaseConfig) GetConnection() string {
	return dc.Connection
}

func (dc DatabaseConfig) GetDbName() string {
	return dc.dbName
}

func (dc DatabaseConfig) SetDbName(dbName string) {
	dc.dbName = dbName
}

func (dc DatabaseConfig) GetUserName() string {
	return dc.userName
}

func (dc DatabaseConfig) SetUserName(userName string) {
	dc.userName = userName
}
