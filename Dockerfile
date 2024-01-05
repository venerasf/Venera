FROM golang

WORKDIR $GOPATH/src/github.com/farinap5/Venera

COPY . .

RUN go install .

RUN go build .

CMD [ "./venera" ]