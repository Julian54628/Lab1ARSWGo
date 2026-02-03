# Lab1ARSWGo
Este repositorio contendrá toda la solución del laboratorio 1 de ARSW en Go.
Autores:
- Julián Arenas
- Miguel Vanegas

## Objetivo del Laboratorio

Implementar y analizar diferentes estrategias de paralelización para resolver un problema real de búsqueda en listas 
negras de direcciones IP, aplicando conceptos de programación concurrente en Go y evaluando el rendimiento según la Ley 
de Amdahl.

## Descripción General

Se desarrolla un sistema de seguridad que valida direcciones IP en miles de listas negras de servidores maliciosos. 
El componente debe:

- Buscar una IP en múltiples servidores de listas negras.
- Reportar como "no confiable" si aparece en al menos 5 listas.
- Implementar búsqueda paralela optimizada.
- Evaluar rendimiento con diferentes configuraciones de hilos.
- Analizar resultados aplicando teoría de concurrencia.

## Estructura del Proyecto

Lab1ARSWGo/
├── main.go                        
├── CountThread.go                  
├── BlackListThread.go               
├── HostBlackListsValidator.go   
├── HostBlacklistsDataSourceFacade.go 
├── desempeñoCompleto.go          
├── go.mod                         
└── README.md    

## Cómo ejecutar el proyecto

En la terminal, colóquese en Lab1ARSWGo y ejecute:
go run .

Para acceder al monitoreo de la parte III, una vez se ejecute el código, abra en el navegador:
http://localhost:6060/debug/pprof/

## Explicación del desarrollo por partes

### Parte I

#### Objetivo:

Comprender la diferencia entre ejecutar de forma concurrente y secuencial.

En Go, para ejecutar de forma concurrente se coloca el prefijo `go` antes de la llamada a la función 
(por ejemplo `go f()`), mientras que para la ejecución secuencial se llama a la función directamente (`f()`). Al usar 
goroutines la salida puede aparecer desordenada por la ejecución concurrente; sin goroutines la salida será ordenada y
secuencial.

### Parte II

#### Objetivo

Implementar un sistema paralelo para buscar IPs maliciosas en listas negras.

#### Implementación

Se creó un nuevo tipo `BlackListThread`, encargado de realizar la búsqueda: iniciar la verificación, buscar la IP en su 
segmento asignado, contar las coincidencias para determinar si alcanza el umbral de 5 y reportar los resultados. Cada 
hilo procesa un segmento específico de servidores.

La clase `HostBlackListsValidator` divide 80,000 servidores entre un número N de hilos, distribuyendo de forma 
equitativa las porciones. Se usa `sync.WaitGroup` para esperar a que todos los hilos terminen, y se reporta el resultado
tan pronto como se encuentren 5 coincidencias.

Para optimizar, se implementó un corte temprano: cuando el contador alcanza 5 o más, se detiene la verificación. Con
esto se reduce el tiempo de búsqueda y se usa `atomic` para acceso seguro al contador.

En el ejemplo con la IP `200.24.34.55` (hallazgos tempranos), aparece en 5 servidores: 23, 50, 200, 500, 1000; por lo 
que se reportó como no confiable.

En el ejemplo con la IP `202.24.34.55` (hallazgos dispersos), aparece en los servidores 29, 10034, 20200, 31000, 70500. 
Este caso requiere más búsqueda, ya que llegó hasta el servidor 70501, mientras que el primero llegó hasta el 1001.

### Parte II.I

Nos preguntan: ¿Cómo modificar la implementación para minimizar consultas cuando ya se encontraron suficientes hallazgos?

Se implementó un corte temprano mediante un contador compartido y una verificación periódica que revisa el estado cada 
vez que se procesan 50 servidores. También se puede usar `context.Context` para cancelar todos los hilos simultáneamente
cuando se alcance el límite.

### Parte III

#### Objetivo

Medir el impacto numérico de hilos en tiempo de ejecución.

Las configuraciones que se probaron con sus debidos tiempos fueron:
- 1 hilo con tiempo de 108.883 ms
- 12 hilos con tiempo de 6.737 ms
- 24 hilos con tiempo de 1.546 ms
- 50 hilos con tiempo de 1.625 ms
- 100 hilos con tiempo de 1.003 ms

En este caso, como no estamos en Java, para Go usamos el servidor de monitoreo pprof.

Algunas métricas que se observan son la memoria, mostrando la escala con cada número de hilos, las goroutines y CPU.
