package erratum

import "fmt"

const testVersion = 2

// Use does something
func Use(o ResourceOpener, input string) (err error) {

	var r Resource

outer:
	for {
		r, err = o()

		// if no error break out of loop
		// if frob error handle it and break out of loop
		// if error is transient go round loop
		switch err.(type) {
		case nil:
			fmt.Println("no error found")
			break outer

		case TransientError:
			fmt.Println("transient error found")

		default:
			fmt.Println(err)
			return err
		}
	}

	defer r.Close()

	defer func(r Resource) {
		if p := recover(); p != nil {

			if fe, ok := p.(FrobError); ok {
				fmt.Println("recovered a frob error", p, fe.defrobTag)
				r.Defrob(fe.defrobTag)
				err = fe
			} else {
				fmt.Println("recovered a normal error", p)
			}
		}
	}(r)

	r.Frob(input)

	return err
}
