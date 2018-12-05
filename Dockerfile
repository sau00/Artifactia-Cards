# build stage
FROM alpine:latest AS build-env

RUN apk add --update musl-dev go git

#RUN mkdir -p /opt/src/artifactia-cards
ADD . /opt/src/artifactia-cards

ENV GOPATH=/opt

RUN cd /opt/src/artifactia-cards && go get -v . && go build -o cards


# final stage
FROM alpine:latest as build-prod
WORKDIR /opt/
COPY --from=build-env /opt/src/artifactia-cards/cards /opt/
COPY --from=build-env /opt/src/artifactia-cards/app/views /opt/app/views
COPY --from=build-env /opt/src/artifactia-cards/static /opt/static
COPY --from=build-env /opt/src/artifactia-cards/uploads /opt/uploads
ENTRYPOINT ./cards

EXPOSE 1233