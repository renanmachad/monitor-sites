package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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
			fmt.Println("Adicione sites à lista de monitoramento")
			addSites()
			os.Exit(0)
		case 3:
			fmt.Println("Encerrando programa")
			os.Exit(0)
		case 4:
			ReadLog()
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
	fmt.Println("2- Adicionar sites monitorados")
	fmt.Println("3- Fechar programa")
	fmt.Println("4- Exibir logs")
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

// restante do código omitido

func testaSite(site string) {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Error", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		WriteLog(site, true)

	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		WriteLog(site, false)
	}
}

func initMonitoring() {
	fmt.Println("Iniciando monitoramento")

	// slice of sites to monitoring
	sites := readFile()

	// map the slice of websites
	for i, site := range sites {
		fmt.Println("Verificando site:", "- ", i, site)
		testaSite(site)
	}

}

func readFile() []string {
	var sites []string
	file, err := os.Open("./config/sites.txt")

	if err != nil {
		fmt.Println("Um erro ocorreu", err)

	}
	reader := bufio.NewReader(file)

	for {

		line, error := reader.ReadString('\n')

		line = strings.TrimSpace(line)

		sites = append(sites, line)
		if error == io.EOF {
			break
		}

	}
	file.Close()
	return sites
}

func WriteLog(site string, status bool) {
	// func to register status in a log file
	arquivo, err := os.OpenFile("./config/logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05 ") + site + " --online: " + strconv.FormatBool(status) + "\n")

	if err != nil {
		fmt.Println("Erro ao escrever arquivo", err)
	}

	arquivo.Close()
}

func ReadLog() {
	file, err := ioutil.ReadFile("./config/logs.txt")

	fmt.Println(string(file))

	if err != nil {
		fmt.Println("Erro ao ler arquivo de log: ", err)
	}

}

func addSites() {
	var new_site string
	file, err := os.Open("./config/config.txt")

	fmt.Scan(&new_site)
	file.WriteString(new_site)

	if err != nil {
		fmt.Println("Erro ao ler ou escrever dados no arquivo")
	}
	file.Close()

}
