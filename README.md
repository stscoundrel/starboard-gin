# Starboard (Gin)

Gin HTTP server to run [Starboard](https://github.com/stscoundrel/starboard). Lists Github users repositories that have received stars.

## Live API

App is hosted at [Render.com](https://starboard-sqrw.onrender.com/v1/stars/YOUR_USERNAME_HERE). Just change your username to the request url.

`https://starboard-sqrw.onrender.com/v1/stars/YOUR_USERNAME_HERE`

## Local development

Or in case of local development:

Start server:
`go run main.go`

HTTP GET request to:

`http://localhost:8080/v1/stars/stscoundrel`

or any other Github username.
