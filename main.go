/*
 * @Author: your name
 * @Date: 2021-07-24 13:34:42
 * @LastEditTime: 2021-08-03 20:46:23
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \WebProject\main.go
 */
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Money    int    `json:"money"`
}

var CurrentUser User

type Response struct {
	Data string `json:"data"`
}
type Shop struct {
	ShopName string `json:"shopname"`
	Address  string `json:"address"`
}
type MilkTea struct {
	TeaName      string `json:"teaname"`
	BasicPrice   int    `json:"basicprice"`
	Profile      string `json:"profile"`      //简介
	Introduction string `json:"introduction"` //详细介绍
	Category     string `json:"category"`
}
type OrderForm struct {
	UserName    string `json:"username"`
	ShopName    string `json:"shopname"`
	TeaName     string `json:"teaname"`
	Sweet       string `json:"sweet"`
	Condition   string `json:"condition"`
	AddMaterial string `json:"addmaterial"`
	Number      int    `json:"number"`
	TotalPrice  int    `json:"totalprice"`
}

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/webapp")
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	http.HandleFunc("/register", Register)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/getuser", GetUserInformation)
	http.HandleFunc("/recharge", Recharge)
	http.HandleFunc("/altername", AlterName)
	http.HandleFunc("/newpassword", AlterPassword)
	http.HandleFunc("/getshops", GetShopsInformation)
	http.HandleFunc("/gettea", GetTeaInformation)
	http.HandleFunc("/purchase", Purchase)
	http.HandleFunc("/gethistory", GetHistoryOrders)

	http.ListenAndServe(":80", nil)
}

//需要用户名和密码
func GetUserInformation(w http.ResponseWriter, r *http.Request) {
	PreProcessService(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	} //返回数据格式是json
	if CurrentUser.ID == 0 {
		w.WriteHeader(403)
	} else {

		config := ProcessJson(r)
		if config["username"] == CurrentUser.Username && config["password"] == CurrentUser.Password {
			JsonResponse(CurrentUser, w)
			return
		} else {
			w.WriteHeader(403)
		}
	}

}
func Register(w http.ResponseWriter, r *http.Request) {
	PreProcessService(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	} //返回数据格式是json
	config := ProcessJson(r)
	CurrentUser = User{
		ID:       0,
		Username: "",
		Password: "",
		Money:    0,
	}
	user := CheckUsername(config["username"].(string))
	if user != nil {
		w.WriteHeader(403)
		return
	}
	AddUser(config["username"].(string), config["password"].(string))
	w.WriteHeader(200)
}
func Login(w http.ResponseWriter, r *http.Request) {

	PreProcessService(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	} //返回数据格式是json
	config := ProcessJson(r)
	user := CheckUsernameAndPassword(config["username"].(string), config["password"].(string))
	if user != nil {
		CurrentUser = *user
		w.WriteHeader(200)
	} else {
		CurrentUser = User{
			ID:       0,
			Username: "",
			Password: "",
			Money:    0,
		}
		w.WriteHeader(403)
	}

}

