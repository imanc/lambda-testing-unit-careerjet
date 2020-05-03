PROJECT_NAME := lambda-testing-unit-careerjet

run:
	go run main.go

build: clean
	go build -o bin/$(PROJECT_NAME)

clean:
	rm -rf bin