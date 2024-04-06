###  Projeto WebCrawler
## Descrição
Se trata de uma aplicação em Go que busca diversos links dentro de algum site, fazendo visitas neles e recuperando ainda mais links. Os dados são salvos em um banco de dados não-relacional (MongoDB). São implementadas dentro do projeto GoRoutines para sua otimização e tornar a aplicação mais performática.

## Tecnologias
- Go
- MongoDB
- Docker

## Como usar
1. Abra o VSCode ou sua IDE compatível de preferência
2. Clone o repositório
```sh
git clone https://github.com/Mario-Juu/crawler.git
```
3. Crie o banco de dados MongoDB por docker
```sh
docker run -d --name mongodb -p 27017:21017 mongo 
```
4. Baixe as dependências
```sh
go mod tidy
```
5. Dê run na aplicação
```sh
go run main.go -link={adicione o seu link de partida da aplicação}
```


## Objetivo 
Afim de manter os dados dos links recuperados salvos, esses são salvos com o link, o domínio e a data em que o mesmo foi visitado.

Os dados são salvos dessa maneira:
| id | link     | website | visitedDate |
|:---|:-----------------|:----------------|:------|
| 6611b86a362e8a3c27bb7064  | https://research.youtube | research.youtube        | ISODate('2024-04-06T21:02:34.792Z')  |

Caso deseje visitar todos os links no banco de dados pelo Docker, basta seguir os seguintes comandos:
```sh
docker exec -it mongodb /bin/bash
```
```sh
mongosh
```
```sh
show dbs;
use crawler
show collections;
db.visited_links.find({})
```
