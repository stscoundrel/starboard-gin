# Starboard (Gin)

Gin HTTP server to run [Starboard](https://github.com/stscoundrel/starboard). Lists Github users repositories that have received stars.

## Usage

Start server:
`go run main.go`

HTTP GET request to:

`http://localhost:8080/v1/stars/stscoundrel`

or any other Github username.