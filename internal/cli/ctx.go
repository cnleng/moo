package cli

import "context"

type Ctx interface {
	context.Context
	Args() []string
	Flag(string) Flag
	Bool(string) bool
	Int(string) int
	Uint(string) uint
	Float(string) float64
	String(string) string
	StringSlice(string) []string
	IntSlice(string) []int
	UintSlice(string) []uint
	FloatSlice(string) []float64
	StringMap(string) map[string]string
	Map(string) map[string]interface{}
}

func newCtx(c context.Context, cmd *Cmd, args []string) Ctx {
	return &ctx{cmd: cmd, args: args, Context: c}
}

type ctx struct {
	context.Context
	cmd  *Cmd
	args []string
}

func (c *ctx) Args() []string {
	return c.args
}

func (c *ctx) Flag(key string) Flag {
	for _, v := range c.cmd.flags {
		if key == v.Key() {
			return v
		}
	}
	return nil
}

// some helper methods

func (c *ctx) Bool(key string) bool {
	return c.Flag(key).Var().(bool)
}

func (c *ctx) Int(key string) int {
	return c.Flag(key).Var().(int)
}

func (c *ctx) Uint(key string) uint {
	return c.Flag(key).Var().(uint)
}

func (c *ctx) Float(key string) float64 {
	return c.Flag(key).Var().(float64)
}

func (c *ctx) String(key string) string {
	return c.Flag(key).Var().(string)
}

func (c *ctx) StringSlice(key string) []string {
	return c.Flag(key).Var().([]string)
}

func (c *ctx) IntSlice(key string) []int {
	return c.Flag(key).Var().([]int)
}

func (c *ctx) UintSlice(key string) []uint {
	return c.Flag(key).Var().([]uint)
}

func (c *ctx) FloatSlice(key string) []float64 {
	return c.Flag(key).Var().([]float64)
}

func (c *ctx) StringMap(key string) map[string]string {
	return c.Flag(key).Var().(map[string]string)
}

func (c *ctx) Map(key string) map[string]any {
	return c.Flag(key).Var().(map[string]any)
}
