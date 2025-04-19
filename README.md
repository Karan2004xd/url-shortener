# URL Shortener

A simple URL shortener built with Go, using the Gin framework. It supports JWT-based authentication, generates unique IDs using Twitter's Snowflake algorithm with Base62 encoding, and stores data in SQLite3. The service also allows custom-named URLs.

## Features

- **Short URL Generation**: Creates compact, Base62-encoded short URLs using Twitter's Snowflake ID generation.
- **Custom URLs**: Supports user-defined custom short URLs.
- **Authentication**: JWT-based authentication for secure access.
- **Database**: Uses SQLite3 for lightweight and efficient data storage.
- **API Framework**: Built with Gin for fast and scalable HTTP routing.

## Project Status

This project is still in development. The core backend functionality is mostly complete, including URL shortening, custom URLs, and authentication. Future plans include:

- Adding a frontend interface for user interaction.
- Deploying the application to a hosting platform.

## Technologies Used

- **Go**: Backend programming language.
- **Gin**: HTTP web framework for routing and middleware.
- **SQLite3**: Lightweight database for storing URLs and user data.
- **JWT**: Token-based authentication for secure access.
- **Snowflake**: Twitter's ID generation algorithm for unique IDs.
- **Base62**: Encoding scheme for compact URL representation.

## Setup and Usage

1. Clone the repository
2. Install dependencies: `go mod download`.
3. Configure and set the environment variables, mentions in the `config/config.go file`
4. Run the application: `go run main.go`

## Next Steps

- Develop a frontend (e.g., using HTML/CSS/JavaScript or a framework like React).
- Deploy the application to a cloud platform (e.g., Heroku, AWS, or Fly.io).
- Add additional features like URL analytics or expiration.

*Contributions and feedback are welcome!*
