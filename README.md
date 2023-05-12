# go-db-bench
Go-DB-Bench is a command line tool for benchmarking the performance of MongoDB and PostgreSQL databases.

## Setting up the Databases

The Go Database Benchmark tool requires a running instance of MongoDB and PostgreSQL to perform benchmarking tests. To make it easy to set up these databases, we provide a Makefile with commands to start and stop them using Docker.

### PostgreSQL Setup

To start a PostgreSQL database, run the following command:

```
make up-postgres
```

This will start a Docker container running PostgreSQL with the following configuration:

- Username: `test`
- Password: `test`
- Database: `test`
- Port: `5432`

To stop the PostgreSQL container, run the following command:

```
make down-postgres
```

### MongoDB Setup

To start a MongoDB database, run the following command:

```
make up-mongo
```

This will start a Docker container running MongoDB with the following configuration:

- Port: `27018`

To stop the MongoDB container, run the following command:

```
make down-mongo
```
## MongoDB Benchmarking
To test the performance of MongoDB, use the `mongo-test` command. This command will run a series of tests against a MongoDB database and output the results to the console.

To generate 1000000 documents in a MongoDB collection, use the `mongo-generate` command.

## PostgreSQL Benchmarking
To test the performance of PostgreSQL, use the `postgres-test` command. This command will run a series of tests against a PostgreSQL database and output the results to the console.

To generate 1000000 records in a PostgreSQL database, use the `postgres-generate` command.

## Configuration
By default, the tool will connect to a MongoDB and PostgreSQL database running on localhost with default ports and no authentication. If you need to connect to a different database, you can modify the default parameters in the config/config.go file.
