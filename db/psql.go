package database

import (
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/ruangnyaman/rna-ecommerce-backend/conf"

	// PostgreSQL DB Driver
	_ "github.com/lib/pq"
)

//PSQLDBInitWrite : init DB
func PSQLDBInit() *sqlx.DB {
	Logger.Info().Msg("PQWrite - DB open started")
	v, err := conf.GetConfig()
	dbname := v.GetString("datastore.psql.db")
	host := v.GetString("datastore.psql.host")
	user := v.GetString("datastore.psql.user")
	password := v.GetString("datastore.psql.password")
	port := v.GetString("datastore.psql.port")
	schema := v.GetString("datastore.psql.schema")
	maxcon := v.GetInt("datastore.psql.maxcon")
	maxidle := v.GetInt("datastore.psql.maxidle")

	//db, err := OpenSqlxViaPgxConnPool(dbname, host, user, password, port, maxcon)

	db, err := sqlx.Open("postgres",
		"host="+host+" port="+port+" user="+user+" password="+password+" dbname="+dbname+" sslmode=require search_path="+schema)

	if err != nil {
		Logger.Fatal().Msgf("PQ - DB open failed %s", err.Error())
		return nil
	}

	err = db.Ping()
	if err != nil {
		Logger.Fatal().Msgf("PQ - DB ping failed %s", err.Error())
		return nil
	}

	db.SetMaxOpenConns(maxcon)
	db.SetMaxIdleConns(maxidle)
	db.SetConnMaxLifetime(5 * time.Minute)
	Logger.Info().Msg("PQ - DB open completed")
	return db
}
