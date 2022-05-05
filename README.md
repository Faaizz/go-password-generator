# Secure Password Generator

REST API (and CLI tool) to generate secure passwords written in Go.

## Features
-  Allow specification of minimum length
-  Allow specification of number of special characters
-  Allow specification of number of numeric characters
-  Allow generation of multiple passwords in one go
-  Unit tested (table-driven testing)
-  Automated Kubernetes deployment

## Usage
To use the command-line (CLI) tool:
```shell
# Build the tool for your platform
go build .
# Start REST API Server at http:localhost:8080
./go-password-generator
# Generate a single 8-character password with 2 special characters and 2 numerals
./go-password-generator generatePassword
```

### REST API
To generate passwords via the REST API, you need to first startup the server:
```shell
./go-password-generator
```

Then make a `POST` request to [http://localhost:8080/](http://localhost:8080/) with a JSON body like:
```json
{
    "min_length": 8,
    "special_characters_count": 2,
    "numbers_count": 2,
    "pwds_to_create": 3
}
```
A sample curl request is:
```shell
curl --request POST \
  --url http://localhost:8080/ \
  --header 'Content-Type: application/json' \
  --data '{
	"min_length": 8,
	"special_chars_count": 2,
	"numbers_count": 1,
	"pwds_to_create": 3
}'
```

### CLI Tool
To generate passwords with the CLI tool:
```shell
# Generate 5 8-character passwords with 2 special characters and 1 numeral each
./go-password-generator generatePassword -m 8 -c 2 -n 1 -p 5
# To get usage help
./go-password-generator help
```

### Running Tests
To run tests:
```shell
go test ./...
```

