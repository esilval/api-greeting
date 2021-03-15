# Greeting API

API to perform a greeting in three different languages

Setup
-----
- Clone and open in any IDE of your preference
- Go into the source folder
- Execute the command ``go run main.go``
- Then, open a browser or any REST API Client (Postman) and open ``http://localhost:8080/greet``

You can run this API by following this curl as well:

```bash
curl --location --request GET 'http://localhost:8080/greet?lan=es&name=Nick'
```