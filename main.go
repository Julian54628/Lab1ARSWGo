package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

func main() {
	// Iniciar pprof para monitoreo
	go func() {
		fmt.Println("Monitor pprof disponible en: http://localhost:6060/debug/pprof/")
		http.ListenAndServe(":6060", nil)
	}()

	fmt.Println("\nPARTE I")
	h1 := &CountThread{0, 99}
	h2 := &CountThread{99, 199}
	h3 := &CountThread{200, 299}
	fmt.Println("\nEjecucion start")
	go h1.Run()
	go h2.Run()
	go h3.Run()
	time.Sleep(2 * time.Second)
	fmt.Println("\nEjecucion run")
	h1.Run()
	h2.Run()
	h3.Run()
	time.Sleep(1 * time.Second)
	fmt.Println("\n PARTE II")
	validador := HostBlackListsValidator{}
	fmt.Println("\nPrueba de la clase BlackListThread:")
	var wg sync.WaitGroup
	hilo1 := NuevoHiloSimple(0, 19999, "200.24.34.55")
	hilo2 := NuevoHiloSimple(20000, 39999, "200.24.34.55")
	hilo3 := NuevoHiloSimple(40000, 59999, "200.24.34.55")
	wg.Add(3)
	go hilo1.Ejecutar(&wg)
	go hilo2.Ejecutar(&wg)
	go hilo3.Ejecutar(&wg)
	wg.Wait()
	fmt.Printf("Hilo 1: revision de %d servidores, se encontro %d: %v\n",
		hilo1.CuantosReviso(), hilo1.CuantosEncontro(), hilo1.CualesEncontro())
	fmt.Printf("Hilo 2: revision de %d servidores, se encontro %d: %v\n",
		hilo2.CuantosReviso(), hilo2.CuantosEncontro(), hilo2.CualesEncontro())
	fmt.Printf("Hilo 3: revision de %d servidores, se encontro %d: %v\n",
		hilo3.CuantosReviso(), hilo3.CuantosEncontro(), hilo3.CualesEncontro())

	fmt.Println("\nPruebas con varias cantidades de hilos:")

	fmt.Println("\nIP 200.24.34.55 (hallazgos tempranos encontraods):")
	for _, n := range []int{1, 2, 4} {
		res := validador.checkHost("200.24.34.55", n)
		fmt.Printf("   %d hilos -> %v\n", n, res)
	}

	fmt.Println("\nIP 202.24.34.55 (hallazgos dispersos encontraods):")
	for _, n := range []int{1, 4, 8} {
		res := validador.checkHost("202.24.34.55", n)
		fmt.Printf("%d hilos -> %v\n", n, res)
	}

	fmt.Println("\nIP 212.24.24.55 (no existe en las listas negras):")
	res := validador.checkHost("212.24.24.55", 4)
	fmt.Printf("4 hilos -> %v\n", res)
	fmt.Println("PARTE III")
	for i := 0; i < 5; i++ {
		fmt.Print(".")
		time.Sleep(1 * time.Second)
	}

	fmt.Println()
	RunDesempeñoCompleto()
}
