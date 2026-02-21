package search

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"

// Feed contendo informacoes que precisamos para processar o feed
type Feed struct {
	Name string `json:"site"`
	URI string `json:"link"`
	Type string `json:"type"`
}

// Recupera os feeds de leitura e processa o arquivo de dados de feed
func RetrieveFeeds() ([]*Feed, error) {
	// Abre o arquivo
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}
	
	// Agenda o fechamento do arquivo assim que a funcao retorna
	defer file.Close() // IMPORTANTE 

	// Codifica o arquivo em pedacos menores de ponteiros para os valores dos Feeds
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	// Nao eh necessario checkagem de erro o chamador consegue fazer
	return feeds, err
} 