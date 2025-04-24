package actions

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/tOMAS-gen/goarcawrap"
	"github.com/tOMAS-gen/goarcawrap/model"
)

// Definimos los flags para las sub-acciones de certificado
// Es importante que estos flags se definan ANTES de llamar a flag.Parse() en main()
// Por convención, los definimos aquí pero podrían estar en main.go si prefieres.
var (
	// Flag principal para activar los comandos de certificado
	certificateCmd = flag.Bool("certificate", false, "Accede a las opciones de gestión de certificados.")
	generate       = flag.Bool("generate", false, "Genera una nueva clave privada y CSR.")
	viewCSR        = flag.Bool("viewCSR", false, "Muestra información del CSR existente (si aplica).")
	delete         = flag.Bool("delete", false, "Eliminar certificados.")
)

var (
	cuit             = flag.String("cuit", "", "Numero de CUIT del cliente.")
	commonName       = flag.String("name", "", "Nombre común para el CSR.")
	organizationName = flag.String("organization", "", "Nombre de la organización para el CSR.")
)

func HandleCertificate() {
	// Si se pasó el flag -certificate
	if *certificateCmd {
		// Llamar a la función que maneja las sub-acciones de certificado
		handleCertificateActions()
	}
}

func helpCertificate() {
	fmt.Fprintf(os.Stderr, "  -%s [opciones]\n", "certificate")
	fmt.Fprintf(os.Stderr, "    -%s		%v\n", "generate", "Genera una nueva clave privada y CSR.")
	fmt.Fprintf(os.Stderr, "    -%s		%v\n", "viewCSR", "Muestra información del CSR existente (si aplica).")
	fmt.Fprintf(os.Stderr, "    -%s		%v\n\n", "delete", "Eliminar certificados.")
}

func handleCertificateActions() {
	// Contar cuántas sub-acciones se especificaron
	actionCount := 0
	if *generate {
		actionCount++
	}
	if *viewCSR {
		actionCount++
	}
	if *delete {
		actionCount++
	}

	// Validar que se haya especificado exactamente una sub-acción
	if actionCount == 0 {
		fmt.Fprintln(os.Stderr, "Error: Debes especificar una acción para -certificate (ej: -generate, -viewCSR, -delete).")
		flag.Usage() // Muestra la ayuda general
		os.Exit(1)
	}
	if actionCount > 1 {
		fmt.Fprintln(os.Stderr, "Error: Solo puedes especificar una acción de certificado a la vez (-generate, -viewCSR, o -delete).")
		os.Exit(1)
	}

	// Ejecutar la acción correspondiente
	if *generate {
		handleGenerate()
	} else if *viewCSR {
		handleViewCSR()
	} else if *delete {
		handleDelete()
	}

	os.Exit(1)
}

func handleGenerate() {

	if *cuit == "" || *commonName == "" || *organizationName == "" {
		fmt.Fprintln(os.Stderr, "Error: Debes especificar -cuit, -name y -organization.")
		flag.Usage() // Muestra la ayuda general
		os.Exit(1)
	}

	cuitInt64, err := strconv.ParseInt(*cuit, 10, 64)

	// Verificar si hubo un error durante la conversión
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: No se pudo convertir el CUIT '%s' a int64: %v\n", *cuit, err)
		// Aquí podrías manejar el error, por ejemplo, saliendo del programa
		os.Exit(1)
	}

	err = goarcawrap.GenerateCertificate(&model.Client{
		CUIT:             cuitInt64,
		CommonName:       *commonName,
		OrganizationName: *organizationName,
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error al generar el certificado:", err)
		os.Exit(1)
	}

	fmt.Println("Certificado generado exitosamente.")
	os.Exit(1)
}

func handleViewCSR() {
	// Leer el contenido del archivo CSR
	csrBytes, err := goarcawrap.GetCSR()
	if err != nil {
		// Manejar el caso específico de archivo no encontrado
		if os.IsNotExist(err) {
			fmt.Fprintln(os.Stderr, "Error: El archivo CSR no existe.")
			os.Exit(1)
		}
		// Otro tipo de error
		fmt.Fprintln(os.Stderr, "Error al leer el archivo CSR:", err)
		os.Exit(1)
	}

	// Imprimir el contenido del archivo a la consola
	fmt.Println("\n--- Contenido de MiPedido.csr ---")
	fmt.Println(string(*csrBytes)) // Convertir bytes a string para imprimir
	fmt.Println("--- Fin del contenido ---")

	// Indicar que el comando ha terminado exitosamente
	fmt.Println("Una vez obtenido el certificado CST, se debe cargar en el portal de AFIP para obtener el archivo CRT.")
	fmt.Println("Guardar el archivo CRT en la carpeta ARCA con el siguiente nombre: MiCertificado.crt")

	os.Exit(0) // Salir exitosamente
}

func handleDelete() {
	err := goarcawrap.DeleteCertificate()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error al eliminar el certificado:", err)
		os.Exit(1)
	}
	fmt.Println("Certificado eliminado exitosamente.")
	os.Exit(1)
}
