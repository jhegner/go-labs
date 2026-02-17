package main

import(
    "log"
    "os"

    _ "github.com/goinaction/code/chapter2/sample/matchers"
    "github.com/goinaction/code/chapter2/sample/search"
)

// init eh chamado antes mesmo da funcao main
func init() {
    // altera o dispositivo para login no stodut
    log.SetOutput(os.Stdout)
}

// entrypoint para o programa
func main() {
    // realiza a busca do termo especifico
    search.Run("president") 
}

