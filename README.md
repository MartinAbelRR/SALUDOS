# Saludos en Go

Este paquete proporciona una forma sencilla de retornar saludos personalizados en Go.

## Instalación
Ejecutar el siguiente comando para instalar el paquete:
```bash
go get -u github.com/MartinAbelRR/saludos
```

## Uso
Aquí tienes un ejemplo de cómo utilizar el paquete en tu código:

```go
package main

import (
    "fmt"
    "github.com/MartinAbelRR/saludos"
)

func main() {
    message, err := saludos.Hello("Martin")

    if err != nil {
        fmt.Println("Ocurrió un error:", err)
        return
    }

    fmt.Println(message)
}
```

Este ejemplo improta el paquete github.com/MartinAbelRR/saludos y llama a la función Hello(name string) y retorna un saludo personalizado. Si ocurre un error, se imprime un mensaje de error. 