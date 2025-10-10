# Golang Web Server Example

This project demonstrates a simple web server written in Go that serves static HTML files and handles a basic form submission.

## Features

- Serves static files (HTML, CSS, JS) from the `static` directory.
- Provides a landing page (`index.html`) with a button to open a form.
- Displays a styled form (`form.html`) for user input (Name, Email, Password).
- Handles form submissions via POST requests to `/form`.

## Project Structure

```
short_projects/
└── webserver/
    ├── main.go
    └── static/
        ├── index.html
        └── form.html
```

## Getting Started

### Prerequisites

- Go 1.18 or higher installed

### Running the Server

1. Navigate to the `webserver` directory:

   ```sh
   cd /path/to/short_projects/webserver
   ```

2. Run the server:

   ```sh
   go run main.go
   ```

3. Open your browser and go to [http://localhost:8080](http://localhost:8080).

### Usage

- The home page (`index.html`) will be displayed.
- Click the **Open Form** button to navigate to the form page.
- Fill out the form and submit.
