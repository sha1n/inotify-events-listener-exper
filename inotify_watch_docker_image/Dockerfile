FROM ubuntu:18.04

WORKDIR /

RUN apt-get update && \
    apt-get install -y inotify-tools && \
    apt-get install -y netcat


VOLUME /sync-data

CMD echo 4524288 > /proc/sys/fs/inotify/max_user_watches && \
    inotifywait \
    -r \
    -e delete \
    -e delete_self \
    -e moved_from \
    -e moved_to \
    -e close_write \
    --csv \
    /sync-data/outputs -m | \
    netcat -vv -k host.docker.internal 8081

