go get github.com/rakyll/statik
statik -src=assets -include=*.png,*.mp3,*.ttf
go build -o spaceinvaders.exe ./cmd/spaceinvaders/
spaceinvaders.exe