package models

type RequestLogin struct {
	Login string `json:"login,omitempty"`
	Password string `json:"password"`
}
