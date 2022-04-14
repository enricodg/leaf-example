package pkgConfig

import (
	leafConfig "github.com/paulusrobin/leaf-utilities/config"
	"time"
)

type (
	MySQLConfig struct {
		// - Toggle
		MySqlEnable bool `envconfig:"MYSQL_ENABLE" default:"false"`

		// - MySQL
		MySqlAddress               []string      `envconfig:"MYSQL_ADDRESS" required:"true"`
		MySqlUsername              string        `envconfig:"MYSQL_USERNAME" required:"true"`
		MySqlPassword              string        `envconfig:"MYSQL_PASSWORD"`
		MySqlDbName                string        `envconfig:"MYSQL_DB_NAME" required:"true"`
		MySqlMaxIdleConnection     int           `envconfig:"MYSQL_MAX_IDLE_CONNECTION" default:"15"`
		MySqlMaxOpenConnection     int           `envconfig:"MYSQL_MAX_OPEN_CONNECTION" default:"15"`
		MySqlMaxLifetimeConnection time.Duration `envconfig:"MYSQL_MAX_LIFETIME_CONNECTION" default:"60s"`
		MySqlLogMode               bool          `envconfig:"MYSQL_LOG_MODE" default:"false"`
	}
)

func NewMySQLConfig() (MySQLConfig, error) {
	configuration := MySQLConfig{}
	if err := leafConfig.NewFromEnv(&configuration); err != nil {
		return MySQLConfig{}, err
	}
	return configuration, nil
}
