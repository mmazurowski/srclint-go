package commands

import "srclint/src/config"

type Args struct {
	Path string
}

func InitHandler(args Args) {
	cfn := config.CreateEmpty()

	config.WriteConfig(cfn, args.Path)
}
