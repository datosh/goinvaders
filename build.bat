go get github.com/rakyll/statik
statik -src=assets -include=*.png,*.mp3,*.ttf
go build ./cmd/spaceinvaders/
spaceinvaders.exe