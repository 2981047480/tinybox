package exception

import "fmt"

type Exception struct {
	code  int
	excep string
}

func (e Exception) Error() string {
	return fmt.Sprintf("Error %v, Error Code %v", e.excep, e.code)
}

var ProcessNotFoundError Exception = Exception{code: 40001, excep: "Error Process Not Found"}
