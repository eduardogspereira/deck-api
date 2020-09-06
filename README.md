<img src="https://avatars1.githubusercontent.com/u/316475?s=200&v=4" width="120px" height="120px" align="left"/>

# Deck API

API to simulate a deck of cards.

## Table of Contents

- [Dependencies](#dependencies)
- [Build](#build)
- [Run](#run)
  - [.env file](#env-file)
  - [API with docker-compose](#api-with-docker-compose)
  - [API without docker-compose](#api-without-docker-compose)
- [API](#api)
  - [POST /deck](#post-/deck)
  - [GET /deck/:deckID](#get-/deck/deckid)
  - [POST /deck/:deckID/draw](#post-/deck/deckid/draw)
- [Development](#development)

## Dependencies

- [Go ^1.15](https://golang.org/)
- [Docker ^19.03](https://www.docker.com/)
- [docker-compose ^1.26.2](https://docs.docker.com/compose/install/)

## Build

Run the following command to generate an executable for the deck-api.

```bash
$ make build
```

This command will output the executable at `./deck-api`.

## Run

### .env file

You must have a .env file with the following environment variables to run the deck-api:

```bash
PGHOST=...
PGPORT=...
PGUSER=...
PGDATABASE=...
PGPASSWORD=...
```

These are the credentials required to connect with the PostgreSQL database. There is a file called `.env.sample` at the root that can be used to run the deck-api in docker-compose. You'll only need to copy the `.env.sample` to `.env`.

### API with docker-compose

Run the following command to run the deck-api:

```bash
$ make up
```

This command will make the deck-api available at http://localhost:3000. It'll also create a temporary database to help explore the API.

### API without docker-compose

For execute the deck-api without the docker-compose it will be required to install the go packages locally with the following command:

```bash
$ go get
```

Then you can start the API with one of the following commands:

```bash
$ go run server.go
# or
$ make build && ./deck-api
```

This command will make the deck-api available at http://localhost:3000.

## API

### POST /deck

Creates a new deck.

#### Request QueryString params

| Name        | Type     | Required | Description                                                                       |
| ----------- | -------- | -------- | --------------------------------------------------------------------------------- |
| **shuffle** | bool     | false    | Defines if the deck should be shuffled.                                           |
| **cards**   | []string | false    | Defines which cards should be inserted at the deck. The default is all the cards. |

#### Request Example

```
curl -X POST http://localhost:3000/deck?cards=AS,KD,AC,2C,KH&shuffle=true --include

HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Date: Tue, 25 Aug 2020 09:47:27 GMT
Content-Length: 80
```

```json
{
  "deck_id": "6a25b964-1d82-4854-b763-c13fb3838f2b",
  "shuffled": true,
  "remaining": 5
}
```

### GET /deck/:deckID

Open a deck.

#### Request Example

```
curl -X GET http://localhost:3000/deck/6a25b964-1d82-4854-b763-c13fb3838f2b --include

HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Tue, 25 Aug 2020 09:48:00 GMT
Content-Length: 310
```

```json
{
  "deck_id": "6a25b964-1d82-4854-b763-c13fb3838f2b",
  "shuffled": true,
  "remaining": 5,
  "cards": [
    { "value": "KING", "suit": "DIAMONDS", "code": "KD" },
    { "value": "2", "suit": "CLUBS", "code": "2C" },
    { "value": "ACE", "suit": "CLUBS", "code": "AC" },
    { "value": "ACE", "suit": "SPADES", "code": "AS" },
    { "value": "KING", "suit": "HEARTS", "code": "KH" }
  ]
}
```

### POST /deck/:deckID/draw

Draw cards from the deck.

#### Request QueryString params

| Name      | Type | Required | Description                           |
| --------- | ---- | -------- | ------------------------------------- |
| **count** | bool | true     | How many cards to draw from the deck. |

#### Request Example

```
curl -X POST http://localhost:3000/deck/6a25b964-1d82-4854-b763-c13fb3838f2b/draw?count=2 --include

HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Tue, 25 Aug 2020 09:50:42 GMT
Content-Length: 100
```

```json
{
  "cards": [
    { "value": "ACE", "suit": "SPADES", "code": "AS" },
    { "value": "KING", "suit": "HEARTS", "code": "KH" }
  ]
}
```

## Development

Execute the tests with the following command:

```bash
$ make test
```
