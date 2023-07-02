curl localhost:8080/debug/pprof/heap -o mem_out.txt
curl localhost:8080/debug/pprof/profile?seconds=15 -o cpu_out.txt

go tool pprof -svg -alloc_objects main mem_out.txt > mem_ao.svg
go tool pprof -svg main cpu_out.txt > cpu.svg