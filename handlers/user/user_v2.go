package user

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nju-iot/edgex_admin/caller"
	"github.com/nju-iot/edgex_admin/dal"
	"github.com/nju-iot/edgex_admin/logs"
	"github.com/nju-iot/edgex_admin/resp"
	"gopkg.in/gomail.v2"
)

// RegisterParamsV2 ...
type RegisterParamsV2 struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Email    string `form:"email" json:"email"`
}

// RegisterCheckParamsV2 ...
type RegisterCheckParamsV2 struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Email    string `form:"email" json:"email"`
	Code     string `form:"code" json:"code" binding:"required"`
}

// CodeCheckerV2 ...
type CodeCheckerV2 struct {
	Email string
	Code  string
}

var checker map[string]CodeCheckerV2

// RegisterV2 ...
func RegisterV2(c *gin.Context) *resp.JSONOutput {
	// Step1. 参数校验
	params := &RegisterParamsV2{}
	err := c.Bind(&params)
	if err != nil {
		logs.Error("[Register] request-params error: params=%+v, err=%v", params, err)
		return resp.SampleJSON(c, resp.RespCodeParamsError, nil)
	}

	// Step2. 查看用户/邮箱是否存在
	userInfo, dbErr := dal.GetEdgexUserByName(params.Username)
	mailInfo, dbErr2 := dal.GetEdgexUserByEmail(params.Email)
	if dbErr != nil {
		logs.Error("[Register] get user failed: username=%s, err=%v", params.Username, dbErr)
		return resp.SampleJSON(c, resp.RespDatabaseError, nil)
	}
	if dbErr2 != nil {
		logs.Error("[Register] email [%s] already exists: err=%v", params.Email, dbErr2)
	}
	if userInfo != nil && mailInfo != nil {
		return resp.SampleJSON(c, resp.RespCodeUserExsit, nil)
	}

	//Step3. 发送邮箱验证码
	mailTo := []string{params.Email}
	subject := string("登录验证")
	code := randomCode()
	body := code
	err = SendMailToV2(mailTo, subject, body)
	if err != nil {
		return resp.SampleJSON(c, resp.RespCodeParamsError, "发送失败")
	}

	if checker == nil {
		checker = make(map[string]CodeCheckerV2)
	}
	var cc CodeCheckerV2
	cc.Email = params.Email
	cc.Code = code
	checker[params.Email] = cc
	return resp.SampleJSON(c, resp.RespCodeSuccess, nil)
}

// RegisterCheckV2 ...
func RegisterCheckV2(c *gin.Context) *resp.JSONOutput {
	params := &RegisterCheckParamsV2{}
	err := c.Bind(&params)
	if err != nil {
		logs.Error("[RegisterCheck] request-params error: params=%+v, err=%v", params, err)
		return resp.SampleJSON(c, resp.RespCodeParamsError, nil)
	}
	//验证验证码
	err = checkCodeV2(params.Email, params.Code)
	if err != nil {
		logs.Error("[RegisterCheck] check-code error: params=%+v, err=%v", params, err)
		return resp.SampleJSON(c, resp.RespCodeParamsError, nil)
	}
	//添加用户
	user := &dal.EdgexUser{
		Username:     params.Username,
		Password:     params.Password,
		Email:        params.Email,
		CreatedTime:  time.Now(),
		ModifiedTime: time.Now(),
	}
	dbErr := dal.AddEdgexUser(caller.EdgexDB, user)
	if dbErr != nil {
		return resp.SampleJSON(c, resp.RespDatabaseError, nil)
	}
	return resp.SampleJSON(c, resp.RespCodeSuccess, nil)
}

type MailParam struct {
	Email string `form:"email" json:"email" binding:"required"`
}

func SendMailToV2(mailTo []string, subject string, body string) error {
	//定义邮箱服务器连接信息，如果是网易邮箱 pass填密码，qq邮箱填授权码

	//mailConn := map[string]string{
	//  "user": "xxx@163.com",
	//  "pass": "your password",
	//  "host": "smtp.163.com",
	//  "port": "465",
	//}

	mailConn := map[string]string{
		"user": "2369351080@qq.com",
		"pass": "inkdesahnqrjdjeg",
		"host": "smtp.qq.com",
		"port": "465",
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()

	m.SetHeader("From", m.FormatAddress(mailConn["user"], "NJU-IOT-EDGEX验证")) //这种方式可以添加别名，即“XX官方”　 //说明：如果是用网易邮箱账号发送，以下方法别名可以是中文，如果是qq企业邮箱，以下方法用中文别名，会报错，需要用上面此方法转码
	//m.SetHeader("From", "FB Sample"+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“FB Sample”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	//m.SetHeader("From", mailConn["user"])
	m.SetHeader("To", mailTo...)    //发送给多个用户
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/html", body)    //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	return err

}

func SendMailV2(c *gin.Context) *resp.JSONOutput {
	params := &MailParam{}
	err := c.Bind(&params)
	if err != nil {
		logs.Error("[SendMail] request-params error: params=%+v, err=%v", params, err)
		return resp.SampleJSON(c, resp.RespCodeParamsError, nil)
	}
	mailTo := []string{params.Email}
	subject := string("登录验证")
	code := randomCode()
	body := code
	err = SendMailToV2(mailTo, subject, body)
	if err != nil {
		return resp.SampleJSON(c, resp.RespCodeParamsError, "发送失败")
	}
	return resp.SampleJSON(c, resp.RespCodeSuccess, nil)
}

func randomCode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return vcode
}

