package generate

import (
	"database/sql"
	"math/rand"
	"time"
	"strconv"
	"log"
	"strings"
	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

type GeneratePerson struct {
	DbPerson *sql.DB
	DbImage *sql.DB
}

var generatePerson GeneratePerson


func init() {
	FilesPerson, _ := ReadFileDb("person.db")
	FilesImages, _ := ReadFileDb("images.db")
	generatePerson.DbPerson = GetSql(FilesPerson)
	generatePerson.DbImage = GetSql(FilesImages)
}

func GetSql(files string) *sql.DB {
	db, _ := sql.Open("sqlite3", "file:" + files)
	return db
}



func randomString(length int) string {
	seed := time.Now().UnixNano()
    rng := rand.New(rand.NewSource(seed))
	const charset = "abcdefghijklmnopqrstuvwxyz"
	randomBytes := make([]byte, length)
	_, err := rng.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	for i := 0; i < length; i++ {
		randomBytes[i] = charset[int(randomBytes[i])%len(charset)]
	}
	return string(randomBytes)
}

func randomNumbers(length int) string {
	seed := time.Now().UnixNano()
    rng := rand.New(rand.NewSource(seed))
	const charset = "0123456789"
	randomBytes := make([]byte, length)
	_, err := rng.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	for i := 0; i < length; i++ {
		randomBytes[i] = charset[int(randomBytes[i])%len(charset)]
	}
	return string(randomBytes)
}


func generatePhone() string {
    seed := time.Now().UnixNano()
    rng := rand.New(rand.NewSource(seed))
	numbers := []int{11,12,13,21,22,21,23,52,53,51,14,55,58,15,16,56,57,17,18,19,59,77,78,79,31,32,33,38,81,82,83,84,85,86,87,88,89}
	randomIndex := rng.Intn(len(numbers))
	randomNumber := numbers[randomIndex]
	numberstr := strconv.Itoa(randomNumber)
	return "+628" + numberstr + randomNumbers(8)
}

func randomBirthdate() string {
    minYear := 1980
    maxYear := 2004
    seed := time.Now().UnixNano()
    rng := rand.New(rand.NewSource(seed))
    year := rng.Intn(maxYear-minYear+1) + minYear
    month := rng.Intn(12) + 1
    daysInMonth := time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC).Day()
    day := rng.Intn(daysInMonth) + 1
    timed := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return timed.Format("02-01-2006")
}



func randomUserAll() map[string]interface{} {
	log.Println("generate random users all")

	db := generatePerson.DbPerson
	var (
		name string
		gender string
	)

	db.QueryRow(`SELECT name, gender FROM users ORDER BY RANDOM() LIMIT 1`).Scan(&name, &gender)
	log.Printf("%s | %s\n", name, gender)

	name_lower := strings.ToLower(name)
	namestr := strings.ReplaceAll(name_lower, " ", "")
	email := namestr + randomString(3) + randomNumbers(5)
	phone := generatePhone()
	birthdate := randomBirthdate()
	return map[string]interface{}{
		"name": name_lower,
		"gender": gender,
		"email": email,
		"phone": phone,
		"ttl": birthdate,
	}
}

func randomUserLaki() map[string]interface{} {
	log.Println("generate random users laki laki")

	db := generatePerson.DbPerson
	var (
		name string
		gender string
	)

	db.QueryRow(`SELECT name, gender FROM users WHERE gender = 'L' ORDER BY RANDOM() LIMIT 1`).Scan(&name, &gender)
	log.Printf("%s | %s\n", name, gender)

	name_lower := strings.ToLower(name)
	namestr := strings.ReplaceAll(name_lower, " ", "")
	email := namestr + randomString(3) + randomNumbers(5)
	phone := generatePhone()
	birthdate := randomBirthdate()
	return map[string]interface{}{
		"name": name_lower,
		"gender": gender,
		"email": email,
		"phone": phone,
		"ttl": birthdate,
	}
}


func randomUserWanita() map[string]interface{} {
	log.Println("generate random users wanita")

	db := generatePerson.DbPerson
	var (
		name string
		gender string
	)

	db.QueryRow(`SELECT name, gender FROM users WHERE gender = 'P' ORDER BY RANDOM() LIMIT 1`).Scan(&name, &gender)
	log.Printf("%s | %s\n", name, gender)

	name_lower := strings.ToLower(name)
	namestr := strings.ReplaceAll(name_lower, " ", "")
	email := namestr + randomString(3) + randomNumbers(5)
	phone := generatePhone()
	birthdate := randomBirthdate()
	return map[string]interface{}{
		"name": name_lower,
		"gender": gender,
		"email": email,
		"phone": phone,
		"ttl": birthdate,
	}
}


func GenerateImages() []byte {
	db := generatePerson.DbImage
	var poto []byte
	db.QueryRow(`SELECT image FROM Images ORDER BY RANDOM() LIMIT 1`).Scan(&poto)
	return poto
}



type DataResultGenerate struct {
	Global map[string]interface{} `json:"global"`
	Laki map[string]interface{} `json:"laki"`
	Wanita map[string]interface{} `json:"wanita"`
}


func GetDataAll() DataResultGenerate {
	userall := randomUserAll()
	laki := randomUserLaki()
	wanita := randomUserWanita()
	return DataResultGenerate{
		Global: userall,
		Laki: laki,
		Wanita: wanita,
	}
}