FROM golang

WORKDIR /go/src/
RUN git clone https://github.com/gophergala2016/Festival.git
WORKDIR /go/src/Festival

RUN go get github.com/HouzuoGuo/tiedot
RUN go get github.com/russross/blackfriday

RUN go build
EXPOSE 3001
ENTRYPOINT /go/src/Festival/Festival
