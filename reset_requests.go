package main

import ()

type ResetRequest struct {
	Roll  string
	Email string
	Token string
}

func GetResetReq(roll string, email string, token string) ResetRequest {

	return ResetRequest{
		roll, email, token,
	}
}
