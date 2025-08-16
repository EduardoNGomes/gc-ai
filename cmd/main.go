package main

import (
	"bufio"
	"os"

	"github.com/eduardongomes/gcai/cli"
	"github.com/eduardongomes/gcai/config"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	conf := config.NewConfig()
	cli := cli.NewCLI()

	cli.Run(conf, reader)

}
