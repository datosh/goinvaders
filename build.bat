go get github.com/rakyll/statik
statik -src=assets -include=*.png,*.mp3
go build ./cmd/spaceinvaders/
spaceinvaders.exe