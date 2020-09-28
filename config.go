package main

type Config map[string]struct {
	TemplatePath string
	BuildCmd     string
	RunCmd       string
}

type TmpConfig struct {
	Language string
	Filename string
}