//充值
func Recharge(w http.ResponseWriter, r *http.Request) {
	PreProcessService(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	} //返回数据格式是json
	if CurrentUser.ID == 0 {
		w.WriteHeader(403)
	}
	config := ProcessJson(r)
	if config["username"] == CurrentUser.Username && config["password"] == CurrentUser.Password {
		temp := config["money"].(float64)
		CurrentUser.Money += int(temp)
		UpdateMoney(CurrentUser.Username, CurrentUser.Money)
		//fmt.Fprintln(w, "  ", reflect.TypeOf(temp))
		w.WriteHeader(200)
	} else {
		w.WriteHeader(403)
	}
}
func AlterName(w http.ResponseWriter, r *http.Request) {
	PreProcessService(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	} //返回数据格式是json
	if CurrentUser.ID == 0 {
		w.WriteHeader(403)
	}
	config := ProcessJson(r)
	if config["username"] == CurrentUser.Username && config["password"] == CurrentUser.Password {
		user := CheckUsername(config["newname"].(string))
		if user != nil {
			w.WriteHeader(403)
			return
		}
		UpdateName(config["newname"].(string), config["username"].(string))
		CurrentUser.Username = config["newname"].(string)
		w.WriteHeader(200)
	} else {
		w.WriteHeader(403)
	}
}
func AlterPassword(w http.ResponseWriter, r *http.Request) {
	PreProcessService(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	} //返回数据格式是json
	if CurrentUser.ID == 0 {
		w.WriteHeader(403)
	}
	config := ProcessJson(r)
	if config["username"] == CurrentUser.Username && config["password"] == CurrentUser.Password {
		UpdatePassword(config["newpassword"].(string), CurrentUser.Username)
		CurrentUser.Password = config["newpassword"].(string)
		w.WriteHeader(200)
	} else {
		w.WriteHeader(403)
	}
}
func GetShopsInformation(w http.ResponseWriter, r *http.Request) {
	PreProcessService(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	} //返回数据格式是json
	JsonResponse(GetShops(), w)
}
func GetTeaInformation(w http.ResponseWriter, r *http.Request) {
	PreProcessService(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	} //返回数据格式是json
	config := ProcessJson(r)
	str := config["category"].(string)
	JsonResponse(GetMilkTeaByCategory(str), w)
}
func Purchase(w http.ResponseWriter, r *http.Request) {
	PreProcessService(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	} //返回数据格式是json
	config := ProcessJson(r)
	if config["username"] == CurrentUser.Username && config["password"] == CurrentUser.Password {
		price := int(config["totalprice"].(float64))
		if price < CurrentUser.Money {
			CurrentUser.Money -= price
			UpdateMoney(CurrentUser.Username, CurrentUser.Money)
			AddOrderForm(config["shopname"].(string), config["teaname"].(string), config["sweet"].(string), config["condition"].(string), config["addmaterial"].(string), int(config["number"].(float64)), int(config["totalprice"].(float64)))

			w.WriteHeader(200)
		} else {
			fmt.Println("余额不足")
			w.WriteHeader(403)
		}

	} else {
		w.WriteHeader(403)
	}
}
func GetHistoryOrders(w http.ResponseWriter, r *http.Request) {
	PreProcessService(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	} //返回数据格式是json
	config := ProcessJson(r)
	if config["username"] == CurrentUser.Username && config["password"] == CurrentUser.Password {
		orders := GetHistoryOrderForm(CurrentUser.Username)
		JsonResponse(orders, w)
		w.WriteHeader(200)
	} else {
		w.WriteHeader(403)
	}
}
func ProcessJson(r *http.Request) map[string]interface{} {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var config map[string]interface{}
	json.Unmarshal(body, &config)
	return config
}
func PreProcessService(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")                                                            //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token") //header的类型
	w.Header().Add("Access-Control-Allow-Credentials", "true")                                                    //设置为true，允许ajax异步请求带cookie信息
	w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")                             //允许请求方法
	w.Header().Set("content-type", "application/json;charset=UTF-8")
}

