package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

type ViaCepResponse struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Cidade     string `json:"localidade"`
	Estado     string `json:"uf"`
}

type BrasilApiResponse struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"street"`
	Bairro     string `json:"neighborhood"`
	Cidade     string `json:"city"`
	Estado     string `json:"state"`
}

type Address struct {
	Cep        string
	Logradouro string
	Bairro     string
	Cidade     string
	Estado     string
	API        string
}

func fetchAPI(url, apiName string, ch chan<- Address) {
	client := http.Client{
		Timeout: 1 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("Erro ao fazer requisição para %s: %v\n", apiName, err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Erro ao ler resposta de %s: %v\n", apiName, err)
		return
	}

	var address Address
	if apiName == "ViaCEP" {
		var data ViaCepResponse
		if err := json.Unmarshal(body, &data); err == nil {
			address = Address{
				Cep:        data.Cep,
				Logradouro: data.Logradouro,
				Bairro:     data.Bairro,
				Cidade:     data.Cidade,
				Estado:     data.Estado,
				API:        "ViaCEP",
			}
		}
	} else {
		var data BrasilApiResponse
		if err := json.Unmarshal(body, &data); err == nil {
			address = Address{
				Cep:        data.Cep,
				Logradouro: data.Logradouro,
				Bairro:     data.Bairro,
				Cidade:     data.Cidade,
				Estado:     data.Estado,
				API:        "BrasilAPI",
			}
		}
	}
	ch <- address
}

func isValidCEP(cep string) bool {
	regex := regexp.MustCompile(`^\d{5}-?\d{3}$`)
	return regex.MatchString(cep)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite o CEP: ")
	cep, _ := reader.ReadString('\n')
	cep = strings.TrimSpace(cep)

	if !isValidCEP(cep) {
		fmt.Println("CEP inválido. Por favor, insira um CEP no formato 00000-000 ou 00000000.")
		return
	}

	ch := make(chan Address, 2)
	go fetchAPI("https://viacep.com.br/ws/"+cep+"/json/", "ViaCEP", ch)
	go fetchAPI("https://brasilapi.com.br/api/cep/v1/"+cep, "BrasilAPI", ch)

	select {
	case address := <-ch:
		fmt.Printf("Resultado da API %s:\nCEP: %s\nLogradouro: %s\nBairro: %s\nCidade: %s\nEstado: %s\n",
			address.API, address.Cep, address.Logradouro, address.Bairro, address.Cidade, address.Estado)
	case <-time.After(1 * time.Second):
		fmt.Println("Erro: Timeout ao buscar o CEP")
	}
}
