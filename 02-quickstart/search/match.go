package search

import (
	"log"
)

// Resultado contem o conteudo da busca
type Result struct {
	Field string
	Content string
}

// O Matcher Define o comportamento requerido pelos tipos que 
// implementam um novo tipo de busca
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}