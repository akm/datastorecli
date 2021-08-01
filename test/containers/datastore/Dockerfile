FROM google/cloud-sdk:alpine

RUN apk add --update --no-cache openjdk8-jre \
    && gcloud components install cloud-datastore-emulator beta --quiet

ADD ./entrypoint /datastore/entrypoint

WORKDIR /datastore

RUN chmod +x ./entrypoint

ENTRYPOINT ["./entrypoint"]
