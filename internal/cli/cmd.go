package cli

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"text/tabwriter"
	"unsafe"
)

type Intercept func(Ctx) error

type Cmd struct {
	Name    string
	Help    string
	Version string
	Args    []string
	Flags   []Flag
	cmds    []*Cmd
	Run     func(Ctx) error

	set *flag.FlagSet
	// root   *Cmd
	parent string

	Intercept Intercept
}

func (c *Cmd) Init() {
	c.init("", c, c.Flags)
}

func (c *Cmd) RunCtx(ctx context.Context) error {
	args := os.Args[1:]
	i, help := seperate(args)
	// see if we need to strip out the position argument
	j := i
	if !help {
		j -= len(c.Args) + 1
	}

	// find command
	cmd := c.find(args[:j], 0)
	if cmd == nil {
		return errors.New("no such command")
	}

	if cmd.Run == nil {
		return cmd.help(os.Stdout)
	}

	// parse flags
	if i < len(args) {
		if err := cmd.set.Parse(args[i:]); err != nil {
			return err
		}
	}

	// validate flags
	if err := cmd.validate(); err != nil {
		return err
	}

	// build context
	context := newCtx(ctx, cmd, args[j:i])

	// run interceptor
	if c.Intercept != nil {
		if err := c.Intercept(context); err != nil {
			return err
		}
	}

	// execute action
	return cmd.Run(context)
}

func (c *Cmd) Register(cmd *Cmd) {
	c.cmds = append(c.cmds, cmd)
}

func (c *Cmd) find(names []string, i int) *Cmd {
	if len(names) == i {
		return c
	}
	for _, cmd := range c.cmds {
		if cmd.Name == names[i] {
			if c := cmd.find(names, i+1); c != nil {
				return c
			}
		}
	}
	return nil
}

func (c *Cmd) validate() error {
	for _, flag := range c.Flags {
		if flag.Invalid() {
			return fmt.Errorf("option `%s` is required", flag.Key())
		}
	}
	return nil
}

func (c *Cmd) help(w io.Writer) error {
	tw := tabwriter.NewWriter(w, 0, 8, 1, '\t', tabwriter.AlignRight)

	if len(c.Help) != 0 {
		fmt.Fprintf(tw, "NAME:\n\t%s %s\n\n", c.Help, c.Version)
	}

	sep := ""
	if len(c.parent) > 0 {
		sep = " "
	}
	fmt.Fprintf(tw, "USAGE:\n\t%s%s%s", c.parent, sep, c.Name)

	if len(c.cmds) > 0 {
		fmt.Fprint(tw, " <command>")
	}

	for _, arg := range c.Args {
		fmt.Fprintf(tw, " <%s>", arg)
	}

	fmt.Fprint(tw, " [options...]\n")

	if c.cmds != nil {
		fmt.Fprint(tw, "\nCOMMANDS:\n")
		for _, cmd := range c.cmds {
			fmt.Fprintf(tw, "\t%s\t%s\n", cmd.Name, cmd.Help)
		}
	}

	if c.Flags != nil {
		fmt.Fprint(tw, "\nOPTIONS:\n")
		for _, flag := range c.Flags {
			fmt.Fprintf(tw, "\t--%s\t%s\t(default %v) \n", flag.Key(), flag.Help(), flag.Var())
		}
	}

	return tw.Flush()
}

func (c *Cmd) init(parent string, root *Cmd, globalFlags []Flag) {
	set := flag.NewFlagSet(c.Name, flag.ExitOnError)

	if c.Flags != nil {
		sort.Slice(c.Flags, func(i, j int) bool {
			return c.Flags[i].Key() < c.Flags[j].Key()
		})

		for _, flag := range c.Flags {
			switch f := flag.(type) {
			case *BoolFlag:
				set.BoolVar(&f.Value, f.Name, f.Value, f.Usage)
			case *IntFlag:
				set.IntVar(&f.Value, f.Name, f.Value, f.Usage)
			case *UintFlag:
				set.UintVar(&f.Value, f.Name, f.Value, f.Usage)
			case *FloatFlag:
				set.Float64Var(&f.Value, f.Name, f.Value, f.Usage)
			case *StringFlag:
				set.StringVar(&f.Value, f.Name, f.Value, f.Usage)
			case *StringSliceFlag:
				set.Var((*StringSlice)(unsafe.Pointer(&f.Value)), f.Name, f.Usage)
			case *IntSliceFlag:
				set.Var((*IntSlice)(unsafe.Pointer(&f.Value)), f.Name, f.Usage)
			case *UintSliceFlag:
				set.Var((*UintSlice)(unsafe.Pointer(&f.Value)), f.Name, f.Usage)
			case *FloatSliceFlag:
				set.Var((*FloatSlice)(unsafe.Pointer(&f.Value)), f.Name, f.Usage)
			case *StringMapFlag:
				f.Value = StringMap{}
				set.Var((*StringMap)(unsafe.Pointer(&f.Value)), f.Name, f.Usage)
			case *MapFlag:
				f.Value = Map{}
				set.Var((*Map)(unsafe.Pointer(&f.Value)), f.Name, f.Usage)
			default:
				fmt.Printf("unsupported flag %s, ignored\n", f)
			}
		}

	}

	set.Usage = func() {
		c.help(os.Stdout)
	}
	c.set = set
	c.parent = parent

	for _, cmd := range c.cmds {
		cmd.Flags = append(cmd.Flags, globalFlags...)
		cmd.init(fmt.Sprintf("%s%s", parent, c.Name), root, globalFlags)
	}
}

func seperate(args []string) (int, bool) {
	var help bool
	for i, arg := range args {
		if arg[0] == '-' {
			switch arg[1] {
			case '-':
				help = arg[2:] == "help"
			case 'h':
				help = true
			}
			return i, help
		}
	}
	return len(args), help
}
