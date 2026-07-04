package db

import (
	"errors"
	"fmt"
	"log"
	"time"

	"room-service-msc/infrastructure/message"

	"room-service-msc/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db       *gorm.DB
	cfg      = config.Get // your config loader
	interval = cfg("db.interval").(int)
)

// auto start pool monitor
func init() {
	go connectionPool()
}

// DBInstance returns singleton instance
func DBInstance() (*gorm.DB, error) {
	if db == nil {
		if err := connect(); err != nil {
			return nil, err
		}
	}
	return db, nil
}

func BeginTransaction() (*gorm.DB, error) {

	db, err := DBInstance()
	if err != nil {
		//logging.Failed(logging.L().Log(), errors.New(message.ErrInitializeDB))
		return nil, errors.New(message.ErrInitializeDB)
	}

	tx := db.Begin()
	if tx.Error != nil {
		//logging.Failed(errors.New(message.ErrBeginTrx))
		return nil, errors.New(message.ErrBeginTrx)
	}

	return tx, nil
}

func FinishTransaction(tx *gorm.DB, err *error) {

	if *err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
}

// ──────────────────────────────────────────────────────────
//
//	CONNECT TO MYSQL (LOCAL OR PRODUCTION)
//
// ──────────────────────────────────────────────────────────
func connect() error {

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg("db.user"),
		cfg("db.password"),
		cfg("db.host"),
		cfg("db.port"),
		cfg("db.name"),
	)

	gdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		log.Println("❌ Failed to connect to MySQL:", err)
		return err
	}

	if err := setupPool(gdb); err != nil {
		return err
	}

	db = gdb
	log.Println("✅ MySQL connected successfully")

	return nil
}

// ──────────────────────────────────────────
//
//	DB POOL CONFIGURATION
//
// ──────────────────────────────────────────
func setupPool(gdb *gorm.DB) error {
	sqlDB, err := gdb.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(cfg("db.idle-conn").(int))
	sqlDB.SetMaxOpenConns(cfg("db.open-conn").(int))
	sqlDB.SetConnMaxLifetime(time.Minute * 5)

	return nil
}

// ─────────────────────────────────────────────────────
//
//	CONNECTION POOL MONITOR (AUTO RECONNECT)
//
// ─────────────────────────────────────────────────────
func connectionPool() {
	for {
		time.Sleep(time.Duration(interval) * time.Second)

		instance, err := DBInstance()
		if err != nil {
			log.Println("❌ MySQL reconnect failed:", err)
			continue
		}

		sqlDB, err := instance.DB()
		if err != nil {
			log.Println("❌ MySQL DB instance error:", err)
			continue
		}

		if err := sqlDB.Ping(); err != nil {
			log.Println("⚠️ MySQL connection lost, reconnecting...")
			_ = connect()
		}
	}
}
