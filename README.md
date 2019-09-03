# mpg-tracker

Input and view your car's MPG data

## Setup

- Prerequisites:

  - ensure [Docker](https://docs.docker.com/install/) is installed and configured

  - ensure [Go](https://golang.org/doc/install) is installed and configured

- create an `.env` file in the root directory

  - set environment varibles as follows, replacing `<value>` with your desired value

  ```config
  COUCHBASE_ADMINISTRATOR_USERNAME=<value>
  COUCHBASE_ADMINISTRATOR_PASSWORD=<value>
  COUCHBASE_BUCKET=<value>
  COUCHBASE_BUCKET_PASSWORD=<value>
  ```

- run `docker-compose up -d` to start container with the Couchbase database

  - with container running you can navigate to `http://localhost:8091` for Administration

- run `go install`

- run `mpg-tracker` to see usage
