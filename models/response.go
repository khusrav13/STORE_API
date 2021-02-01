package models

type ResponseStatus struct {
	Ok      bool   `json:"ok,omitempty"`
	Message string `json:"message,omitempty"`
}

type ResponseToken struct {
	Ok    bool   `json:"ok,omitempty"`
	Token string `json:"token,omitempty"`
	User  ResponseUser
}

type StatusResponse struct {
	Ok      bool   `json:"ok,omitempty"`
	Message string `json:"message,omitempty"`
}

type UserRoleJWT struct {
	Role   int    `json:"role,omitempty"`
	Name   string `json:"name,omitempty"`
	UserID int    `json:"user_id,omitempty"`
}

type CredentialStatus struct {
	Ok       bool `json:"ok,omitempty"`
	Login    bool `json:"login,omitempty"`
	Password bool `json:"password,omitempty"`
}
