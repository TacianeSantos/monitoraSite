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

	boasVindas()

	for {
		menuSistema()
		opcao := leComando()

		switch opcao {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimeLog()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Opção desconhecida")
			os.Exit(-1)
		}
	}
}

func boasVindas() {
	var versao float32 = 1.1
	var nome string = "Sra/Sr"
	fmt.Println("Olá ", nome)
	fmt.Println("Este programa está na versão: ", versao)
}

func menuSistema() {
	fmt.Println("1- Para monitorar site")
	fmt.Println("2- Para verificar logs")
	fmt.Println("0- Sair do programa")
}

func leComando() int {
	var opcao int
	fmt.Scan(&opcao)

	fmt.Println("A escolha do usuário foi:", opcao)

	return opcao
}

const nMonitoramento = 5
const delayParaMonitorar = 5

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	//Slice

	sites := leSitesDoArquivo()

	for i := 0; i < nMonitoramento; i++ {
		for i, site := range sites {
			fmt.Println("Testando: ", i, site)
			testaSite(site)
		}

		time.Sleep(delayParaMonitorar * time.Minute)
	}

	fmt.Println("")

}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		registraLog(site, true)
		fmt.Println("")
		fmt.Println("Retorno de sucesso")
	} else {
		registraLog(site, false)
		fmt.Println("")
		fmt.Println("Error:", resp.StatusCode)
		fmt.Println("Não foi possível acessar o site")
	}
}

func leSitesDoArquivo() []string {

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)
	var sites []string

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()

}

func imprimeLog() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Foi encontrado um erro", err)
	}

	fmt.Println(string(arquivo))
}
