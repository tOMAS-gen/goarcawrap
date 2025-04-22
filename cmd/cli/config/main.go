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
	// Si se pasó el flag -h o -help, mostrar la ayuda
	actions.HandleCertificate()
	// Si no se pasó ningún flag conocido o subcomando válido, mostrar la ayud
	fmt.Fprintln(os.Stderr, "Error: Comando no reconocido.")
	os.Exit(0)
}
