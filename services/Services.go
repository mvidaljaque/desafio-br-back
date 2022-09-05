package services

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	. "desafio-br-back/database"
	. "desafio-br-back/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// URL pokeAPI
const apiurl = "https://pokeapi.co/api/v2/pokemon/"

// slice de pokemones
var pokemones []Pokemon

// estructura de la tabla de base de datos
type registrosPokemon struct {
	gorm.Model
	//id_registros       int    `gorm:"primary_key;auto_increment"`
	id              int    `gorm:"type:int;not null;"`
	nombre          string `gorm:"type:varchar(45);not null;"`
	experiencia     int    `gorm:"type:int;not null;"`
	ataque          int    `gorm:"type:int;not null;"`
	ataque_especial int    `gorm:"type:int;not null;"`
	defensa         int    `gorm:"type:int;not null;"`
	url_sprite      string `gorm:"type:char(45);not null;"`
}

// Funcion para guardar pokemones en la base de datos
func guardarPokemons(poke Pokemon) {
	dbConfig := Configure("./", "mysql")
	DB = dbConfig.InitMysqlDB()

	//fmt.Println("esto se esta tratando de insertar")
	//fmt.Println(poke)
	DB.Create(poke)
	defer DB.Close()
}

// Funcion para leer registros desde la base de datos
func LeerPokemonesDB(w http.ResponseWriter, r *http.Request) {
	dbConfig := Configure("./", "mysql")
	DB = dbConfig.InitMysqlDB()

	fmt.Println("esto se esta leyendo base de datos")

	var result []Pokemon

	// obtener todos los registros
	DB.Find(&result)

	responsePokemons(w, 200, result)
}

// Funcion para leer registros desde la base de datos
func GenerarArchivo(w http.ResponseWriter, r *http.Request) {
	dbConfig := Configure("./", "mysql")
	DB = dbConfig.InitMysqlDB()

	fmt.Println("esto se esta leyendo base de datos")

	var result []Pokemon

	// obtener todos los registros
	DB.Find(&result)
	//fmt.Println(result)
	crearArchivo()

	//Escribir en el archivo todos los registros obtenidos de la db
	escribeArchivo(result)

	//responsePokemons(w, 200, pokemones)
	var resultado Message
	resultado.setMessage("Archivo Generado")
	resultado.setMessage("ok")

	responseMessages(w, 200, resultado)

}

// funcion para obtener json de un pokemon desde poke api mediante id del pokemon
func ObtenerPokemon(id int) {
	//var pokeNum string
	pokeNum := strconv.Itoa(id)
	response, err := http.Get(apiurl + pokeNum)
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(response.Body)

	var pokemon_data PokemonFull
	decoder.Decode(&pokemon_data)

	Poke := Pokemon{
		Id:              pokemon_data.ID,
		Nombre:          pokemon_data.Species.Name,
		Experiencia:     pokemon_data.BaseExperience,
		Hp:              pokemon_data.Stats[0].BaseStat,
		Ataque:          pokemon_data.Stats[1].BaseStat,
		Ataque_Especial: pokemon_data.Stats[3].BaseStat,
		Defensa:         pokemon_data.Stats[2].BaseStat,
		Url_Sprite:      pokemon_data.Sprites.Other.Dream_world.Front_default,
	}

	//cerrar lectura
	defer response.Body.Close()

	//log.Println(pokemon_data)
	pokemones = append(pokemones, Poke)

	//err = collection.Insert(pokemon_data)

	//if err != nil{
	//	w.WriteHeader(500)
	//	return
	//}

	//responsePokemon(w, 200, pokemon_data)
}

// funcion que devuelve json de un pokemon
func responsePokemon(w http.ResponseWriter, status int, results Pokemon) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}

// funcion que devuelve json del listado de pokemones
func responsePokemons(w http.ResponseWriter, status int, results []Pokemon) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}

// funcion que genera array de 10 numeros randoms entre 1 y 493 (cuarta generación de pokemons)
func randomArray(len int) []int {
	a := make([]int, len)
	for i := 0; i <= len-1; i++ {
		a[i] = rand.Intn(493)
	}
	return a
}

// listar 10 pokemon randoms
func PokemonList10(w http.ResponseWriter, r *http.Request) {

	//Generar arreglo con 10 numeros random
	rand.Seed(time.Now().UnixNano())
	a := randomArray(10)

	//reinicializamos slice de pokemones
	pokemones = pokemones[:0]

	//obtenemos pokemons desde la poke api
	for i := 0; i < len(a); i++ {
		ObtenerPokemon(a[i])
	}

	//Guardamos pokemons en la base de datos
	for i := 0; i < len(a); i++ {
		guardarPokemons(pokemones[i])

	}
	// //print por consola de los pokemons obtenidos
	// for i := 0; i < len(pokemones); i++ {
	// 	fmt.Println(pokemones[i])

	// }

	responsePokemons(w, 200, pokemones)
}

// listar los primeros 151 pokemon en un json
func PokemonList151(w http.ResponseWriter, r *http.Request) {

	//reinicializamos slice de pokemones
	pokemones = pokemones[:0]

	//Obtenemos los 151 pokemones desde la pokeApi
	for i := 1; i <= 151; i++ {
		ObtenerPokemon(i)

	}

	//Guardamos pokemons en la base de datos
	for i := 1; i < len(pokemones); i++ {
		//guardar en la base de datos
		guardarPokemons(pokemones[i])

	}

	responsePokemons(w, 200, pokemones)

}

type Message struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (this *Message) setStatus(data string) {
	this.Status = data
}

func (this *Message) setMessage(data string) {
	this.Message = data
}

func responseMessages(w http.ResponseWriter, status int, results Message) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}
