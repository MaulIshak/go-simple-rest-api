package connection

import (
	"database/sql"
	"fmt"
	"go-simple-rest-api/internal/config"
	"log"

	_ "github.com/lib/pq"
)

func GetDatabase(config config.Database) *sql.DB{
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable Timezone=%s",
	config.Host,
	config.Port,
	config.User,
	config.Pass,
	config.Name,
	config.Tz,	
	)
	db, err := sql.Open("postgres", dsn)

	if err != nil{
		log.Fatal("failed to open connection: ", err.Error())
	}
	err = db.Ping()
	if err != nil{
		log.Fatal("failed to ping connection: ", err.Error())
	}

	return db
}