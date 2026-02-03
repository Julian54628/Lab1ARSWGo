package main

import (
	"sync"
	"sync/atomic"
)

type BlackListThread struct {
	inicio       int
	fin          int
	ip           string
	encontrados  []int
	revisados    int
	contadorGral *int32
}

func NuevoHiloSimple(inicio, fin int, ip string) *BlackListThread {
	return &BlackListThread{
		inicio:      inicio,
		fin:         fin,
		ip:          ip,
		encontrados: []int{},
	}
}

func NuevoHiloOptimizado(inicio, fin int, ip string, contador *int32) *BlackListThread {
	return &BlackListThread{
		inicio:       inicio,
		fin:          fin,
		ip:           ip,
		encontrados:  []int{},
		contadorGral: contador,
	}
}

func (h *BlackListThread) Ejecutar(wg *sync.WaitGroup) {
	defer wg.Done()
	datasource := GetInstance()
	for server := h.inicio; server <= h.fin; server++ {
		if h.contadorGral != nil && h.revisados%50 == 0 {
			if atomic.LoadInt32(h.contadorGral) >= 5 {
				break
			}
		}

		h.revisados++
		if datasource.IsInBlackListServer(server, h.ip) {
			h.encontrados = append(h.encontrados, server)
			if h.contadorGral != nil {
				atomic.AddInt32(h.contadorGral, 1)
			}
		}
	}
}

func (h *BlackListThread) CuantosEncontro() int {
	return len(h.encontrados)
}

func (h *BlackListThread) CualesEncontro() []int {
	return h.encontrados
}

func (h *BlackListThread) CuantosReviso() int {
	return h.revisados
}
