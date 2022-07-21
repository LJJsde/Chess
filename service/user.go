package service

import (
	"crypto/md5"
	"fmt"
	"regexp"
	"Chess/module"
)

//先放着
func NewUser() *module.SUser {
	p := &module.SUser{}
	return p
}

func Password(password, salt string) string {
	p := md5.New()
	p.Write([]byte(salt + password))
	return fmt.Sprintf("%x", p.Sum(nil))
}

func VerifyEmail(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
