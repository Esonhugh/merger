FROM golang AS builder
LABEL authors="esonhugh"

WORKDIR /build

ADD go.mod .
COPY . .
RUN go build -o doc_merger cmd/main/main.go


FROM golang

WORKDIR /

COPY --from=builder /build/test/ /test/
COPY --from=builder /build/init.sh /init.sh
COPY --from=builder /build/doc_merger /doc_merger
COPY --from=builder /build/go.mod /go.mod
RUN go get github.com/esonhugh/sculptor

CMD ["sh", "/init.sh"]