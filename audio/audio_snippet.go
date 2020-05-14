a, err := audio.NewContext(48000)
if err != nil {
	log.Panic("NewContext")
}
f, err := ebitenutil.OpenFile("audio/ochnoe.mp3")
if err != nil {
	log.Panic("OpenFile")
}
m, err := mp3.Decode(a, f)
if err != nil {
	log.Panic("Decode")
}
player, err := audio.NewPlayer(a, m)
if err != nil {
	log.Panic("NewPlayer")
}