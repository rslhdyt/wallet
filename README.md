# Wallet
The wallet is a wallet application that allows you to store and manage your digital assets.

Endpoints:
- CRUD users
- CRUD wallets
- CRUD cards

Rate Limit:
Currently, the rate limit middlware applied to all endpoints, see the limit.go file.

## Prerequisites
- go 1.17.*
- mysql 5.7.*
- docker (optional)

## Installation
- Clone the repository

## How to run using docker (recomended)
- go to the repository root
- run `docker-compose up -d`

## Running Locust
- go to the locust web interface
  `http://localhost:8089`

### TODO

- implement authentication
