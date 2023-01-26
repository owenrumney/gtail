package auth

import (
	"errors"
	"fmt"
	"os"

	"github.com/googleapis/gax-go/v2/apierror"
)

func CheckErrorForAuth(err error) {
	var apiErr *apierror.APIError
	if errors.As(err, &apiErr) {
		if apiErr.GRPCStatus().Code().String() == "Unauthenticated" {
			fmt.Print("\nYou do not have active ADC credentials. Please run `gcloud auth application-default login` to set up ADC credentials.\n")
			os.Exit(1)
		}
	}
}
