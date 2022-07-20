package common

type EvalDetail struct {
	Variation interface{}
	Id        int64
	Reason    string
	Name      string
	KeyName   string
}

func (e *EvalDetail) IsSuccess() bool {
	return e.Id > 0
}
