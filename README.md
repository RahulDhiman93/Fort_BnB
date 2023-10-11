# FortBnB

FortBnB is a room reservation platform developed using Golang, HTML, CSS, and JavaScript. It uses PostgreSQL for the database. With FortBnB, users can search for rooms, book accommodations, and manage their reservations.

## Features

- User registration and login
- Search for available rooms
- View room details and images
- Make room reservations
- Manage bookings

## Getting Started

### Prerequisites

- Install Golang
- Install Node.js
- Set up PostgreSQL
- Clone the repository:

   ```sh
   git clone https://github.com/yourusername/FortBnB.git
   cd FortBnB
   ```

### Installation

1. Configure the PostgreSQL database in `database.yaml.example` and rename it to `database.yaml`.

2. Install dependencies:

   ```sh
   go mod tidy
   cd ..
   ```

3. Start the server:

   ```sh
   go build -o fortbnb cmd/web/*.go && ./fortbnb -dbname="YOUR_DB_NAME" -dbuser="YOUR_DB_USER" -cache=false -production=false
   ```

4. Access FortBnB at `http://localhost:8080` in your browser.

## Usage

1. Search for rooms and select one.
2. View room details.
3. Make a reservation.
4. Manage your bookings in your profile.
5. log in to the admin dashboard for CRUD operations.

## Contributing

We welcome contributions. Please check our [contribution guidelines](CONTRIBUTING.md).
