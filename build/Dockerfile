# Build the mattermost reminders
ARG DOCKER_BUILD_IMAGE=golang:1.17.0
ARG DOCKER_BASE_IMAGE=alpine:3.14

FROM ${DOCKER_BUILD_IMAGE} AS build
WORKDIR /mattermost-reminders/
COPY . /mattermost-reminders/
RUN apt-get update -yq && apt-get install -yq unzip
RUN make build


# Final Image
FROM ${DOCKER_BASE_IMAGE}
LABEL name="Mattermost Reminders" \
  maintainer="cloud-team@mattermost.com" \
  vendor="Mattermost" \
  distribution-scope="public" \
  url="https://mattermost.com" \
  io.k8s.description="Mattermost Reminders is responsible for sending reminder notifications to Mattermost" \
  io.k8s.display-name="Mattermost Reminders"

ENV CLOUD=/mattermost-reminders/main \
    USER_UID=10001 \
    USER_NAME=cloud
    
RUN  apk update && apk add libc6-compat && apk add ca-certificates && apk --no-cache add tzdata
COPY --from=build /mattermost-reminders/build/_output/bin/main /mattermost-reminders/main
COPY --from=build /mattermost-reminders/build/bin /usr/local/bin

RUN  /usr/local/bin/user_setup
WORKDIR /mattermost-reminders/

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
