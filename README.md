Car Gasoline Tracker
====================

I wanted a web service to track how my car is doing for gas mileage.

## API

### Locations

Aka, gas stations.

GET /locations - list locations
GET /locations/:id - fetch a specific location
GET /locations/search/:address - search for an address

POST /locations - create a location

### Records

One for each time I fill up my car.

GET /records - list records
GET /records/:id - fetch a specific record

POST /records - store the data from a receipt
