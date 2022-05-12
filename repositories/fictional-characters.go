package repositories

import (
	"database/sql"
	"fantasy-info-center-api-gin/models"
	"flag"
	"fmt"
	"log"

	_ "github.com/microsoft/go-mssqldb"
)

var (
	debug         = flag.Bool("debug", false, "enable debugging")
	password      = flag.String("password", "Cx123456", "the database password")
	port     *int = flag.Int("port", 1433, "the database port")
	server        = flag.String("server", "localhost", "the database server")
	user          = flag.String("user", "sa", "the database user")
)

var db *sql.DB

func init() {
	flag.Parse()

	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
	}

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=FantasyInfoCenter", *server, *user, *password, *port)
	_db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal(err)
	}
	db = _db
}

func GetFictionalCharacterByID(id string) models.FictionalCharacter {
	selectSql := "select id, name from FictionalCharacters where id = " + id
	row := db.QueryRow(selectSql)
	var fictionalCharacterId int64
	var fictionalCharacterName string
	row.Scan(&fictionalCharacterId, &fictionalCharacterName)
	return models.FictionalCharacter{ID: int(fictionalCharacterId), Name: fictionalCharacterName}
}

func GetFictionalCharactersFromDatabase() []models.FictionalCharacter {
	selectSql := "select id, name from FictionalCharacters"
	allFictionalCharacters := make([]models.FictionalCharacter, 0)
	rows, _ := db.Query(selectSql)
	for rows.Next() {
		var fictionalCharacterId int64
		var fictionalCharacterName string
		var f models.FictionalCharacter
		rows.Scan(&fictionalCharacterId, &fictionalCharacterName)
		f.ID = int(fictionalCharacterId)
		f.Name = fictionalCharacterName
		allFictionalCharacters = append(allFictionalCharacters, f)
	}

	return allFictionalCharacters
}

func InsertFictionalCharacterInDatabaseDatabase(name string) sql.Result {
	insertSql := "insert into FictionalCharacters (name) values (@p1)"
	stmt, _ := db.Prepare(insertSql)
	result, err := stmt.Exec(name)
	log.Println(err)
	return result
}
