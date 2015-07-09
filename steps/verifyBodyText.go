package steps

import (
	"fmt"
	"strings"
)

func init() {
	StepList["verifyBodyText"] = verifyBodyText
}

func verifyBodyText(a interface{}) {

	SCRIPT := SCRIPT_getElementsByXPath + `
        return document.body.textContent;
	`

	attr := a.(string)

	fmt.Print("[verifyBodyText]: " + attr)

	arg := []interface{}{}

	b, err := WD.ExecuteScript(SCRIPT, arg)
	StepFailure(err)

	body := b.(string)
	m := strings.Index(body, attr)

	if m != -1 {
		StepSuccess()
	} else {
		VerificationFailure()
	}

}
