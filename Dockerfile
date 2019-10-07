FROM registry.access.redhat.com/ubi7/ubi

ADD ./battlefield-player-golang /

ENTRYPOINT ["/battlefield-player-golang"]