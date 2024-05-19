FROM golang:1.20

# Create a group named "gophers"
RUN groupadd -r gophers

# Create a user named "gopher" in the "gophers" group
RUN useradd -r -m -g gophers gopher

# Create Work dir
WORKDIR /home/gopher

# Swith current user to "gopher"
USER gopher

# Copy code
COPY --chown=gopher:gophers go.mod *.go .
RUN go mod tidy

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o urlScraper

ENTRYPOINT ["/home/gopher/urlScraper"]