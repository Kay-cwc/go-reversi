package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ValidatePromptResp[O uint | []uint] func(string) (O, string, bool)

func Ask[O uint | []uint](q string, validateFuncs []ValidatePromptResp[O]) O {
	var ans string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, q+":\n")
		ans, _ = r.ReadString('\n')
		if ans != "" {
			ans = strings.TrimSpace(ans)
			// should perform validation here
			for _, validateFunc := range validateFuncs {
				output, errorMsg, validationError := validateFunc(ans)
				if validationError {
					fmt.Println(errorMsg)
					break
				}
				return output
			}
			continue
		}
	}
}
