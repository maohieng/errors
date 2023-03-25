package errs

type Stack struct {
	Msg  string `json:"msg,omitempty"`
	Op   Op     `json:"op,omitempty"`
	Kind Kind   `json:"kind,omitempty"`
	//Severe string `json:"severe,omitempty"`
	Err *Stack `json:"err,omitempty"`
}

func Errors(err error) *Stack {
	if err == nil {
		return nil
	}

	e, ok := err.(*Error)
	if !ok {
		return &Stack{
			Msg:  err.Error(),
			Op:   "",
			Kind: 0,
			//Severe: "",
			Err: nil,
		}
	}

	return &Stack{
		Msg:  e.msg,
		Op:   e.op,
		Kind: e.kind,
		//Severe: e.sev.String(),
		Err: Errors(e.err),
	}
}
