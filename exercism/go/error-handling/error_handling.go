package erratum

func Use(ro ResourceOpener, s string) (err error) {
	var res Resource

	for res, err = ro(); err != nil; res, err = ro() {
		if _, ok := err.(TransientError); !ok {
			return err
		}
	}
	defer func() { _ = res.Close() }()

	defer func(){
		if r := recover(); r != nil {
			switch e := r.(type) {
			case FrobError:
				res.Defrob(e.defrobTag)
				err = e
			case error:
				err = e
			}
		}
	}()

	res.Frob(s)

	return err

}