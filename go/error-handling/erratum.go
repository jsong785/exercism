package erratum

func Use(o ResourceOpener, input string) (err error) {
	resource, err := OpenResource(o)
	if err != nil {
		return err
	}

	defer resource.Close()
	defer HandleFrobError(resource, &err)

	resource.Frob(input)
	return nil
}

func OpenResource(o ResourceOpener) (Resource, error) {
	resource, err := o()
	for {
		if _, ok := err.(TransientError); ok {
			resource, err = o()
			continue
		}
		break
	}
	return resource, err
}

func HandleFrobError(resource Resource, err *error) {
	if r := recover(); r != nil {
		switch e := r.(type) {
		case FrobError:
			resource.Defrob(e.defrobTag)
			*err = e.inner
		case error:
			*err = e
		}
	}
}
