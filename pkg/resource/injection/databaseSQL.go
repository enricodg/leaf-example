package injection

import (
	pkgConfig "github.com/enricodg/leaf-example/pkg/config"
	leafGormMySql "github.com/paulusrobin/leaf-utilities/database/sql/integrations/gorm/mysql"
	leafSql "github.com/paulusrobin/leaf-utilities/database/sql/sql"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
)

type SQL struct {
	MySQL leafSql.ORM
}

func NewDatabaseSQL(mysqlConfig pkgConfig.MySQLConfig, logger leafLogger.Logger) (SQL, error) {
	var (
		mysql = leafSql.NoopORM()
		err   error
	)

	if mysqlConfig.MySqlEnable {
		mysql, err = newMySQL(mysqlConfig, logger)
		if err != nil {
			return SQL{}, err
		}
	}

	return SQL{
		MySQL: mysql,
	}, nil
}

func newMySQL(mysqlConfig pkgConfig.MySQLConfig, logger leafLogger.Logger) (leafSql.ORM, error) {
	return leafGormMySql.New(
		leafGormMySql.DbConnection{
			Address:  mysqlConfig.MySqlAddress,
			Username: mysqlConfig.MySqlUsername,
			Password: mysqlConfig.MySqlPassword,
			DbName:   mysqlConfig.MySqlDbName,
		},
		leafSql.WithConnMaxLifetime(mysqlConfig.MySqlMaxLifetimeConnection),
		leafSql.WithMaxIdleConnection(mysqlConfig.MySqlMaxIdleConnection),
		leafSql.WithMaxOpenConnection(mysqlConfig.MySqlMaxOpenConnection),
		leafSql.WithLogMode(mysqlConfig.MySqlLogMode),
		leafSql.WithLogger(logger),
	)
}
