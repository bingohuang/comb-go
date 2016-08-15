package cc

import "github.com/codegangsta/cli"

var (
	containerFlags = []cli.Flag{
		cli.BoolFlag{
			Name:  "i",
			Usage: "List all containers' images.",
		},
		cli.BoolFlag{
			Name:  "a",
			Usage: "List all containers.",
		},
		cli.BoolFlag{
			Name:  "f",
			Usage: "Get specified container's flow.",
		},
		cli.BoolFlag{
			Name:  "c",
			Usage: "Create container.",
		},
		cli.BoolFlag{
			Name:  "u",
			Usage: "Update container.",
		},
		cli.BoolFlag{
			Name:  "r",
			Usage: "Restart container.",
		},
		cli.BoolFlag{
			Name:  "t",
			Usage: "Tag container.",
		},
		cli.BoolFlag{
			Name:  "d",
			Usage: "Delete container.",
		},
	}

	clusterFlags = []cli.Flag{
		cli.BoolFlag{
			Name:  "a",
			Usage: "List all clusters.",
		},
	}

	repositoryFlags = []cli.Flag{
		cli.BoolFlag{
			Name:  "a",
			Usage: "List all repositories.",
		},
	}

	secretkeyFlags = []cli.Flag{
		cli.BoolFlag{
			Name:  "a",
			Usage: "List all secret keys.",
		},
		cli.BoolFlag{
			Name:  "c",
			Usage: "create a secret key with arg key name.",
		},
		cli.BoolFlag{
			Name:  "d",
			Usage: "delete the secret key with arg id.",
		},
	}
)
