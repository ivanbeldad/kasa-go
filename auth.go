package kasa

import (
	uuid "github.com/gofrs/uuid"
)

type auth struct {
	Username string
	Password string
	URL      string
	Token    string
	UUID     string
}

func (a *auth) generateToken() error {
	req, err := a.getRequest()
	if err != nil {
		return err
	}
	res, err := req.execute()
	if err != nil {
		return err
	}
	a.Token = res.Token
	return err
}

func (a *auth) generateUUID() error {
	if a.UUID != "" {
		return nil
	}
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	a.UUID = uuid.String()
	return nil
}

func (a *auth) getRequest() (request, error) {
	var r request
	var err error
	err = a.generateUUID()
	if err != nil {
		return r, err
	}
	return request{
		URL: a.URL,
		RequestBody: requestBody{
			methodLogin,
			params{
				AppType:       appType,
				CloudUserName: a.Username,
				CloudPassword: a.Password,
				TerminalUUID:  a.UUID,
			},
		},
	}, nil
}
