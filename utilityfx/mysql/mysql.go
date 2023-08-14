package mysql

import (
	"context"
	"fmt"
	"github.com/dlmiddlecote/sqlstats"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/viper"
	"github.com/uptrace/opentelemetry-go-extra/otelsqlx"
	"go.uber.org/fx"
	"log"
	"time"
)

var Module = fx.Provide(newMySQL)

func newMySQL(lifecycle fx.Lifecycle) (*sqlx.DB, error) {
	mysql := "mysql"
	username := viper.GetString(fmt.Sprintf("%s.username", mysql))
	password := viper.GetString(fmt.Sprintf("%s.password", mysql))
	url := viper.GetString(fmt.Sprintf("%s.url", mysql))
	schema := viper.GetString(fmt.Sprintf("%s.schema", mysql))

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=5s",
		username, password, url, schema)

	log.Printf("Connecting to database: %s-%s!", url, schema)
	db, err := otelsqlx.Connect("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	return initTuningConfig(lifecycle, mysql, db, url, schema)
}

func initTuningConfig(lifecycle fx.Lifecycle, mysql string, db *sqlx.DB, url string, schema string) (*sqlx.DB, error) {
	maxIdle := viper.GetInt(fmt.Sprintf("%s.max-idle-connections", mysql))
	maxOpen := viper.GetInt(fmt.Sprintf("%s.max-open-connections", mysql))
	maxIdleTime := viper.GetInt(fmt.Sprintf("%s.max-idle-conn-time-in-millis", mysql))
	maxLifeTime := viper.GetInt(fmt.Sprintf("%s.max-life-conn-time-in-millis", mysql))

	db.SetMaxIdleConns(maxIdle)
	db.SetMaxOpenConns(maxOpen)
	db.SetConnMaxIdleTime(time.Duration(maxIdleTime) * time.Millisecond)
	db.SetConnMaxLifetime(time.Duration(maxLifeTime) * time.Millisecond)

	collector := sqlstats.NewStatsCollector(mysql, db)
	prometheus.MustRegister(collector)

	log.Printf("Connecting to database: %s-%s successfully!", url, schema)

	lifecycle.Append(fx.Hook{OnStop: func(ctx context.Context) error {
		log.Printf("Closing DB: %s-%s!", url, schema)
		return db.Close()
	}})

	return db, nil
}
