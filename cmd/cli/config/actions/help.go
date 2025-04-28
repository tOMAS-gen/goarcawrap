package actions

import (
	"flag"
	"fmt"
	"os"
)

var helpCmd = flag.Bool("help", false, "Muestra la ayuda.")

func Help() {
	fmt.Fprintf(os.Stderr, "\nHerramienta para gestionar configuración y certificados de ARCA.\n\n")
	fmt.Fprintf(os.Stderr, "Uso: %s [opciones]\n\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Opciones:\n\n")
	// Mostrar certificate
	helpCertificate()
	// Mostrar WSAA
	helpWsaa()
	// Mostrar help
	fmt.Fprintf(os.Stderr, "  -help -h		Genera una nueva clave privada y CSR.\n\n")
	// Salir
	os.Exit(1)
}
