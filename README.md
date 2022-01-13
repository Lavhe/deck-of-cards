 <h1 align="center">Deck of cards</h1>
  <p align="center">
    A cross-platform api built in GO_LANG that exposes a RESTful interface on port 8000 and allows the user to be able to handle the deck of cards basic operations that can be used in most card games like Poker and Blackjack.
 </p>
<p align="center">
  <a href="https://toggl.com/">
    <img src="https://i.pinimg.com/originals/61/a6/d6/61a6d6164ffa19786136ef2b7ef4b37e.png" alt="Logo" width="30%" height="auto">
  </a>
</p>

<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#features">Features</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgements">Acknowledgements</a></li>
  </ol>
</details>

## About The Project

This is a GO_LANG application that exposes a restful API on port `8000`, It exposes the basic APIs for card games.

It has APIs for Creating a deck, (shuffled and unshuffled) the same API also has an optional parameter to filter the type of cards a user needs i.e **/api/create/?cards=AS,KD,2C,KH**.

### Built With

The programming language used in this project is [GO_LANG](https://golang.org)
The following additional plugins/frameworks are also used.

- [gorilla/mux](https://github.com/gorilla/mux) - HTTP Handler for incoming traffic
- [jwt-go](https://github.com/dgrijalva/jwt-go) - Implementation of JSON Web Tokens used for Auth
- [google/uuid](https://github.com/google/uuid) - The uuid package generates and inspects UUIDs based on RFC 4122 and DCE 1.1: Authentication and Security Services.

## Getting Started

### Prerequisites

- A Computer 😅 (Linux/Windows/Mac)
- [Git](https://git-scm.com/) Installed
- [Go](https://golang.org) Installed
- A web browser, a software to make rest APIs like [PostMan](https://www.postman.com/) or [FireCamp](https://firecamp.io/)
  If using vs-code I recommend the [rest-client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) plugin

### Installation

1. Clone the repo
   ```sh
   git clone git@github.com:Lavhe/deck-of-cards.git
   cd deck-of-cards
   ```
2. Run using the terminal

   ```sh
   # Starts up the application using the default port (8000), and default data folder (./data)
   go run app.go

   # Alternatively, You can provide your own port and data folder
   PORT=8080 DATA_DIR=/tmp go run app.go
   ```

3. Run go tests

   ```sh
   # To run all tests
   go test -cover -v deck-of-cards/...

   # To run only deck tests
   go test -cover -v deck-of-cards/deck_util

   # To run only server tests
   go test -cover -v deck-of-cards/server

   # To run only db tests
   go test -cover -v deck-of-cards/db
   ```

## Usage

##### VS-CODE

I recommend you open the [test.http](./test.http) file to make HTTP requests.

##### CURL

```sh
# Health check - Checks if the application is up and running

curl --request GET \
  --url http://localhost:8000/health
```

```sh
# Create Deck
curl --request POST \
  --url http://localhost:8000/deck/create \
  --header 'content-type: application/json'
```

```sh
# Create Deck shuffled deck
curl --request POST \
  --url http://localhost:8000/deck/create?shuffle=true \
  --header 'content-type: application/json'
```

```sh
# Create Deck filtered
curl --request POST \
  --url http://localhost:8000/deck/create?cards=AS,KD,AC,2C,KH \
  --header 'content-type: application/json'
```

```sh
# Create Deck filtered and shuffled
curl --request POST \
  --url http://localhost:8000/deck/create?cards=AS,KD,AC,2C,KH&shuffle=true \
  --header 'content-type: application/json'
```

```sh
# Open a deck
curl --request POST \
  --url http://localhost:8000/deck/open/890f5d6d-3403-447d-b42b-c166256c2644 \
  --header 'content-type: application/json'
```

```sh

# Draw on card from the deck
curl --request POST \
  --url http://localhost:8000/deck/draw/890f5d6d-3403-447d-b42b-c166256c2644 \
  --header 'content-type: application/json'
```

```sh

# Draw on 3 cards from the deck
curl --request POST \
  --url http://localhost:8000/deck/draw/890f5d6d-3403-447d-b42b-c166256c2644?count=3 \
  --header 'content-type: application/json'
```

## Features

- [x] golang (1.16)
- [x] Unit Tests for all the .go modules
- [x] Health check
- [x] Logger middleware that shows the RemoteAddr (source ip address) and the response time for each request
- [x] Deck of cards functionality
  - [x] Creates a deck of unshuffled 52 cards
  - [x] Creates a deck of shuffled 52 cards
  - [x] Creates a deck of filter cards (shuffled or unshuffled)
  - [x] Gets a deck of cards using a unique deckId
  - [x] Draws one/many cards from the provided deckId

## Roadmap

We will see whats next 🙈

## Contact

Joseph Sirwali - mulavhe@gmail.com

## Acknowledgements

- [gorilla mux](https://github.com/gorilla/mux)
- [google uuid](https://github.com/google/uuid)
- [jwt-go - JSON Web Tokens](https://github.com/dgrijalva/jwt-go)
- [golang ioutil/](https://golang.org/pkg/io/ioutil/)
