package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ValidatePromptResp[O uint | [2]uint] func(string) (O, string, bool)

func Ask[O uint | [2]uint](q string, validateFunc ValidatePromptResp[O]) O {
	var ans string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, q+":\n")
		ans, _ = r.ReadString('\n')
		if ans != "" {
			ans = strings.TrimSpace(ans)
			// should perform validation here
			output, errorMsg, validationError := validateFunc(ans)
			if validationError {
				fmt.Println(errorMsg)
				continue
			}
			return output
		}
	}
}
