package main

import (
	"fmt"
	"time"
)

type HostBlacklistsDataSourceFacade struct{}

var instancia *HostBlacklistsDataSourceFacade

func GetInstance() *HostBlacklistsDataSourceFacade {
	if instancia == nil {
		instancia = &HostBlacklistsDataSourceFacade{}
	}
	return instancia
}

func (f *HostBlacklistsDataSourceFacade) GetRegisteredServersCount() int {
	return 80000
}

func (f *HostBlacklistsDataSourceFacade) IsInBlackListServer(numeroServer int, ip string) bool {
	time.Sleep(1 * time.Millisecond)

	if ip == "200.24.34.55" {
		if numeroServer == 23 || numeroServer == 50 || numeroServer == 200 || numeroServer == 500 || numeroServer == 1000 {
			return true
		}
		return false
	}

	if ip == "202.24.34.55" {
		if numeroServer == 29 || numeroServer == 10034 || numeroServer == 20200 || numeroServer == 31000 || numeroServer == 70500 {
			return true
		}
		return false
	}

	if ip == "202.24.34.54" {
		if numeroServer == 39 || numeroServer == 10134 || numeroServer == 20300 || numeroServer == 70210 {
			return true
		}
		return false
	}
	return false
}

func (f *HostBlacklistsDataSourceFacade) ReportAsNotTrustworthy(host string) {
	fmt.Printf("INFO: El host %s fue reportado como no confiable\n", host)
}

func (f *HostBlacklistsDataSourceFacade) ReportAsTrustworthy(host string) {
	fmt.Printf("INFO: El host %s fue reportado como confiable\n", host)
}
