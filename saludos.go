package saludos

import (
	"errors"
	"fmt"
	"math/rand"
)

// Devolver un saludo para la persona especificada.
func Hello(name string) (string, error) {

	if name == "" {
		return name, errors.New("nombre vacío")
	}
	// Devuelve un saludo que incluye el nombre en un mensaje.
	message := fmt.Sprintf(RandomFormat(), name)
	return message, nil
}	

// Hellos returns un mapa que asocia cada nombre de la persona
// con un mensaje saludo.
func Hellos(names []string) (map[string]string, error){
	messages := make(map[string]string)

	for _ , name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}
	
	return messages, nil
}

// RandomFormat devulve uno de varios formatos de mensajes de saludo
// que se selecciona de forma aleatoria.
func RandomFormat() string {
	formats := []string {
		"¡Hola, %s! ¡Bienvenido!",
		"¡Que bueno verte, %s!",
		"¡Saludos, %s! ¡Encantado de conocerte!",
	}

	return formats[rand.Intn(len(formats))]
}