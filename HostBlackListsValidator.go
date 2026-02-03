package main

import (
	"fmt"
	"sync"
	_ "sync/atomic"
)

type HostBlackListsValidator struct{}

const Limite = 5

func (v *HostBlackListsValidator) checkHost(ip string, cantidadHilos int) []int {
	datasource := GetInstance()
	totalServers := datasource.GetRegisteredServersCount()

	if cantidadHilos == 1 {
		resultados := []int{}
		contador := 0
		revisados := 0

		for i := 0; i < totalServers && contador < Limite; i++ {
			revisados++

			if datasource.IsInBlackListServer(i, ip) {
				resultados = append(resultados, i)
				contador++
			}
		}

		if contador >= Limite {
			datasource.ReportAsNotTrustworthy(ip)
		} else {
			datasource.ReportAsTrustworthy(ip)
		}

		fmt.Printf("Revisadas las blacklist: %d of %d\n", revisados, totalServers)
		return resultados
	}

	tamanoSegmento := totalServers / cantidadHilos
	resto := totalServers % cantidadHilos

	var encontradosTotal int32 = 0

	hilos := make([]*BlackListThread, cantidadHilos)
	var wg sync.WaitGroup

	inicio := 0
	for i := 0; i < cantidadHilos; i++ {
		fin := inicio + tamanoSegmento - 1
		if i < resto {
			fin++
		}

		hilos[i] = NuevoHiloOptimizado(inicio, fin, ip, &encontradosTotal)
		wg.Add(1)
		go hilos[i].Ejecutar(&wg)
		inicio = fin + 1
	}

	wg.Wait()
	todosResultados := []int{}
	totalRevisados := 0
	for _, hilo := range hilos {
		resultadosHilo := hilo.CualesEncontro()
		todosResultados = append(todosResultados, resultadosHilo...)
		totalRevisados += hilo.CuantosReviso()
	}

	resultadoFinal := todosResultados
	if len(todosResultados) > Limite {
		resultadoFinal = todosResultados[:Limite]
	}

	if len(resultadoFinal) >= Limite {
		datasource.ReportAsNotTrustworthy(ip)
	} else {
		datasource.ReportAsTrustworthy(ip)
	}

	fmt.Printf("Checked Black Lists: %d of %d (using %d threads)\n",
		totalRevisados, totalServers, cantidadHilos)

	return resultadoFinal
}
