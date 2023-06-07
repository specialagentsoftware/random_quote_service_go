# random_quote_service_go
## Three different language challenge

I decided to write a simple web endpoint that displays random quotes (Author, Category, Quote) when someone hits the endpoint in three different languages. (Python, Go and Rust) just to see how the implementation was different and to figure out what I wanted to start using for other projects. This is the go version of the quotes service. 

### Overview
Basic webserver is using echo due to the simplicty of this project.
The quote_client uses the encoding/csv module from the go standard library.
Also included is a go.mod and the go.sum file. 

The project is broken down into modules for ease of understanding and maintenance. 

The application entry point is server.go. It starts the Echo webserver and instantiates the quote_client.
The quote client handles connecting to the csv file and parsing that into a list of Quote objects that are genrated by the quote_model. It then selects a random quote and utilizes the Quote_model to retrieve a formatted html unordered list of the author, category and quote. This is then passed back to the webserver for display to the requester. 

There are comments in each module about the specifics of what is happening when it makes sense.

### Instalation
I might do some installation notes.