FROM alpine:3.6

ADD rungopherrun /rungopherrun

ENTRYPOINT ["/rungopherrun"]
