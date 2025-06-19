Here's the updated README with the formatting improved and the text simplified for beginners to better understand what the project does and how to set it up:

---

# Joke and Meme API Service

### What is this project?

This project is a simple web service built with Go. It fetches programming jokes and memes from external APIs and stores jokes in a PostgreSQL database. You can access these jokes and memes through HTTP endpoints.

### Key Features:

1. **Fetch a Programming Joke**: The application fetches a random programming joke from the JokeAPI and saves it in a PostgreSQL database.
2. **Fetch a Programming Meme**: The application fetches a random programming meme from the Meme API and displays it as an HTML page.

---

### Prerequisites

Before running the application, make sure you have **Docker** installed on your machine. Docker will help you run the application and its dependencies easily in containers.

---

### How to Set Up and Run the Project Using Docker

Follow these steps to get the project running on your local machine using Docker.

1. **Clone the Repository:**
   First, clone the repository to your local machine:

   ```bash
   git clone https://github.com/Achanandhi-M/golang-api-project.git
   cd golang-api-project
   ```

2. **Build and Run the Application Using Docker Compose:**
   Docker Compose is used to set up both the Go application and PostgreSQL database. To build and start the application, run:

   ```bash
   docker-compose up
   ```

   This command will download the required images, build the Go application, and start the containers for both the Go app and PostgreSQL database.

3. **Access the Application:**
   Once the containers are running, you can access the application in your browser:

   * The Go application will be available at: [http://localhost:8080](http://localhost:8080)
   * PostgreSQL will be available at: `localhost:5432` (for accessing the database).

---

### Available Endpoints

Once the application is running, you can interact with it using the following HTTP endpoints:

* **`GET /joke`**:

  * Fetches a random programming joke from JokeAPI.
  * Saves the joke in the PostgreSQL database.
  * Returns a JSON response with the joke data.

* **`GET /meme`**:

  * Fetches a random programming meme from Meme API.
  * Displays the meme in an HTML page with a link to the original post on Reddit.

---

### How to Access the Database

To interact with the PostgreSQL database, you can run commands inside the database container. Here’s how you can do that:

1. Open a terminal and run the following command to access the PostgreSQL container:

   ```bash
   docker-compose exec postgres psql -U postgres -d jokes_db
   ```

2. Inside the database, you can run the following commands:

   * **List tables**:

     ```bash
     \dt
     ```

   * **See the first 10 jokes**:

     ```bash
     SELECT * FROM jokes LIMIT 10;
     ```

   * **Exit**:

     ```bash
     \q
     ```

---

### Conclusion

This project is a beginner-friendly API service that demonstrates how to fetch and display programming jokes and memes using Go, external APIs, and a PostgreSQL database. Docker Compose simplifies the process of running both the application and the database in isolated containers.

By following the instructions above, you can quickly get the service up and running on your local machine.
