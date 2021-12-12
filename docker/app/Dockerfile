ARG GOLANG_VERSION

FROM golang:${GOLANG_VERSION} AS build

COPY . .

ENV GOPATH=

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/app /go/app/main.go

FROM scratch AS runtime

COPY --from=build /go/bin/app ./

ARG PORT
ENV PORT=${PORT}

ARG APP_NAME
ENV APP_NAME=${APP_NAME}

ARG ENV
ENV ENV=${ENV}

ARG DB_HOST
ENV DB_HOST=${DB_HOST}

ARG DB_NAME
ENV DB_NAME=${DB_NAME}

ARG DB_USER
ENV DB_USER=${DB_NAME}

ARG DB_PASS
ENV DB_PASS=${DB_PASS}

ARG DB_PORT
ENV DB_PORT=${DB_PORT}

EXPOSE ${PORT}/tcp

CMD ["./app"]
