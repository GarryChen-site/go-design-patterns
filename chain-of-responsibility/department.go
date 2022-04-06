package chainofresponsibility

//  Handler interface

type department interface {
	execute(* patient)

	setNext(department)
}
