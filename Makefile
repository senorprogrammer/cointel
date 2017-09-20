install:
	go clean
	go install

run:
	go run cointel.go --format=table --persist=true

clean:
	go clean
	rm -f cointel
