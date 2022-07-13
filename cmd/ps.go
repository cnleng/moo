package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/moobu/moo/client"
	"github.com/moobu/moo/client/http"
	"github.com/moobu/moo/internal/cli"
	"github.com/moobu/moo/runtime"
)

func init() {
	cmd.Register(&cli.Cmd{
		Name:  "ps",
		About: "list running pods",
		Run:   List,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "ns",
				Usage: "filter by namespace",
				Value: "moo",
			},
			&cli.BoolFlag{
				Name:  "all",
				Usage: "list all",
			},
			&cli.StringFlag{
				Name:  "name",
				Usage: "specify the pod's name",
			},
			&cli.StringFlag{
				Name:  "tag",
				Usage: "specify the pod's tag",
			},
			&cli.StringFlag{
				Name:  "server",
				Usage: "address of Moo server",
				Value: defaultServerAddr,
			},
		},
	})
}

func List(c cli.Ctx) error {
	cli := http.New(client.Server(c.String("server")))
	pods, err := cli.List(
		runtime.Name(c.String("name")),
		runtime.Tag(c.String("tag")))
	if err != nil {
		return err
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 4, 1, '\t', tabwriter.AlignRight)
	fmt.Fprint(tw, "NAME\tTAG\tSTATUS\tSOURCE")

	for _, pod := range pods {
		meta := pod.Metadata
		fmt.Fprintf(tw, "%s\t%s\t%s\t%s", pod.Name, pod.Tag, meta["status"], meta["source"])
	}
	return tw.Flush()
}