func JsonResponse(response interface{}, w http.ResponseWriter) {

	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func AddUser(name, password string) {
	sqlStr := "insert into users(username,password,money) values(?,?,?)"
	inSmt, err := Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译异常:", err)
		panic(err.Error())
	}
	_, err1 := inSmt.Exec(name, password, 100)
	if err1 != nil {
		fmt.Println("执行异常:", err1)
		panic(err.Error())
	}
}
func DeleteUserByName(name string) error {
	sqlStr := "delete from users where username =?"
	inSmt, err := Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译异常:", err)
		return err
	}
	_, err1 := inSmt.Exec(name)
	if err1 != nil {
		fmt.Println("执行异常:", err1)
		return err1
	}
	return nil
}
func GetUserByName(name string) *User {
	sqlStr := "select id , username , password , money from users where username = ?"
	row := Db.QueryRow(sqlStr, name)
	var id int
	var username string
	var password string
	var money int
	err := row.Scan(&id, &username, &password, &money)
	if err != nil {
		fmt.Println("查询一位用户失败")
		panic(err.Error())
	}
	u := &User{
		ID:       id,
		Username: username,
		Password: password,
		Money:    money,
	}
	return u
}
func GetUsers() []*User {
	sqlStr := "select id , username, password, money from users"
	rows, err := Db.Query(sqlStr)
	if err != nil {
		fmt.Println("查询全部用户失败")
		panic(err.Error())
	}
	var users []*User
	for rows.Next() {
		var id int
		var username string
		var password string
		var money int
		err := rows.Scan(&id, &username, &password, &money)
		if err != nil {
			fmt.Println("查询失败")
			panic(err.Error())
		}
		u := &User{
			ID:       id,
			Username: username,
			Password: password,
			Money:    money,
		}
		users = append(users, u)
	}
	return users
}
func CheckUsernameAndPassword(name, password string) *User {
	sqlStr := " select id , username , password , money from users where username= ? and password = ? "
	row := Db.QueryRow(sqlStr, name, password)
	u := &User{}
	err := row.Scan(&u.ID, &u.Username, &u.Password, &u.Money)
	if err != nil {
		return nil
	} else {
		return u
	}
}
func CheckUsername(name string) *User {
	sqlStr := "select id,username,password,money from users where username= ? "
	row := Db.QueryRow(sqlStr, name)
	user := &User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Money)
	if user.ID == 0 {
		return nil
	} else {
		return user
	}

}
func UpdateMoney(name string, money int) {
	sqlStr := "update users set money = ? where username = ? "
	inSmt, err := Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译异常:", err)
		panic(err.Error())
	}
	_, err1 := inSmt.Exec(money, name)
	if err1 != nil {
		fmt.Println("执行异常:", err1)
		panic(err.Error())
	}
}
func UpdateName(newName, oldName string) {
	sqlStr := "update users set username = ? where username = ? "
	inSmt, err := Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译异常:", err)
		panic(err.Error())
	}
	_, err1 := inSmt.Exec(newName, oldName)
	if err1 != nil {
		fmt.Println("执行异常:", err1)
		panic(err.Error())
	}
}
func UpdatePassword(newPassword, name string) {
	sqlStr := "update users set password = ? where username = ? "
	inSmt, err := Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译异常:", err)
		panic(err.Error())
	}
	_, err1 := inSmt.Exec(newPassword, name)
	if err1 != nil {
		fmt.Println("执行异常:", err1)
		panic(err.Error())
	}
}
func GetShops() []Shop {
	sqlStr := "select  shopname, address from shop "
	rows, err := Db.Query(sqlStr)
	if err != nil {
		fmt.Println("初始化查询商店失败")
		panic(err.Error())
	}
	var shops []Shop
	for rows.Next() {
		var shop Shop
		err := rows.Scan(&shop.ShopName, &shop.Address)
		if err != nil {
			fmt.Println("查询店铺失败")
			panic(err.Error())
		}
		shops = append(shops, shop)
	}
	return shops
}
func GetMilkTeaByCategory(category string) []MilkTea {
	sqlStr := "select teaname,basicprice,profile,introduction  from milktea where category = ? "
	rows, err := Db.Query(sqlStr, category)
	if err != nil {
		fmt.Println("初始化查询奶茶失败")
		panic(err.Error())
	}
	var milks []MilkTea
	for rows.Next() {
		var milk MilkTea
		milk.Category = category
		err = rows.Scan(&milk.TeaName, &milk.BasicPrice, &milk.Profile, &milk.Introduction)
		if err != nil {
			fmt.Println("查询一个奶茶失败")
			panic(err.Error())
		}
		milks = append(milks, milk)
	}
	return milks
}
func AddOrderForm(shopname, teaname, sweet, cond, addmaterial string, num, total int) {
	sqlStr := "insert into orderform(username,shopname,teaname,sweet,cond,addmaterial,num,totalprice) values(?,?,?,?,?,?,?,?)"
	inSmt, err := Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译异常:", err)
		panic(err.Error())
	}
	_, err1 := inSmt.Exec(CurrentUser.Username, shopname, teaname, sweet, cond, addmaterial, num, total)
	if err1 != nil {
		fmt.Println("执行异常:", err1)
		panic(err.Error())
	}
}
func GetHistoryOrderForm(username string) []OrderForm {
	sqlStr1 := "select shopname , teaname , sweet , cond from orderform where username = ? "
	sqlStr2 := "select addmaterial , num , totalprice from orderform where username = ? "
	rows1, err := Db.Query(sqlStr1, username)
	if err != nil {
		fmt.Println("初始化查询1失败")
		panic(err.Error())
	}
	rows2, err := Db.Query(sqlStr2, username)
	if err != nil {
		fmt.Println("初始化查询2失败")
		panic(err.Error())
	}
	var orders []OrderForm
	for rows1.Next() && rows2.Next() {
		var order OrderForm
		order.UserName = username
		err = rows1.Scan(&order.ShopName, &order.TeaName, &order.Sweet, &order.Condition)
		if err != nil {
			fmt.Println("逐个查询失败")
			panic(err.Error())
		}
		err = rows2.Scan(&order.AddMaterial, &order.Number, &order.TotalPrice)
		if err != nil {
			fmt.Println("逐个查询失败")
			panic(err.Error())
		}
		orders = append(orders, order)
	}
	return orders
}
