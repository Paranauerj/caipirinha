# Caipirinha

![alt text](https://miro.medium.com/max/250/1*KgxtxcL5vvC6jZMZa9CxQQ.png)

## The fastest code generator for Go APIs



Caipirinha is a simple and fast code generator for building APIs with Go



## Features

- MVC
- ORM
- Gin Based
- Simple Architecture
- Fast Development

## Installation

Caipirinha requires Go 1.12+, though it can work with previous versions
All projects built with Caipirinha support Modules.
To Install it:
```sh
go get github.com/Paranauerj/caipirinha
```

## Usage
Generating things is really easy, just type:
``` sh
caipirinha create [option] name
```

The current options are **Project, Model, Controller and Middleware**

## Tech Stack

Caipirinha is built on top of the following technologies

| Feature | Project |
| ------ | ------ |
| ORM | [Gorp](https://github.com/go-gorp/gorp) |
| Web Framework | [Gin-gonic](https://github.com/gin-gonic/gin) |
| MySQL Driver | [Go-MySQL-Driver](github.com/go-sql-driver/mysql) |
| DotEnv Reader | [GoDotEnv](github.com/joho/godotenv) |
