package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dshills/my-ai/assistant"
	"github.com/dshills/my-ai/config"
)

const defConfig = ".my-ai.conf"

func main() {
	var err error
	home := os.Getenv("HOME")
	path := fmt.Sprintf("%v/%v", home, defConfig)
	path, err = filepath.Abs(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conf, err := config.Load(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	assistant, err := assistant.NewAssistant(
		conf.APIKey,
		conf.Model,
		conf.Name,
		conf.Instructions,
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(assistant.Info())

	if err := assistant.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
