# GoLinkSniffer


Web Crawler is a Go application designed to crawl websites, extract links, and store visited links in a MongoDB database.

## Features

- Visits a starting URL and recursively crawls linked web pages.
- Extracts links from HTML documents and checks if they've been visited before.
- Stores visited links in a MongoDB database.
- Configurable starting URL and MongoDB connection.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Go (Golang) installed on your system.
- A MongoDB server running and accessible.
- FerretDB
- docker-compose installed on your system.


## Installation

1. Clone the repository:

 ```bash
git clone https://github.com/alvarorichard/GoLinkSniffer.git
```

2. Navigate to the project directory:

```bash
cd GoLinkSniffer
```
3. Run the application with docker-compose:

```bash
docker-compose up
```

4. The application will start crawling the configured starting URL and storing visited links in the MongoDB database.
 ```bash
go run main.go
```
or 
```bash
go run main.go -url=https://example.com
```

## Usage

Replace `https://example.com` with the URL you want to start crawling from.

The application will start crawling the specified website, extracting links, and storing visited links in the MongoDB database.

To stop the application, press `Ctrl + C`.

## Configuration

You can configure the following aspects of the application:

MongoDB connection: Modify the MongoDB connection settings in db/db.go.

## Contributing

To contribute to GoLinkSniffer, follow these steps:

1. Fork this repository.
2. Create a new branch for your feature or bug fix.
3. Make changes and commit your changes.
4. Push your changes to the forked repository.






