package user

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
	"study-go/helper/crypto"
	"study-go/helper/errorstoos"
	"study-go/helper/timetools"
	"study-go/models"
)

var (
	RegistInfoError, MissingPassword, MissingAccount, MissingNickname, AccountAlreadyExists, AccountNotExist, IncorrectPassword, UserLocked error
)

func init() {
	RegistInfoError = errorstoos.New("RegistInfoError")
	MissingPassword = errorstoos.New("MissingPassword")
	MissingAccount = errorstoos.New("MissingAccount")
	MissingNickname = errorstoos.New("MissingNickname")
	AccountAlreadyExists = errorstoos.New("AccountAlreadyExists")
	AccountNotExist = errorstoos.New("AccountNotExist")
	IncorrectPassword = errorstoos.New("IncorrectPassword")
	UserLocked = errorstoos.New("UserLocked")
}

// 注册
func Regist(user map[string]string) (userId int64, err error) {
	o := orm.NewOrm()
	userInfo := new(models.User)
	userProfile := new(models.UserProfile)

	if userInfo == nil {
		return 0, RegistInfoError
	}

	if _, ok := user["password"]; ok == false {
		return 0, MissingPassword
	}

	if _, ok := user["account"]; ok == false {
		return 0, MissingAccount
	}

	if _, ok := user["nickname"]; ok == false {
		return 0, MissingNickname
	}

	if _, err := strconv.Atoi(user["account"]); nil == err {
		userInfo.Mobile = user["account"]
		o.Read(userInfo, "Mobile")
	} else {
		userInfo.Email = user["account"]
		o.Read(userInfo, "Email")
	}

	if userInfo.Id > 0 {
		return 0, AccountAlreadyExists
	}

	salt := crypto.RandInt64Str()
	password := crypto.Md5(user["password"] + salt)

	dataTime := timetools.NowDateTime()

	sql := "insert into " + userInfo.TableName() + "( `email`, `mobile`, `password`, `salt`,  `ip`, `create_time`, `update_time`) values(?, ?, ?, ?, ?, ?, ?)"

	o.Begin()
	//userId, err = o.Insert(userInfo)
	res, err := o.Raw(sql, userInfo.Email, userInfo.Mobile, password, salt, user["ip"], dataTime, dataTime).Exec()
	if nil != err {
		o.Rollback()
		return
	}

	userId, _ = res.LastInsertId()
	userProfile.Id = userId
	userProfile.Nickname = user["nickname"]

	_, err = o.Insert(userProfile)
	if nil != err {
		o.Rollback()
		return
	}

	o.Commit()
	return
}

func Login(account string, password string) (user *models.User, err error) {
	o := orm.NewOrm()
	userInfo := new(models.User)
	sql := "select `id`, `email`, `mobile`, `password`, `locked`, `salt` from " + userInfo.TableName() + " where "

	if _, err := strconv.Atoi(account); nil == err {
		sql += "mobile= ?"
	} else {
		sql += "email= ?"
	}

	//var maps []orm.Params
	//num, err := o.Raw(sql, account).Values(&maps)

	var userTmp models.UserLogin
	err = o.Raw(sql, account).QueryRow(&userTmp)

	if nil != err {
		err = AccountNotExist
		return
	}

	if userTmp.Locked == 1 {
		err = UserLocked
		return
	}

	//fmt.Println(crypto.Md5(password+userTmp.Salt))
	//fmt.Println(userTmp.Password)
	if crypto.Md5(password+userTmp.Salt) != userTmp.Password {
		err = IncorrectPassword
		return
	}
	userId := userTmp.Id
	user, err = models.GetUserById(userId)
	return
}

func JwtLogin(account string, password string) (jwt string, user *models.User, err error) {
	user, err = Login(account, password)
	if nil != err {
		return
	}
	return
}

func GetUserById(id int64) (*models.User, error) {
	user, err := models.GetUserById(id)
	return user, err
}

func GetAllUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	return models.GetAllUser(query, fields, sortby, order, offset, limit)
}
