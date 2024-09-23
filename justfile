version := `sed -n 's/.*VERSION_STRING = "\(.*\)"/\1/p' utils/version.go`

version:
	@echo "Version: {{version}}"

default:
    @just --list

build:
	mkdir -p dist
	go build -o dist/dutip main.go

clean:
	rm -rf dist

install:
	just bump
	go install

run:
	just build
	./dist/dutip -from "Google Chrome" -to "Brave Browser"


bump:
	go run main.go -bump

dev *args:
	go run main.go {{args}}