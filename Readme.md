## Benchmark Gorm and database/sql

It is mainly focussing on comparing gorm and database/sql 

## To get memory profile of inserting 3000 records into db

memory profile will be stored in data/ directory

> `go run cmd/profile/main.go`

## To analyse alloc_space

> `go tool pprof --alloc_space data/mem.pprof`

then type th following in cli (eg: `list RowInsert` )

> `list <functionName>`

## To use web ui

> `go tool pprof -http=:3030 data/mem.pprof`

## To run benchmark

> `cd db`

> `go test -bench=. -benchmem`