
````markdown
# Joke and Meme API Service

This project provides a simple web service to fetch jokes and memes via HTTP, using a Go application with a PostgreSQL database to store jokes.

## Features
1. **Fetch a Programming Joke**: Fetch a random programming joke from JokeAPI and save it in a PostgreSQL database.
2. **Fetch a Programming Meme**: Fetch a random programming meme from Meme API and display it on the webpage.

## Prerequisites

- **Docker**: Make sure Docker is installed on your machine.

## How to Run the Project Using Docker

### Steps:
1. Clone the repository:
    ```bash
    git clone https://github.com/Achanandhi-M/golang-api-project.git
    cd golang-api-project
    ```

2. Build and run the application using Docker Compose:
    ```bash
    docker-compose up
    ```

3. Once the containers are up and running, the Go application will be available at `http://localhost:8080`, and the PostgreSQL database will be running at `localhost:5432`.

### Available Endpoints:
- **`GET /joke`**: Fetches a random programming joke from JokeAPI and saves it to the PostgreSQL database.
    - Response: JSON with joke data.
- **`GET /meme`**: Displays a random programming meme fetched from Meme API.
    - Response: HTML page displaying the meme.

### Access Database:
You can access the database container via the following command:

```bash
docker-compose exec postgres psql -U postgres -d jokes_db
````

* List tables: `\dt`
* See the first 10 jokes: `SELECT * FROM jokes LIMIT 10;`
* Exit: `\q`


## Conclusion

This project provides an easy-to-use API to fetch and save programming jokes and memes. It uses Go, PostgreSQL, and external APIs for joke and meme retrieval. Docker Compose is used for seamless setup and configuration of the app and the database.

```
