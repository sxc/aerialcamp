FROM node:latest AS tailwind-builder
WORKDIR /tailwind
RUN npm init -y && \
    npm install tailwindcss && \
    npx tailwind init 

COPY ./templates /templates
COPY ./tailwind/tailwind.config.js /src/tailwind.config.js
COPY ./tailwind/styles.css /src/styles.css
RUN npx tailwindcss -c /src/tailwind.config.js -i /src/styles.css -o  /styles.css --minify



# Start with the official GoLang base image
FROM golang:alpine AS builder
# Set the working directory inside the container
WORKDIR /app
# Copy the GoLang source code to the container
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Build the GoLang application inside the container
RUN go build -v -o ./server ./cmd/server/

FROM alpine
WORKDIR /app
COPY ./assets ./assets
COPY .env .env
COPY --from=builder /app/server ./server
COPY --from=tailwind-builder /styles.css /app/assets/styles.css
CMD ./server