package services

import (
	. "desafio-br-back/models"
	"fmt"
	"os"
	"strconv"
)

var path = "C:/Users/Manuel/Documents/repositorio/RepoArchivo/prueba.csv"

func crearArchivo() {

	//eliminar archivo
	os.Remove(path)

	//Crea el archivo
	var file, err = os.Create(path)
	if existeError(err) {
		return
	}
	defer file.Close()

	fmt.Println("Archivo creado exitosamente", path)
}

func escribeArchivo(pokemones []Pokemon) {
	// Abre archivo usando permisos READ & WRITE
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if existeError(err) {
		return
	}
	defer file.Close()

	// Escribe algo de texto linea por linea
	for i := 0; i < len(pokemones); i++ {

		_, err = file.WriteString("ID: " + strconv.Itoa(pokemones[i].Id) + "; " + "Nombre: " + pokemones[i].Nombre + "; " + "Experiencia: " + strconv.Itoa(pokemones[i].Experiencia) + "; " + "HP: " + strconv.Itoa(pokemones[i].Hp) + "; " + "Ataque: " + strconv.Itoa(pokemones[i].Ataque) + "; " + "Ataque Especial: " + strconv.Itoa(pokemones[i].Ataque_Especial) + "; " + "Defensa: " + strconv.Itoa(pokemones[i].Defensa) + "; " + "\n")
		if existeError(err) {
			return
		}
	}

	// Salva los cambios
	err = file.Sync()
	if existeError(err) {
		return
	}
	fmt.Println("Archivo grabado existosamente.")
}

func existeError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}
