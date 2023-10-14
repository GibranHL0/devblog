FROM golang:1.21.2-alpine AS builder
WORKDIR /bin/app
COPY . .
RUN go mod download && go mod verify
RUN go build -v -o /bin/app

FROM golang:1.21.2-alpine AS runner
WORKDIR /bin/app
LABEL version="1.0"
LABEL maintainer="gibran.devops@gmail.com"
COPY --from="builder" /bin/app/devblog .
COPY --from="builder" /bin/app/templates/files ./templates/files
COPY --from="builder" /bin/app/static ./static
COPY --from="builder" /bin/app/.env .
EXPOSE 8000
CMD [ "/bin/app/devblog" ]
