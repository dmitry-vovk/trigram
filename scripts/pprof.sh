#!/bin/sh

# Generate performance reports as diagrams

REPORTS_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")"; pwd -P)"/performance_reports/"

go tool pprof -png http://localhost:8080/debug/pprof/heap > "${REPORTS_DIR}"heap_usage.png
go tool pprof -png http://localhost:8080/debug/pprof/allocs > "${REPORTS_DIR}"allocations.png
go tool pprof -png http://localhost:8080/debug/pprof/goroutine > "${REPORTS_DIR}"goroutines.png
