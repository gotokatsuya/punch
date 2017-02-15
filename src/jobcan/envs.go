package jobcan

import "os"

func loginURL() (s string) {
	if s = os.Getenv("JOBCAN_LOGIN_URL"); len(s) == 0 {
		panic("JOBCAN_LOGIN_URL is empty")
	}
	return
}

func loginEmail() (s string) {
	if s = os.Getenv("JOBCAN_LOGIN_EMAIL"); len(s) == 0 {
		panic("JOBCAN_LOGIN_EMAIL is empty")
	}
	return
}

func loginPassword() (s string) {
	if s = os.Getenv("JOBCAN_LOGIN_PASSWORD"); len(s) == 0 {
		panic("JOBCAN_LOGIN_PASSWORD is empty")
	}
	return
}
