# receipt-processor-challenge

This is my submission for the receipt processor challenge from Fetch's application process.

In order to run this application ensure you have Docker, Go, and make installed on your machine.

Run `make init` to ensure all dependencies are retrieved and the postgres container is created.

Run `make run` to run the application. 

Run `make clean` to stop and delete the Postgres container. If you do not run this after running the application the created container will remain on your system.