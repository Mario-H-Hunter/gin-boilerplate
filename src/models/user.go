package models

import "time"

type User struct {
	ID              string    `json:"id"`
	Email           string    `json:"email"`
	Confirmed       bool      `json:"confirmed"`
	Fullname        string    `json:"fullname"`
	InvitedAccount  bool      `json:"invitedAccount"`
	TrialExpiration time.Time `json:"trialExpiration"`
	Phone           string    `json:"phone"`
	OptInLocation   string    `json:"optInLocation"`
	Company         string    `json:"company"`
	JobTitle        string    `json:"jobTitle"`
	Credit          int       `json:"credit"`
	LastLoginAt     time.Time `json:"lastLoginAt"`
	Organization    string    `json:"organization"`
	CostCenter      string    `json:"cost_center"`
}

type SignInBackendResponse struct {
	Status  int    `json:"status"`
	Token   string `json:"token"`
	User    *User  `json:"user"`
	Crypted string `json:"crypted"`
}

type UserSignInResponse struct {
	Token string `json:"token" binding:"required"`
	User  User   `json:"user" binding:"required"`
}
