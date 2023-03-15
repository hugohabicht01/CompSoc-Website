FROM alpine:3.17

RUN apk add --update --no-cache hugo go npm nodejs

WORKDIR /build

COPY . .

RUN npm install
RUN npm run build


EXPOSE 8000

RUN hugo


WORKDIR /build/src

RUN go build .

CMD ./CompSoc-Website
