package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const (
	myKey = "xtron46cfabrica460123456789"
)

func main() {

	fmt.Println("Debugg 1: Con Llave correcta")
	dataKey, err := ioutil.ReadFile("key.txt")

	if err != nil {
		log.Fatal(err)
	}

	dataFile, err := ioutil.ReadFile("image.jpg")

	if err != nil {
		log.Fatal(err)
	}

	err = encript(&dataFile, string(dataKey))

	if err != nil {
		log.Fatal(err)
	}

	err = crearArchivoEncriptado(dataFile, "image_encript.jpg")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\tSe encripto con éxito")

	err = desencript(&dataFile, string(dataKey))

	if err != nil {
		log.Fatal(err)
	}

	err = crearArchivoEncriptado(dataFile, "image_desencript.jpg")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\tSe desencripto con éxito")
	fmt.Println()

	//---------------------------------------------------

	fmt.Println("Debugg 2: Con Llave incorrecta")
	dataKey, err = ioutil.ReadFile("falsekey.txt")

	if err != nil {
		log.Fatal(err)
	}

	err = encript(&dataFile, string(dataKey))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\tSe encripto con éxito")

	err = desencript(&dataFile, string(dataKey))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\tSe desencripto con éxito")
	fmt.Println()

}

func encript(b *[]byte, key string) (err error) {

	if key != myKey {
		err = errors.New("llave invalida")
		return
	}

	for i := range *b {

		bInt := int((*b)[i])

		bInt++

		(*b)[i] = byte(bInt)

	}
	return
}

func desencript(b *[]byte, key string) (err error) {

	if key != myKey {
		err = errors.New("llave invalida")
		return
	}

	for i := range *b {

		bInt := int((*b)[i])

		bInt--

		(*b)[i] = byte(bInt)

	}
	return
}

func crearArchivoEncriptado(dataFile []byte, nameFile string) (err error) {

	newFile, err := os.Create(nameFile)

	if err != nil {
		return
	}

	defer newFile.Close()

	_, err = newFile.Write(dataFile)

	if err != nil {
		return
	}

	return
}
