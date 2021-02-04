# go-study
(...)

Para o projeto de introdução de Go espera-se que implementes um servidor com API HTTP que se conecte a uma base de dados SQLite. Esse servidor devera implementar uma rota que incrementa um contador armazenado em base de dados e responde o valor atual do contador.

Opcionalmente podes ainda explorar como implementar Log de acesso nesta aplicação, uso de ORM (e.g. GORM) para o acesso á base de dados, manipulação de dados em JSON, etc. A proposta acima serve apenas como uma proposta base para explorares a linguagem.

Posteriormente iremos para uma segunda fase, onde iremos explorar Linux e deploy remoto, onde se espera instalares Linux (Ubuntu server) numa VM e implementes scripts de deploy remoto do programa Go que desenvolveste. Mas apos finalizares a primeira parte discutimos com mais detalhe esta etapa.


(...)


https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql
create repository, clone it:
$ git clone YOUR_REPO_URL
$ cd YOUR_REPO_DIRECTORY

Initialize the Go modules with your GitHub repository address:
$ go mod init github.com/REST_OF_YOUR_REPO_ADDRESS
$ go get -u github.com/gorilla/mux 

https://tutorialedge.net/golang/golang-orm-tutorial/
$ go get -u github.com/jinzhu/gorm

Run in Visual Code Terminal
go run main.go model.go db_connection.go

Test app with postman
GET - http://localhost:8081/counter
PUT - http://localhost:8081/counter