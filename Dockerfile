

FROM node:latest AS node_build

# Set the current working directory inside the container
WORKDIR /temp

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum package.json style.css ./

# Download all node dependencies
RUN npm install

# run tailwind stuff
RUN npx tailwindcss -i ./style.css -o ./dist/tailwind.css

FROM golang:latest AS go_build

WORKDIR /app

COPY --from=node_build /temp /app

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the workspace
COPY . .

# Build the Go app
RUN go build -o /main 

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/main"]