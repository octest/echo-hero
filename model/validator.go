package model

const (
	MAX_VALIDATIONS = 10
	VALID_MISSING   = "%s is missing"
	VALID_MIN       = "%s may have min %d characters"
	VALID_MAX       = "%s may have max %d characters"
	VALID_MINMAX    = "%s may have between %d and %d characters"
)

type Validator interface {
	Valid() bool
}

type Validators struct {
	Messages []string
}

func NewValidators() (val *Validators) {
	val = &Validators{}
	val.Messages = make([]string, 0, MAX_VALIDATIONS)
	return
}

/*func (v *Validator) Validate(i interface{}) (val bool) {
	return
}*/
