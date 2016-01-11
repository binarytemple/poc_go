.PHONY: clean run_list run_read

all: list_leases read_leases 

clean:
	rm -f list_leases

run_list:
	go run list_leases.go

run_read:
	go run list_leases.go

list_leases:
	go build list_leases.go

read_leases:
	go build read_leases.go
