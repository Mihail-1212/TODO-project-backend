package main

import (
	"flag"
	"fmt"

	"github.com/gchaincl/dotsql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // for sqlx.Connect postgres
	"github.com/mihail-1212/todo-project-backend/internal/config"
	"github.com/mihail-1212/todo-project-backend/pkg/logger"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	includeTables := []string{
		"task",
		"user",
	}

	log, _ := logger.NewLoggerDev()

	if err := config.InitConfig(); err != nil {
		log.Panic("error initializing configs",
			zap.String("package", "migrate.main"),
			zap.String("function", "initConfig"),
			zap.Error(err))
		return
	}

	// Parse flags
	SQLMode := flag.String("mode", "", "mode for migrate. up - create table. down - drop table.")
	flag.Parse()
	log.Info("Migration mode", zap.String("mode", *SQLMode))

	// Database connection
	con := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s host=%s",
		viper.GetString("db.username"),
		viper.GetString("db.dbname"),
		viper.GetString("db.password"),
		viper.GetString("db.sslmode"),
		viper.GetString("db.host"))

	db, err := sqlx.Connect("postgres", con)
	if err != nil {
		log.Fatal("Error connection to database",
			zap.String("package", "main"),
			zap.String("function", "sqlx.Connect"),
			zap.Error(err))
	}

	// Upload SQL
	var tables []*dotsql.DotSql

	for _, table := range includeTables {
		dotSqlTable, err := dotsql.LoadFromFile(fmt.Sprintf("internal/schema/%s.sql", table))

		if err != nil {
			log.Fatal("sql upload error",
				zap.String("package", "main"),
				zap.String("function", "dotsql.LoadFromFile"),
				zap.Error(err))
		}

		_ = append(tables, dotSqlTable)
	}

	if *SQLMode == "up" {

		for _, table := range tables {
			_, err = table.Exec(db, "migrate")
			if err != nil {
				log.Fatal("error creating tables",
					zap.String("package", "main"),
					zap.String("function", fmt.Sprintf("%d.Exec", table)),
					zap.Error(err))
			}
		}

		log.Info("UP DATABASE")
	}

	if *SQLMode == "drop" {

		for _, table := range tables {
			_, err = table.Exec(db, "drop")
			if err != nil {
				log.Fatal("error creating tables",
					zap.String("package", "main"),
					zap.String("function", fmt.Sprintf("%s.Exec", table)),
					zap.Error(err))
			}
		}

	}
	if *SQLMode == "data" {
		// _, err = gender.Exec(db, "data")
		// if err != nil {
		// 	log.Fatal("error creating tables",
		// 		zap.String("package", "main"),
		// 		zap.String("function", "gender.Exec"),
		// 		zap.Error(err))
		// }
	}
}
