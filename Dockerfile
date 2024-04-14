#syntax=docker/dockerfile:1

# ------
# Builds the application.
# ------
FROM golang:1.20 AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /shin-psmapi

RUN mkdir -p /logs

CMD ["/shin-psmapi"]

# ------
# Runs the application.
# ------
FROM gcr.io/distroless/base-debian11 AS release
WORKDIR /

COPY --from=build /shin-psmapi /shin-psmapi
COPY --from=build --chown=nonroot:nonroot /logs /logs

EXPOSE 8080

USER nonroot:nonroot
ENTRYPOINT ["/shin-psmapi"]
