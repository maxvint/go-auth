package config

import "code.bigogo.com/gopkg/dbutil"

type AppConfig struct {
	DBAddress dbutil.MySQLAuth `yaml: "db_auth_address"`
}

func (a *AppConfig) Verify() error {
	return nil
}