func checkCodeV2(email string, code string) error {
	if checker == nil {
		checker = make(map[string]CodeCheckerV2)
		return errors.New("no email found")
	}
	c, ok := checker[email]
	if !ok {
		return errors.New("no email found")
	}
	if c.Code == code {
		return nil
	}
	return errors.New("wrong code")
}

// func setEntryptedQuestions(questionCode int, user_id int64, answer string) {

// }

// func updateEntrypted(user_name string, entrypted string) error {
// 	var fieldsMap map[string]interface{} = map[string]interface{}{"entrypted": entrypted}
// 	err := dal.UpdateEdgexUser(user_name, fieldsMap)
// 	return err
// }

func updatePasswordV2(userID int64, password string) error {
	var fieldsMap map[string]interface{} = map[string]interface{}{"password": password}
	err := dal.UpdateEdgexUser(userID, fieldsMap)
	return err
}

// type EntryptedParams struct {
// 	UserName   string `form:"user_name" json:"user_name"`
// 	QuestionId int    `form:"question_id" json:"question_id"`
// 	Answer     string `form:"answer" json:"answer"`
// }

type PasswordParams struct {
	UserID   int64  `form:"user_id" json:"user_id"`
	Password string `form:"password" json:"password"`
}

// func UpdateUserEntrypted(c *gin.Context) *resp.JSONOutput {
// 	params := &EntryptedParams{}
// 	err := c.Bind(&params)
// 	if err != nil {
// 		logs.Error("[UpdateEntrypted] request-params error: params=%+v, err=%v", params, err)
// 		return resp.SampleJSON(c, resp.RespCodeParamsError, nil)
// 	}
// 	entrypted := fmt.Sprintf("%d %s", params.QuestionId, params.Answer)
// 	err = updateEntrypted(params.UserName, entrypted)
// 	if err != nil {
// 		return resp.SampleJSON(c, resp.RespCodeParamsError, "更新失败")
// 	}
// 	return resp.SampleJSON(c, resp.RespCodeSuccess, nil)
// }

func UpdateUserPasswordV2(c *gin.Context) *resp.JSONOutput {
	params := &PasswordParams{}
	err := c.Bind(&params)
	if err != nil {
		logs.Error("[UpdatePassword] request-params error: params=%+v, err=%v", params, err)
		return resp.SampleJSON(c, resp.RespCodeParamsError, nil)
	}
	password := params.Password
	if updatePasswordV2(params.UserID, password) != nil {
		return resp.SampleJSON(c, resp.RespCodeParamsError, "更新失败")
	}
	return resp.SampleJSON(c, resp.RespCodeSuccess, nil)
}

// func CheckUserEntrypted(c *gin.Context) *resp.JSONOutput {
// 	params := &EntryptedParams{}
// 	err := c.Bind(&params)
// 	if err != nil {
// 		logs.Error("[CheckEntrypted] request-params error: params=%+v, err=%v", params, err)
// 		return resp.SampleJSON(c, resp.RespCodeParamsError, nil)
// 	}

// 	user_name := params.UserName
// 	var (
// 		userInfo *dal.EdgexUser
// 		dbErr    error
// 	)
// 	userInfo, dbErr = dal.GetEdgexUserByName(user_name)
// 	if dbErr != nil {
// 		logs.Error("[CheckEntrypted] get userInfo failed: params=%+v, err=%v", params, err)
// 		return resp.SampleJSON(c, resp.RespDatabaseError, nil)
// 	}
// 	if userInfo == nil {
// 		logs.Error("[CheckEntrypted] user is Not Exsit: params=%+v, userInfo=%+v", params, userInfo)
// 		return resp.SampleJSON(c, resp.RespCodeParamsError, nil)
// 	}
// 	if userInfo.Entrypted == "" {
// 		logs.Error("[CheckEntrypted] entrypted Not Exist: params=%+v, userInfo=%+v", params, userInfo)
// 		return resp.SampleJSON(c, resp.RespDatabaseError, "密保不存在")
// 	}
// 	entrypted := fmt.Sprintf("%d %s", params.QuestionId, params.Answer)
// 	if userInfo.Entrypted != entrypted {
// 		logs.Error("[CheckEntrypted] entrypted error: params=%+v, userInfo=%+v", params, userInfo)
// 		return resp.SampleJSON(c, resp.RespDatabaseError, "密保错误")
// 	}

// 	return resp.SampleJSON(c, resp.RespCodeSuccess, nil)
// }
