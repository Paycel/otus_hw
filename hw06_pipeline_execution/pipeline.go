package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	for _, stage := range stages {
		in = run(stage(in), done)
	}
	return in
}

func run(in, done In) Out {
	out := make(Bi)
	go func() {
		defer close(out)
		for {
			select {
			case <-done:
				return
			case res, ok := <-in:
				if !ok {
					return
				}
				out <- res
			}
		}
	}()
	return out
}
