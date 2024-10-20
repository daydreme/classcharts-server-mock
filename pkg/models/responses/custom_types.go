package responses

type YesNoBool string

const (
	Yes YesNoBool = "yes"
	No  YesNoBool = "no"
)

func (ynb YesNoBool) AsBool() bool {
	return ynb == Yes
}
