// AUTORES: Javier Fuster Trallero / Javier Herrer Torres
// NIAs: 626901 / 776609
// FICHERO: calculartiempo.go
// FECHA: 04-oct-2020
// TIEMPO: 5'
// DESCRIPCIÓN: calcula el tiempo medio que le cuesta a la máquina calcular los primos hasta 60000

package main

import (
	"../utils"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		utils.FindPrimes(60000)
	}
	elapsed := time.Now().Sub(start)
	fmt.Print(elapsed.Milliseconds() / 10)
}
