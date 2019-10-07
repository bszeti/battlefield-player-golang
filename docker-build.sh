env GOOS=linux GOARCH=amd64 go build
docker build -t quay.io/bszeti/battlefield-player-golang .
docker push quay.io/bszeti/battlefield-player-golang
rm -f battlefield-player-golang
