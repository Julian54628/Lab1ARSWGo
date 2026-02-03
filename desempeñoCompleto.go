package main

import (
	"fmt"
	"runtime"
	"time"
)

func RunDesempeñoCompleto() {
	validador := HostBlackListsValidator{}

	nucleos := runtime.NumCPU()

	configs := []int{1, nucleos, nucleos * 2, 50, 100}

	fmt.Println("tiempos de ejecucion para la IP 202.24.34.55:")
	for _, hilos := range configs {
		inicio := time.Now()
		validador.checkHost("202.24.34.55", hilos)
		duracion := time.Since(inicio)

		fmt.Printf("%d hilos: %v\n", hilos, duracion)
	}
}
