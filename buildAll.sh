#!/usr/bin/env sh
build_master(){
  go build -o master MasterWorker/master.go
}
build_worker(){
  go build -o worker MasterWorker/worker.go
}

build_secuencial(){
  go build -o secuencial Secuencial/servidorsecuencial.go
}

build_concurrente(){
  go build -o concurrente Concurrente/servidorconcurrente.go
}

build_concurrente_pool(){
  go build -o concurrente_pool ConcurrentePool/servidorconcurrentepool.go
}

###
# Construimos todos los ejecutables
###

build_secuencial
build_concurrente
build_concurrente_pool
build_master
build_worker

