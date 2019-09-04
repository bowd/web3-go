package ethbalance

import "fmt"

// RequestError wraps a request with the error received
type RequestError struct {
	req *BalanceRequest
	err error
}

func (re RequestError) section() string {
	return fmt.Sprintf("    %d %s %s", re.req.Block, re.req.Address, re.req.Source)
}

// CollectBalancesError wraps errors returned from the RPC requests
type CollectBalancesError struct {
	errors []*RequestError
}

// Error message aggregates all unique errors
func (cbe CollectBalancesError) Error() string {
	var fullErrorMessage string
	errorMessages := make(map[string]string)
	index := 0

	for _, reqError := range cbe.errors {
		errorMessageKey := reqError.err.Error()
		sectionError, ok := errorMessages[errorMessageKey]
		if !ok {
			sectionError = fmt.Sprintf("[%d] %s on:\n", index, errorMessageKey)
			index++
		}
		sectionError += reqError.section() + "\n"
		errorMessages[errorMessageKey] = sectionError
	}

	for _, errorMessage := range errorMessages {
		fullErrorMessage += errorMessage
	}

	return fmt.Sprintf("Unable to collect balances because of these errors: \n%s", fullErrorMessage)
}