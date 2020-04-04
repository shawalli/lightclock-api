FROM golang:1.14-buster

ARG USERNAME=app
ARG USER_ID=1000
ARG PROJECT=shawalli/lightclock-api

RUN adduser \
    --system \
    --disabled-password \
    --uid ${USER_ID} \
    --group \
    --shell /bin/bash \
    ${USERNAME}

ARG APT_PACKAGES="""vim \
less \
"""

RUN apt-get update && \
    apt-get install --yes ${APT_PACKAGES}

USER ${USERNAME}

WORKDIR /go/src/github.com/${PROJECT}
