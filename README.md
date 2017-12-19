### Hashing SHA512 and base64 with Go
## This program uses only the standard library so no need to ``$go get ...`` anything 

Clone the repo
cd into the directory and run the following command 

``$go run main.go``

Now the server is up and running 
While running the server is listening on 127.0.0.1:8081.
The server is set to take a post request with a given password

Either curl with the following command 

``curl --data "password=whatever-your-passord-is-here" -X POST http://localhost:8081``

or with postman send as x-www-form-urlencoded data with key as password and value as whatever-your-password-is

When a connection is made, it pauses the current go routine for 5 seconds before continuing
Then returns the final hashed password

The server will handle many processes asynchronously  

When you decide to kill the server it will finish up final processes and then shutdown 






