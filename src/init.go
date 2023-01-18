package src

var Version	float32
var Stable 	bool

func Start(v float32, stb bool) {
	Version = v
	Stable = stb

	x := new(Profile)
	x.BPath = "scripts/"
	x.Globals = make(map[string]string)
	x.Globals["chain"] = "on"

	x.InitCLI()
}