[![Go](https://github.com/shivamsinghal212/tiger-sightings-logger/actions/workflows/go.yml/badge.svg?branch=main&event=push)](https://github.com/shivamsinghal212/tiger-sightings-logger/actions/workflows/go.yml)

# **TigerHall Kittens**

## Project Setup:

1. Initialize an ENV file in project root with:\
`  PG_DB="tigers" ` \
`   PG_HOST="localhost"  `\
`   PG_PASSWORD="password" `\
   `PG_USER="shivam"`\
2. Download and install PostgreSQL

3. Run: \
`go mod tidy` \
`go build <projectroot>/cmd/server \`

## Run Tests
`go test <projectroot>/internal/services`

## APIs
###  POST Tiger
`curl --location --request POST 'http://localhost:8001/api/tiger' \
--header 'Content-Type: application/json' \
--data-raw '{
"name": "test9",
"latitude":21.593683,
"longitude":22.593683,
"dob": "2009-01-02",
"last_seen": 1652819700
}'`

### GET All Tigers (with Pagination)
`curl --location --request GET 'http://localhost:8001/api/tigers?page=1&page_size=9' `
 
### POST a sighting (with image)
#### Note: Converts a JPEG image to 240X200
`curl --location --request POST 'http://localhost:8001/api/tiger-sighting/<tiger_id>' \
--form 'latitude="21.593683"' \
--form 'longitude="22.593683"' \
--form 'last_seen="1652819790"' \
--form 'file=@"/Users/shivam/Downloads/jpeg2000-home.jpeg"'`

### GET all Tiger Sighting of a tiger
`curl --location --request GET 'http://localhost:8001/api/tiger-sighting/<tiger_id>?page=1&page_size=100' \`
