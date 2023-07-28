package network

import "net"

type Args struct {
	N, M int
}

func (t *Args) Multiply(args *Args, reply *int) net.Error {
	*reply = args.N * args.M
	return nil
}
