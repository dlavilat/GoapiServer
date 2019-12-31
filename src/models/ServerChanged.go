package models

type ServerChanged struct {
	serversChanged   bool
	sslGrade         string
	previousSslGrade string
	logo             string
	title            string
	isDown           bool
}
