package validator

import (
	"fmt"
	"net/http"

	jsonio "github.com/GitCollabCode/GitCollab/internal/jsonhttp"
	"github.com/GitCollabCode/GitCollab/internal/models"
	"github.com/sirupsen/logrus"
)

// Helper that decodes and validates incoming JSON request body
func (v *Validation) GetJSON(structure interface{}, w http.ResponseWriter, r *http.Request, log *logrus.Logger) error {
	err := jsonio.FromJSON(structure, r.Body)
	if err != nil {
		log.Errorf("GetJSON failed to decode JSON request into struct: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		err1 := jsonio.ToJSON(&models.ErrorMessage{Message: "bad request"}, w)
		if err != nil {
			log.Fatalf("GetJSON failed to send error response: %s", err1.Error())
			return err1
		}
		return err
	}

	errs := v.Validate(structure)
	if len(errs) != 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		err = jsonio.ToJSON(&ValidationErrorResp{Messages: errs.Errors()}, w)
		if err != nil {
			log.Fatalf("GetJSON failed to send error response: %s", err.Error())
		}
		return fmt.Errorf("GetJSON validation errors found in struct")
	}

	return nil
}
