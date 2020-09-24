package main

type Config map[string]struct {
	TempPath string
	BuildCmd string
	RunCmd   string
}

type CommandConfig struct {
	comment  string
	filepath string
	isPipe   bool
}
