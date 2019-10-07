FROM registry.access.redhat.com/ubi7/ubi

ADD ./battlefield-golang-player /

ENTRYPOINT ["/battlefield-golang-player"]