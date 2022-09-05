package routes

import (
	"net/http"

	. "desafio-br-back/services"

	"github.com/gorilla/mux"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandleFunc)
	}

	return router
}

var routes = Routes{
	// 	Route{
	// 		"Index",
	// 		"GET",
	// 		"/",
	// 		Index,
	// 	},

	Route{
		"PokemonList151",
		"GET",
		"/Pokemones151",
		PokemonList151,
	},

	Route{
		"PokemonList10",
		"GET",
		"/Pokemones10",
		PokemonList10,
	},

	Route{
		"GenerarArchivo",
		"GET",
		"/PokemonesArchivo",
		GenerarArchivo,
	},

	Route{
		"LeerPokemonesDB",
		"GET",
		"/LeerPokemonesDB",
		LeerPokemonesDB,
	},

	// Route{
	// 	"ShowPokemon",
	// 	"GET",
	// 	"/Pokemon/{id}",
	// 	ShowPokemon,
	// },
	// Route{
	// 	"GuardarPokemons",
	// 	"POST",
	// 	"/GuardarPokemon",
	// 	GuardarPokemons,
	// },
	//Route{
	//	"MovieUpdate",
	//	"PUT",
	//	"/pelicula/{id}",
	//	MovieUpdate,
	//},
	//Route{
	//	"MovieRemove",
	//	"DELETE",
	//	"/pelicula/{id}",
	//	MovieRemove,
	//},

}
