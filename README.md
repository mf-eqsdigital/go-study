# go-study
(...)
Para o projeto de introdução de Go espera-se que implementes um servidor com API HTTP que se conecte a uma base de dados SQLite. Esse servidor devera implementar uma rota que incrementa um contador armazenado em base de dados e responde o valor atual do contador.

Opcionalmente podes ainda explorar como implementar Log de acesso nesta aplicação, uso de ORM (e.g. GORM) para o acesso á base de dados, manipulação de dados em JSON, etc. A proposta acima serve apenas como uma proposta base para explorares a linguagem.

Posteriormente iremos para uma segunda fase, onde iremos explorar Linux e deploy remoto, onde se espera instalares Linux (Ubuntu server) numa VM e implementes scripts de deploy remoto do programa Go que desenvolveste. Mas apos finalizares a primeira parte discutimos com mais detalhe esta etapa.
(...)

<br/>https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql
<br/> create repository, clone it:
<br/> $ git clone YOUR_REPO_URL
<br/> $ cd YOUR_REPO_DIRECTORY

<br/>Initialize the Go modules with your GitHub repository address:
<br/> $ go mod init github.com/REST_OF_YOUR_REPO_ADDRESS
<br/> $ go get -u github.com/gorilla/mux 

<br/>https://tutorialedge.net/golang/golang-orm-tutorial/
<br/> $ go get -u github.com/jinzhu/gorm

<br/>Run in Visual Code Terminal
<br/> go run main.go model.go db_connection.go

<br/>Test app with postman
<br/> GET - http://localhost:8081/counter
<br/> PUT - http://localhost:8081/counter

<br/>Known issues
<br/> cgo: exec /missing-cc: exec: "/missing-cc": file does not exist
<br/> solution: install the "TDM-GCC" -> https://jmeubank.github.io/tdm-gcc/download/
