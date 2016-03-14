package gilmour

type PipeComposition struct {
	composition
}

func (c *PipeComposition) Execute(m *Message) (resp *Response, err error) {
	do := func(do recfunc, m *Message, f *Response) {
		cmd := c.lpop()
		resp, err = performJob(cmd, m)

		if len(c.executables()) > 0 && resp.Code() == 200 && err == nil {
			resp = inflateResponse(resp)
			do(do, resp.Next(), f)
			return
		}
	}

	do(do, copyMessage(m), resp)
	return
}

//New Pipe composition
func (g *Gilmour) NewPipe(cmds ...Executable) *PipeComposition {
	c := new(PipeComposition)
	c.setEngine(g)
	c.add(cmds...)
	return c
}
