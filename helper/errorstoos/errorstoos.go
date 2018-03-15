package errorstoos

import (
	"errors"
	"study-go/helper/lang"
)

func New(str string) error {
	locale := lang.Init()
	return errors.New(locale.Tr(str))
}
