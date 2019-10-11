package main

type CustomErr struct {
	Message string
}

func (e CustomErr) Error() string {
	return e.Message
}
