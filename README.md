## Hashing SHA512 with Go
Clone the repo

``$go run main.go``

While running the server is listening on 127.0.0.1:8081.

When a connection is made, it pauses current go routine for 5 seconds before continuing

Either curl with the following command 

``curl --data "password=whatever-your-passord-is-here" -X POST http://localhost:8081``

or 

with postmon send as url encoded data with key as password and value and whatever-your-pasword-is


