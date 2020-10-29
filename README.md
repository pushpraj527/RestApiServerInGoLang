#This is simple Rest Api server which will handle a post request

#To Build image run the command 

-> docker build -t myApiServer:1.0.1 .

#To start the server 

-> docker run -p 8080:8080 myApiServer:1.0.1 

#To verify the server is working or not run the below command

-> curl --location --request POST 'http://localhost:8080' --header 'Content-Type: text/plain' --data-raw '"27"'


