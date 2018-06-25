package courier

func Run(router *Router, transports ...Transport) {
	errs := make(chan error)

	for i := range transports {
		s := transports[i]
		go func() {
			if err := s.Serve(router); err != nil {
				errs <- err
			}
		}()
	}

	select {
	case err := <-errs:
		panic(err)
	}
}
