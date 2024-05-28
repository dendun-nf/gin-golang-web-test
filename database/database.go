package database

import (
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"fmt"
	"github.com/dendun-nf/gin-golang-web-test/internal/todo"
	goMysql "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"os"
)

var GormDb *gorm.DB

func init() {
	_ = godotenv.Load(".env.development")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbSSLCertPath := os.Getenv("DB_SSL_CERT_PATH")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=custom&charset=utf8mb4&parseTime=True",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName)

	rootCertPool := x509.NewCertPool()
	caPem, err := os.ReadFile(dbSSLCertPath)
	if err != nil {
		log.Fatalf("Error reading CA file: %v", err)
	}

	if ok := rootCertPool.AppendCertsFromPEM(caPem); !ok {
		log.Fatalf("Error appending CA certs")
	}

	err = goMysql.RegisterTLSConfig("custom", &tls.Config{
		RootCAs: rootCertPool,
	})
	if err != nil {
		log.Fatalf("Error registering custom TLS config: %v", err)
	}

	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	//	NamingStrategy: schema.NamingStrategy{
	//		SingularTable: true,
	//	},
	//})

	sqlDb, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: sqlDb}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   "t_",
		},
	})
	if err != nil {
		log.Fatalf("Error attached to GORM: %v", err)
	}

	err = gormDB.AutoMigrate(&todo.Todo{})
	if err != nil {
		log.Fatalf("Error Migrating Data: %v", err)
	}

	GormDb = gormDB
	log.Println("Connected to database successfully")
}
