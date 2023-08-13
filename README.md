# Volta 🔐

> Password vault API

## About

Password Vault 🔐 built safelly that runs everywhere.

## Features

Main features:

- [ ] api has CRUD oppearations on entries
- [ ] repo factory can interact with different types of database
- [ ] database can be set on local file, docker or remote
- [ ] api can be accessed with cli, app or custom
- [ ] initial configuration can setup service
- [ ] possiblity to have multiple logins per domain

API consists of CRUD operations on password:

- [ ] create with given password
- [ ] create with random password
- [ ] read all entries
- [ ] read entries that match with given domain sub-string
- [ ] update entry with the given fields, keep the not given
- [ ] delete entry

Base **Entry** model:

- id: uuid unique not null
- domain: string unique not null
- login: string
- password: string not null
- created: date not null auto create
- updated: date not null auto update
- meta: string

Storage:

- [ ] The cloud provider choice is done on the configuration file

## Encryption:

- The database is stored **always** encrypted

## Libraries

This project is only possible due to the great effort of those libraries:

- [dbdiagram](https://dbdiagram.io): Create db diagrams
- [golang-migrate](https://github.com/golang-migrate/migrate): Database migrations
- [sqlc](https://sqlc.dev/): Compile SQL to type-safe Go

