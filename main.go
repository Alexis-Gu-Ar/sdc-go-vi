// add pictures of execution
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Multimedia interface {
	mostrar()
}

type Imagen struct {
	titulo  string
	formato string
	canales uint
}

func (i *Imagen) mostrar() {
	fmt.Println("Titulo: " + i.titulo)
	fmt.Println("Formato: " + i.formato)
	fmt.Printf("Canales: %d\n", i.canales)
}

type Audio struct {
	titulo   string
	formato  string
	duracion uint
}

func (a *Audio) mostrar() {
	fmt.Println("Titulo: " + a.titulo)
	fmt.Println("Formato: " + a.formato)
	fmt.Printf("Duracion: %d\n", a.duracion)
}

type Video struct {
	titulo  string
	formato string
	frames  uint
}

func (v *Video) mostrar() {
	fmt.Println("Titulo: " + v.titulo)
	fmt.Println("Formato: " + v.formato)
	fmt.Printf("Frames: %d\n", v.frames)
}

type ContenidoWeb struct {
	Contenido []Multimedia
}

func (c *ContenidoWeb) mostrar() {
	for _, c := range c.Contenido {
		c.mostrar()
		fmt.Println("")
		fmt.Println("|--------------------------------------|")
		fmt.Println("")
	}
}

func main() {
	exit := false

	var myStr1 string
	var myStr2 string
	var myInt uint64

	scanner := bufio.NewScanner(os.Stdin)
	contenidoWeb := ContenidoWeb{
		Contenido: []Multimedia{},
	}

	for !exit {
		fmt.Println("1) Capturar Imagen")
		fmt.Println("2) Capturar Audio")
		fmt.Println("3) Capturar Video")
		fmt.Println("4) Mostrar")
		fmt.Println("5) Salir")
		scanner.Scan()
		opc, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Formato de opcion invalido")
		}

		switch opc {
		case 1:
			fmt.Println("Agregando imagen")
			fmt.Print("Dame el titulo: ")
			scanner.Scan()
			myStr1 = scanner.Text()
			fmt.Print("Dame el formato: ")
			scanner.Scan()
			myStr2 = scanner.Text()
			fmt.Print("Dame los canales: ")
			scanner.Scan()
			myInt, _ = strconv.ParseUint(scanner.Text(), 10, 64)

			imagen := Imagen{
				titulo:  myStr1,
				formato: myStr2,
				canales: uint(myInt),
			}

			contenidoWeb.Contenido = append(contenidoWeb.Contenido, &imagen)

		case 2:
			fmt.Println("agregando audio")
			fmt.Print("Dame el titulo: ")
			scanner.Scan()
			myStr1 = scanner.Text()
			fmt.Print("Dame el formato: ")
			scanner.Scan()
			myStr1 = scanner.Text()
			fmt.Print("Dame la duracion: ")
			scanner.Scan()
			myInt, _ = strconv.ParseUint(scanner.Text(), 10, 64)

			audio := Audio{
				titulo:   myStr1,
				formato:  myStr2,
				duracion: uint(myInt),
			}

			contenidoWeb.Contenido = append(contenidoWeb.Contenido, &audio)

		case 3:
			fmt.Println("agregando video")
			fmt.Print("Dame el titulo: ")
			scanner.Scan()
			myStr1 = scanner.Text()
			fmt.Print("Dame el formato: ")
			scanner.Scan()
			myStr2 = scanner.Text()
			fmt.Print("Dame los frames: ")
			scanner.Scan()
			myInt, _ = strconv.ParseUint(scanner.Text(), 10, 64)

			video := Video{
				titulo:  myStr1,
				formato: myStr2,
				frames:  uint(myInt),
			}

			contenidoWeb.Contenido = append(contenidoWeb.Contenido, &video)

		case 4:
			contenidoWeb.mostrar()
		case 5:
			exit = true
		default:
			fmt.Println("invalid option")
		}
	}
}
