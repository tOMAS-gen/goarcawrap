package actions

import (
	"flag"
	"fmt"
	"os"

	"github.com/tOMAS-gen/goarcawrap/wsaa"
	wsfev1_api "github.com/tOMAS-gen/goarcawrap/wsfev1/api"
)

var (
	wsaaCmd = flag.Bool("wsaa", false, "Accede a las opciones de gestión de WSAA.")
	get     = flag.Bool("get", false, "Obtener datos autenticación WSAA.")
	view    = flag.Bool("view", false, "Ver información de WSAA.")
)

func HandleWSAA() {
	// Si se pasó el flag -certificate
	if *wsaaCmd {
		// Llamar a la función que maneja las sub-acciones de certificado
		handleWsaaActions()
	}
}

func helpWsaa() {
	fmt.Fprintf(os.Stderr, "  -%s [opciones]\n", "wsaa")
	fmt.Fprintf(os.Stderr, "    -%s		%v\n", "get", "Obtener datos autenticación WSAA.")
	fmt.Fprintf(os.Stderr, "    -%s		%v\n\n", "view", "Ver información de WSAA.")
}

func handleWsaaActions() {
	// Contar cuántas sub-acciones se especificaron
	actionCount := 0
	if *get {
		actionCount++
	}
	if *view {
		actionCount++
	}

	// Validar que se haya especificado exactamente una sub-acción
	if actionCount == 0 {
		fmt.Fprintln(os.Stderr, "Error: Debes especificar una acción para -wsaa (ej: -get, -view).")
		flag.Usage() // Muestra la ayuda general
		os.Exit(1)
	}
	if actionCount > 1 {
		fmt.Fprintln(os.Stderr, "Error: Solo puedes especificar una acción de certificado a la vez (-get o -view).")
		os.Exit(1)
	}

	// Ejecutar la acción correspondiente
	if *get {
		handleGet()
	} else if *view {
		handleViewWSAA()
	}
	os.Exit(1)
}

func handleGet() {
	data, err := wsaa.GetWSAA(wsfev1_api.ServiceID)
	if err != nil {
		fmt.Println("Error al iniciar sesión en WSAA:", err)
		os.Exit(1)
	}
	fmt.Println("Inicio de sesión en WSAA exitoso.")
	fmt.Println("Token:", data.Token)
	fmt.Println("Sign:", data.Sign)
	fmt.Println("Expires:", data.ExpirationTime)
	os.Exit(1)
}
func handleViewWSAA() {
	data, err := wsaa.GetWSAA(wsfev1_api.ServiceID)
	if err != nil {
		fmt.Println("Error al obtener información de WSAA:", err)
		os.Exit(1)
	}
	fmt.Println("Token:", data.Token)
	fmt.Println("Sign:", data.Sign)
	fmt.Println("Expires:", data.ExpirationTime)
	os.Exit(1)
}
