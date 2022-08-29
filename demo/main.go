package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func exists(m map[string]string, k string) (v string, ok bool) {
	v, ok = m[k]
	return v, ok
}

func add2ptr(n *int) {
	*n += 2
}

type user struct {
	name     string
	password string
}

func (u user) checkPassword(password string) bool {
	return u.password == password
}

func (u *user) resetPassword(password string) {
	u.password = password
}

func findUser(users []user, name string) (v *user, err error) {
	for _, u := range users {
		if u.name == name {
			return &u, nil
		}
	}
	return nil, errors.New("not found")
}

type point struct {
	x, y int
}

type userInfo struct {
	Name  string
	Age   int `json:"age"`
	Hobby []string
}

func temp() {
	s := make([]string, 6)
	s[0] = "aaa"
	s = append(s, "ddd")
	fmt.Println(s)

	m := make(map[string]int)
	m["one"] = 1
	m["nine"] = 9
	fmt.Println(m)
	fmt.Println(m["nine"])

	delete(m, "one")
	fmt.Println(m)

	s1 := []int{1, 2, 3, 4, 5}
	for i, num := range s1 {
		fmt.Println(i, num)
	}
	for i := range s1 {
		fmt.Println(i)
	}

	m1 := map[string]string{"a": "A", "b": "B"}
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	for k := range m1 {
		fmt.Println(k)
	}

	fmt.Println(exists(m1, "a"))
	fmt.Println(exists(m1, "one"))

	n := 1
	add2ptr(&n)
	fmt.Println(n)

	a := user{name: "n", password: "p"}
	b := user{"n", "p"}
	c := user{name: "n"}
	c.password = "p"
	var d user
	d.name = "n"
	d.password = "p"
	fmt.Println(a, b, c, d)
	fmt.Println(a.checkPassword("a"))
	a.resetPassword("a")
	fmt.Println(a.checkPassword("a"))

	users := []user{{"a", "a"}, {"b", "b"}}
	u, err := findUser(users, "a")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u.name)

	s2 := "hello"
	fmt.Println(strings.Contains(s2, "ll"))
	fmt.Println(strings.Count(s2, "l"))
	fmt.Println(strings.HasPrefix(s2, "hel"))
	fmt.Println(strings.Index(s2, "ll"))
	fmt.Println(strings.Join([]string{"he", "llo", "azz"}, ""))
	fmt.Println(strings.Repeat(s2, 2))             //return a new string
	fmt.Println(strings.Replace(s2, "e", "E", -1)) //var: string tbr r num, return a new string
	fmt.Println(strings.Split("s2---b-c", "-"))    //can split empty string
	fmt.Println(strings.ToLower("ABCDE"))
	fmt.Println(strings.ToUpper("abcde"))
	fmt.Println(len(s2))
	fmt.Println(len("你好")) //Chinese chars

	s3 := "hello"
	n3 := 123
	p3 := point{1, 2}
	fmt.Printf("s=%v\n", s3)
	fmt.Printf("n=%v\n", n3)
	fmt.Printf("p=%v\n", p3)
	fmt.Printf("p=%+v\n", p3)
	fmt.Printf("p=%#v\n", p3)

	f3 := 3.141592653
	fmt.Println(f3)
	fmt.Printf("%.2f\n", f3)

	a4 := userInfo{Name: "n", Age: 18, Hobby: []string{"Golang", "Java"}}
	buf, err := json.Marshal(a4)
	if err != nil {
		panic(err)
	}
	fmt.Println(buf)
	fmt.Println(string(buf))

	buf, err = json.MarshalIndent(a4, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))

	var b4 userInfo
	err = json.Unmarshal(buf, &b4)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", b4)

	now := time.Now()
	fmt.Println(now) //Current time, format yyyy-mm-dd hh:mm:ss + nsec nanoseconds

	t5 := time.Date(2022, 8, 26, 21, 2, 0, 0, time.UTC)

	fmt.Println(t5)
	fmt.Println(t5.Year(), t5.Month(), t5.Day(), t5.Hour(), t5.Minute())
	fmt.Println(t5.Format("2006-01-02 15:04:05"))
	fmt.Println(t5.Format("06-Jan-2 3:4:5")) //other format

	t51 := time.Date(2022, 8, 26, 22, 32, 30, 0, time.UTC)
	diff := t51.Sub(t5)
	fmt.Println(diff)

	t52, err := time.Parse("2006-01-02 15:04:05", "2022-08-26 21:02:00")
	if err != nil {
		panic(err)
	}
	fmt.Println(t52 == t5)
	fmt.Println(now.Unix())

	f6, _ := strconv.ParseFloat("1.234", 64)
	n6, _ := strconv.ParseInt("111", 10, 64)
	n61, _ := strconv.ParseInt("111", 0, 64)
	n62, _ := strconv.ParseInt("1000", 16, 64)  //n63 is same value but diff case
	n63, _ := strconv.ParseInt("0x1000", 0, 64) // case base = 0, check prefix: 2 for 0b, 8 for 0 or 0o, 16 for 0x
	fmt.Println(f6, n6, n61, n62, n63)
	fmt.Println(strconv.Itoa(int(n6)))
	n64, err := strconv.Atoi("AAA")
	fmt.Println(n64, err)

	fmt.Println(os.Args)
	fmt.Println(os.Getenv("PATH"))
	//fmt.Println(os.Setenv("key", "value)) //set env var

	buf, err = exec.Command("grep", "127.0.0.1", "/etc/hosts").CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))
}

func main() {

}
