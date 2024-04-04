FROM ubuntu:latest
LABEL authors="nualimov"

ENTRYPOINT ["top", "-b"]