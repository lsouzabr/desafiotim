# desafiotim

Clone o repositório usando o comando:
#git clone https://github.com/lsouzabr/desafiotim.git

Estando na pasta do projeto, executar:
#docker compose up

Após executar o comando "docker compose up" a aplicação baixará as dependências e subir os containers

Sugiro usar o software postman e importar as collections.
As collections do postman estão com o post e body configurados, estão na pasta postman, estão só importar.

seguem as urls
URLs:

(post)
http://localhost:8080/sequence
body json:
{
"Letters":["DUHBHB", "DUBUHD", "UBUUHU","BHBDHH","DDDDUB","UDBDUH"]
}

(get)
http://localhost:8080/stats
