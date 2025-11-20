package main

import "fmt"

func main() {
	fmt.Println(login("admin", "123456"))
	fmt.Println(login("admin1", "123456"))
	fmt.Println(login("admin", "1231456"))
	fmt.Println(login("ad2min", "12321456"))
}

func login(account string, pass string) string {
	if account == "admin" && pass == "123456" {
		return "Login Successful"
	} else if account != "admin" && pass == "123456" {
		return "username incorrect"
	} else if account == "admin" && pass != "123456" {
		return "password incorrect"
	} else if account != "admin" && pass != "123456" {
		return "login failed"
	}
	return "Login Failed"
}
