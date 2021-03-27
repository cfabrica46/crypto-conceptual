package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
)

type option string

const (
	myKey = "xtron46cfabrica460123456789"
)

const (
	encript    option = "+"
	desencript option = "-"
)

func main() {

	fmt.Println("Debugg 1: Con Llave correcta")
	dataKey, err := ioutil.ReadFile("key.txt")

	if err != nil {
		log.Fatal(err)
	}

	s := "hola"

	b := []byte(s)

	fmt.Printf("\tOrigen:%s\n", b)

	err = myCrypto(&b, encript, string(dataKey))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\tEncriptado: %s\n", b)

	err = myCrypto(&b, desencript, string(dataKey))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\tDesencriptado: %s\n", b)
	fmt.Println()

	//---------------------------------------------------

	fmt.Println("Debugg 2: Con Llave incorrecta")
	dataKey, err = ioutil.ReadFile("falsekey.txt")

	if err != nil {
		log.Fatal(err)
	}

	s = "hola"

	b = []byte(s)

	fmt.Printf("\tOrigen:%s\n", b)

	err = myCrypto(&b, encript, string(dataKey))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\tEncriptado: %s\n", b)

	err = myCrypto(&b, desencript, string(dataKey))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\tDesencriptado: %s\n", b)

}

func myCrypto(b *[]byte, sign option, key string) (err error) {

	if key != myKey {
		err = errors.New("llave de autenticacion incorrecta")
		return
	}

	switch sign {
	case "+":

		for i := range *b {

			bInt := int((*b)[i])

			bInt++

			(*b)[i] = byte(bInt)

		}

	case "-":

		for i := range *b {

			bInt := int((*b)[i])

			bInt--

			(*b)[i] = byte(bInt)

		}

	default:

		err = errors.New("signo invalido")

		return
	}

	return
}
