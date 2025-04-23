package wsaa

import (
	"time"

	"github.com/tOMAS-gen/goarcawrap/model"
)

func CheckExpirationTimeWSAA(wsaa *model.WSAA) bool {
	// Obtener la hora actual de la región horaria de Argentina
	loc, err := time.LoadLocation("America/Argentina/Buenos_Aires")
	if err != nil {
		return false
	}
	currentTime := time.Now().In(loc)
	
	// Comparar la hora actual con la hora de expiración del loginTicketResponse y devolver true si aún es válido
	expirationTime := wsaa.ExpirationTime

	// Devolver true si la hora actual es anterior a la hora de expiración
	return currentTime.Before(expirationTime)
}