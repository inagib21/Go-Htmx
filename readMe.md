

# HTMX and Go Demo

This repository contains a simple demonstration of using HTMX and Go to create a dynamic web application for managing a film list. The application allows users to view a list of films and add new films to the list using a web form.

## Prerequisites

- Go (Golang) installed on your machine
- Basic understanding of HTML, Go, and HTMX

## How to Run

1. Clone this repository to your local machine.

2. Open a terminal and navigate to the cloned directory.

3. Run the following command to start the Go server:
   ```
   go run main.go
   ```

4. Open your web browser and visit `http://localhost:8000` to access the application.

## Features

- **Film List**: The main page displays a list of films with their titles and directors. The film list is initially populated with a few sample films.

- **Add Film**: Users can add new films to the list using the "Add Film" section. Enter the title and director of the film in the input fields and click the "Submit" button. The new film will be added to the list without requiring a page refresh, thanks to HTMX.

## Technologies Used

- Go (Golang): Used to build the backend server.
- HTMX: A library that allows seamless interaction between the client and server using HTML attributes.
- HTML and Bootstrap: Used for creating the user interface.
- Template Rendering: Go's `html/template` package is used for rendering HTML templates.

## Code Structure

- `main.go`: Contains the Go code for setting up the server, handling routes, and managing the film data.
- `index.html`: The HTML template that defines the user interface and includes HTMX attributes for dynamic interactions.

## Acknowledgments

Thank you to : bugbytes-io[https://github.com/bugbytes-io].This demo was created to showcase the integration of HTMX and Go for building dynamic web applications. Feel free to explore, modify, and expand upon this codebase for your own projects. Youtube[https://www.youtube.com/watch?v=F9H6vYelYyU]
