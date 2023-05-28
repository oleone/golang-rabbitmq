FROM golang:1.19.9-alpine3.18

COPY . .

CMD ["ls -la"]