package main

import "context"

func reqTask(ctx context.Context, name string)  {
	for {
		select {
		case <- ctx.Done():

		}

	}
}
