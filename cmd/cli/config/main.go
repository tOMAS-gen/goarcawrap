package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tOMAS-gen/goarcawrap/cmd/cli/config/actions"
)

func main() {
	// Definir los flags
	flag.Usage = actions.Help
	// Parsear los flags (¡Importante! Hacerlo una sola vez)
	flag.Parse()
	//
	actions.HandleCertificate()
	actions.HandleWSAA()
	// Si no se pasó ningún flag conocido o subcomando válido, mostrar la ayud
	fmt.Fprintln(os.Stderr, "Error: Comando no reconocido.")
	os.Exit(0)
}
