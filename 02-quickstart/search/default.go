package search

// implementa o matcher default
type defaultMatcher struct{}

// inicializa registro do matcher com o programa
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

// implementacao de busca para o compartamento do matcher padrao
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error){
	return nil, nil
}

 