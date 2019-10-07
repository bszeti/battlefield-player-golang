export BATTLEFIELD_MAX_HEALTH=3
export BATTLEFIELD_HIT_PERIOD_MS=2000
export BATTLEFIELD_PLAYER_NAME=test
export BATTLEFIELD_PLAYER_URLS=a,b,c
export TERMINATION_LOG_PATH=./termlog
 
go build && \
./battlefield-player-golang 
