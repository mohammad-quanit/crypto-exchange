package programming

import (
	"strings"

	"github.com/google/uuid"
)

func (pf *ProgrammingFunctions) NewUuid(withoutHyphen bool) string {
	uuidWithHyphen := uuid.New()

	if withoutHyphen {
		return strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	}

	return uuidWithHyphen.String()
}
