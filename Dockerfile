

FROM node:latest AS node_build

# Set the current working directory inside the container
WORKDIR /temp

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum package.json style.css public ./

# Download all node dependencies
RUN npm install

# run tailwind stuff
RUN npx tailwindcss -i ./style.css -o ./dist/tailwind.css

FROM golang:1.21-bullseye as go_build

WORKDIR /app

COPY --from=node_build /temp /app

RUN go mod download
RUN go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main .

FROM gcr.io/distroless/static-debian11

COPY --from=go_build /main .

ADD public ./public

CMD ["./main"]