# dynamic config
ARG             BUILD_DATE
ARG             VCS_REF
ARG             VERSION

# build
FROM            golang:1.18.0-alpine as builder
RUN             apk add --no-cache git gcc musl-dev make
ENV             GO111MODULE=on
WORKDIR         /go/src/moul.io/hacker-typing
COPY            go.* ./
RUN             go mod download
COPY            . ./
RUN             make install

# minimalist runtime
FROM            alpine:3.13.5
LABEL           org.label-schema.build-date=$BUILD_DATE \
                org.label-schema.name="hacker-typing" \
                org.label-schema.description="" \
                org.label-schema.url="https://moul.io/hacker-typing/" \
                org.label-schema.vcs-ref=$VCS_REF \
                org.label-schema.vcs-url="https://github.com/moul/hacker-typing" \
                org.label-schema.vendor="Manfred Touron" \
                org.label-schema.version=$VERSION \
                org.label-schema.schema-version="1.0" \
                org.label-schema.cmd="docker run -i -t --rm moul/hacker-typing" \
                org.label-schema.help="docker exec -it $CONTAINER hacker-typing --help"
COPY            --from=builder /go/bin/hacker-typing /bin/
ENTRYPOINT      ["/bin/hacker-typing"]
#CMD             []
