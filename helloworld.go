package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoringTimes = 5
const sleepTime = 2

func main() {
	introduction()

	for {
		option := readOption()

		if option == 1 {
			monitoringSites()
		} else if option == 2 {
			fmt.Println("Executando opção 2...")
		} else if option == 0 {
			fmt.Println("Programa Encerrado")
			os.Exit(0)
		} else {
			fmt.Println("Opção Inválida!")
			os.Exit(-1)
		}
	}
}

func introduction() {
	var name string = "Budenga"
	var version float32 = 1.1
	age := 1

	fmt.Println("Olá Mundo! Aplicação", name)
	fmt.Println("Aplicação com", age, "ano(s)")
	fmt.Println("Código versão", version)
	fmt.Println()
}

func readOption() int {
	var option int
	fmt.Println("1 - Monitoramento")
	fmt.Println("2 - opção 2")
	fmt.Println("0 - Encerrar")
	fmt.Scan(&option)
	fmt.Println("Opção escolhida", option)

	return option
}

func monitoringSites() {
	fmt.Println("Monitorando...")
	// sites := []string{"https://www.alura.com.br/", "https://www.alura.com.br/2", "https://www.ig.com.br/"}
	sites := readFile()

	for i := 0; i < monitoringTimes; i++ {
		for i, site := range sites {
			fmt.Println("Monitorando site ", i, ":", site)
			checkSite(site)
		}
		fmt.Println()
		time.Sleep(sleepTime * time.Second)
	}

	fmt.Println()
}

func checkSite(site string) {
	ans, _ := http.Get(site)

	status := ans.StatusCode

	if status == 200 {
		fmt.Println("Site funcionando normalmente")
		recordLog(site, true)
	} else {
		fmt.Println("Erro de requisição no site ", site, "\nStatus Code Error:", status)
		recordLog(site, false)
	}
}

func readFile() []string {
	var sites []string

	file, _ := os.Open("./data/sites.txt")

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		// remove \n, \t and others from the text boundary
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}

func recordLog(site string, status bool) {
	filepath := "./data/log.txt"
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Write error file", filepath)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + "-" + site + "- online:" + strconv.FormatBool(status) + "\n")

	file.Close()
}
