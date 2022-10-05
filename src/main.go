package main

import (
	"fmt"
	"os"
)

func main() {
	introduction()
	Options()

	comand := InputComand()

	switch comand {
	case 1:
		fmt.Println("Iniciando monitoramento")
		os.Exit(0)
	case 2:
		fmt.Println("Adicione ou remova os valores")
		os.Exit(0)
	case 3:
		fmt.Println("Encerrando programa")
		os.Exit(0)
	default:
		fmt.Print("Comando inválido")
		os.Exit(0)
	}

}

func Options() {
	fmt.Println("Escolha a opção desejada abaixo")

	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Modificar sites monitorados")
	fmt.Println("3- Fechar programa")
}

func introduction() {
	nome := "Renan"
	version := 1.0
	fmt.Print("Projeto feito por:", nome)
	fmt.Print("Versao:", version)
}

func InputComand() int {
	var comand int
	fmt.Scan(&comand)

	return comand
}
