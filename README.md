# cake_store
cake store is a sample REST API build using echo server.
The code implementation was inspired by port and adapter pattern or known as [hexagonal](https://blog.octo.com/en/hexagonal-architecture-three-principles-and-an-implementation-example)

-   **Business**
Contains all the logic in domain business. Also called this as a service. All the interface of repository needed and the implementation of the service itself will be put here.
-   **Modules**
Contains implementation of interfaces that defined at the business (also called as server-side adapters in hexagonal's term)
-   **App**
App is handler or controller (also called user-side adapters in hexagonal's term)

## Installation
1. Create `cake_store` in mysql database
2. `$ cp .env.example .env` 
3. `$ go run app/migrationMySQL/main.go `
4. `$ go run app/main/main.go`

if you run the installation correctly the application should run in port 4001 (default)

## Testing
Cause we use hexagonal the business logic in domain business should able to test and make higher code covered by test. to ensure there are no redundant or unwanted code (even can not be test) Execute this command `$ make cover-html`.

## API Doc
[![Run in Postman](https://run.pstmn.io/button.svg)](https://documenter.getpostman.com/view/1806312/2s8YswRBYf)
