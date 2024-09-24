package psql_test

import (
	"os"
	"testing"
)

func Test_ConnStr(t *testing.T) {
	t.Setenv("DB_HOST", "localhost")
	t.Setenv("DB_PORT", "5432")
	t.Setenv("DB_USERNAME", "root")
	t.Setenv("DB_NAME", "defaultdb")
	t.Setenv("DB_PASSWORD", "passw")

	connConf := os.ExpandEnv("host=${DB_HOST} user=${DB_USERNAME} password=${DB_PASSWORD} dbname=${DB_NAME} port=${DB_PORT} sslmode=disable")

	t.Logf("%s", connConf)
}
