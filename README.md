# go-pdf-builder

API escrita em Go que recebe arquivos **.txt** e transforma-os em arquivos **.pdf** 

### 📋 Pré-requisitos

Antes de começar, você vai precisar ter instalado em sua máquina a linguagem Go. As instruções estão disponíveis neste [link](https://go.dev/doc/install). Você pode verificar a instalação digitando o comando:

```
which go
```

Em seu terminal. A resposta deverá ser o path onde o compilador se encontra.

Também será necessário ter instalado o sistema de controle de versões Git. As instruções de instalação estão neste [link](https://git-scm.com/book/pt-br/v2/Come%C3%A7ando-Instalando-o-Git)

## 🚀 Começando

- Clone o repositório **go-pdf-builder** em seu computador:

```
git clone git@github.com:ddvalim/go-pdf-builder.git
```

- Instale as dependências do projeto:

```
go get ./...
```

## 🔧 Execução

Para executar o servidor:

```
go run main.go
```

A API estará executando na porta 8080.

## 📍 Rotas

- Enviar um .txt para obter um .pdf:

```
curl --request POST \
  --url http://localhost:8080/pdf \
  --header 'Content-Type: multipart/form-data' \
  --form sample={your_file.txt}
```

## 🛠️ Construído com

* [gofpdf](https://github.com/jung-kurt/gofpdf/tree/master) - Biblioteca utilizada para criar arquivos PDF
* [Mux](https://pkg.go.dev/github.com/gorilla/mux) - Request router

## ✒️ Autora

Este projeto foi desenvolvido por Diovana Rodrigues Valim, bacharela em Sistemas de Informação pela Universidade Federal de Santa Catarina & Engenheira de Software no Mercado Livre.

---
