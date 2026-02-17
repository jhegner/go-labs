package search

import (
	"log"
	"sync"
)

// Cria um mapa que registra matchers para busca
var matchers = make(map[string]Matcher)

// Executa a logica de busca
func Run(searchTerm string)

	// Obtem a lista de feeds para realizar a busca
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}
	
	// Cria um channel nao bufferizado para receber o resultado correspondente
	results := make(chan *Result)

	// Define um grupo de espera para que seja possivel processar todos os feeds
	var waitGroup sync.WaitGroup

	// Define o numero de goroutines necessarios esperar em quanto processa feeds individuais
	waitGroup.Add(len(feeds))

	// Inicia uma goroutine para cada feed para buscar o resultado
	for _, feed := range feeds {
		// Obtem um matcher para busca
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// Inicia a goroutine para realizar a busca
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	// Executa uma goroutine para monitorar quando todo trabalho eh concluido (closures)
	go func(){
		// Espera ate tudo ser processado
		waitGroup.Wait()

		// Fecha o canal para indicar para funcao Display 
		// que o programa pode ser encerrado
		close(results)
	}()

	// Inicia a apresentacao do resultado assim que disponivel e retorna depois 
	// que o ultimo resultado eh exibido
	Display(results)