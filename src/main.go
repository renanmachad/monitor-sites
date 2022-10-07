package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	introduction()

	for {
		Options()

		comand := InputComand()

		switch comand {
		case 1:
			initMonitoring()
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
}

// display options
func Options() {
	fmt.Println("Escolha a opção desejada abaixo")

	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Modificar sites monitorados")
	fmt.Println("3- Fechar programa")
}

// just print the introduction
func introduction() {
	nome := "Renan"
	version := 1.0
	fmt.Println("Projeto feito por:", nome)
	fmt.Println("Versao:", version)
}

// recieve the input for the user
func InputComand() int {
	var comand int
	fmt.Scan(&comand)

	return comand
}

func initMonitoring() {
	fmt.Println("Iniciando monitoramento")

	// slice of sites to monitoring
	sites := []string{"https://wavetech-st.com", "https://dynamox.net.com", "https://random-status-code.herokuapp.com/"}

	// map the slice of websites
	for i, site := range sites {
		fmt.Println(i)
		resp, error := http.Get(site)

		if resp.StatusCode == 200 {
			fmt.Println("Site ativo")
		} else {
			fmt.Println("Site com problemas:", resp.StatusCode)
		}

		fmt.Print(error)
	}

}
