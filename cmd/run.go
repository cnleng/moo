package cmd

import (
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"

	"github.com/moobu/moo/builder"
	"github.com/moobu/moo/client"
	"github.com/moobu/moo/client/http"
	"github.com/moobu/moo/internal/cli"
	"github.com/moobu/moo/runtime"
)

func init() {
	cmd.Register(&cli.Cmd{
		Name:  "run",
		About: "run a pod",
		Pos:   []string{"source"},
		Run:   Run,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "name",
				Usage: "name of the deploy",
			},
			&cli.StringFlag{
				Name:  "ref",
				Usage: "reference of the source to run",
			},
			&cli.IntFlag{
				Name:  "replicas",
				Usage: "replicas to be deployed",
				Value: 1,
			},
			&cli.BoolFlag{
				Name:  "gpu",
				Usage: "enable GPU support",
			},
			&cli.BoolFlag{
				Name:  "output",
				Usage: "enable the standard output",
				Value: true,
			},
			&cli.StringSliceFlag{
				Name:  "env",
				Usage: "environment variables to run with",
			},
			&cli.StringFlag{
				Name:  "server",
				Usage: "address of Moo server",
				Value: defaultServerAddr,
			},
			&cli.StringFlag{
				Name:  "image",
				Usage: "specify an image to run",
			},
		},
	})
}

// TODO: how can we watch the log output by the server end?
func Run(c cli.Ctx) error {
	// connect to the Moo server
	cli := http.New(client.Server(c.String("server")))
	// parse the given remote address of the source.
	rawURL := c.Pos()[0]
	u, err := url.Parse(rawURL)
	if err != nil {
		return err
	}
	// use the base name of the source if no name is given.
	name := c.String("name")
	if len(name) == 0 {
		name = filepath.Base(u.Path)
	}
	// use the reference name of the source for tagging the
	// pod if a reference is specified.
	tag := "latest"
	ref := c.String("ref")
	if len(ref) > 0 {
		tag = ref
	}
	source := &builder.Source{
		Name:   name,
		Remote: rawURL,
	}
	// tell the server to build the source, so we can use the
	// runtime later to run the returned bundle.
	bundle, err := cli.Build(source, builder.Ref(ref))
	if err != nil {
		return err
	}
	// specify the output
	// TODO: use a default file on the machine running the CLI.
	output := io.Discard
	if c.Bool("output") {
		output = os.Stdout
	}
	pod := &runtime.Pod{
		Name: name,
		Tag:  tag,
	}
	// build the optional functions used by the Moo runtime.
	// TODO: spscify the user's namespace to use
	opts := []runtime.CreateOption{
		runtime.CreateWithNamespace("default"),
		runtime.Env(c.StringSlice("env")...),
		runtime.Image(c.String("image")),
		runtime.Replicas(c.Int("replicas")),
		runtime.GPU(c.Bool("gpu")),
		runtime.Output(output),
		runtime.Bundle(bundle),
	}
	// tell the server to run a pod containing the bundle.
	if err := cli.Create(pod, opts...); err != nil {
		return err
	}
	fmt.Printf("Deployed.\n")
	return nil
}
