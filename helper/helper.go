package helper

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jordan-wright/email"
	"im/define"
	"net/smtp"
)

var myKey = []byte("im")

type UserClaims struct {
	Identity string
	//Identity primitive.ObjectID
	Email string
	jwt.StandardClaims
}

func Md5(s string) string {
	hash := md5.New()
	return fmt.Sprintf("%x", hash.Sum([]byte(s)))
}

func GenerateToken(identity, email string) (string, error) {
	claim := &UserClaims{
		Identity:       identity,
		Email:          email,
		StandardClaims: jwt.StandardClaims{}, // 暂时未做刷新处理
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func AnalyseToken(tokenString string) (*UserClaims, error) {
	claims, err := jwt.ParseWithClaims(tokenString, new(UserClaims), func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claim, ok := claims.Claims.(*UserClaims); ok && claims.Valid {
		return claim, nil
	}
	return nil, errors.New("token错误")
}

func MailSendCode(toUserEmail string, code string) error {
	e := email.NewEmail()
	e.From = "Get <949244762@qq.com>"
	e.To = []string{toUserEmail}
	e.Subject = "验证码已发送，请查收"
	e.HTML = []byte("您的验证码：<b>" + code + "</b>")
	//e.SendWithTLS("smtp.163.com:465",
	//	smtp.PlainAuth("", "getcharzhaopan@163.com", define.MailPassword, "smtp.163.com"),
	//	&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	err := e.Send("smtp.qq.com:587", smtp.PlainAuth("", "949244762@qq.com", define.MailPassword, "smtp.qq.com"))
	if err != nil {
		return err
	}

	return nil
}
