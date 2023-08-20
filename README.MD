# Movie REST API

This is a simple Go application that sets up a REST API for managing a list of movies. The application uses the [Gorilla Mux](https://github.com/gorilla/mux) router to handle different HTTP requests.

## Features

- Retrieve a list of movies.
- Retrieve a single movie by its ID.
- Create a new movie entry.
- Update an existing movie.
- Delete a movie by its ID.

## Getting Started

1. **Clone the Repository:** Clone this repository to your local machine.

2. **Navigate to the Directory:** Open a terminal and navigate to the directory where you cloned the repository.

3. **Run the Application:** Execute the following command to run the application:

   ```bash
   go run main.go

## Access the API Endpoints

- Retrieve all movies: [http://localhost:8000/movies](http://localhost:8000/movies)
- Retrieve a specific movie: [http://localhost:8000/movies/{id}](http://localhost:8000/movies/{id})
- Create a new movie: Send a POST request to [http://localhost:8000/movies](http://localhost:8000/movies) with a JSON payload containing movie details.
- Update a movie: Send a PUT request to [http://localhost:8000/movies/{id}](http://localhost:8000/movies/{id}) with a JSON payload containing updated movie details.
- Delete a movie: Send a DELETE request to [http://localhost:8000/movies/{id}](http://localhost:8000/movies/{id}).

## API Structure

- **GET /movies:** Retrieve a list of all movies.
- **GET /movies/{id}:** Retrieve details of a specific movie.
- **POST /movies:** Create a new movie entry.
- **PUT /movies/{id}:** Update an existing movie.
- **DELETE /movies/{id}:** Delete a movie.

## Requirements

- Go (Golang) must be installed on your machine to run this application.
- The application uses the Gorilla Mux package. You can install it using the following command:

   ```bash
   go get -u github.com/gorilla/mux
