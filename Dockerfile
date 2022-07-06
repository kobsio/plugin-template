# The following Docker image is used to build the frontend for the plugin. This allows us to distribute the frontend
# for the plugin via the created Docker image and we do not have to care about other distribution methods.
FROM --platform=linux/amd64 node:16.13.0 as app
WORKDIR /kobs
COPY . .
# To build the frontend of a plugin we have to call yarn install and then yarn build.
RUN yarn install --frozen-lockfile --network-timeout 3600000
RUN yarn build

# Finally we are using an alpine image to distribute the frontend files. We recommend to copy the "build" folder from
# the first build stage to a folder which the name of the plugin into the "kobs" directory.
#
# The instructions how this plugin can then be used can be found in the kobsio/app-template repository.
FROM alpine:3.16.0
RUN apk update && apk add --no-cache ca-certificates
RUN mkdir /kobs
COPY --from=app /kobs/build /kobs/helloworld
WORKDIR /kobs
USER nobody
