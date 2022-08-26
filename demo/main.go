package main

import (
	"errors"
	"fmt"
	"strings"
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
}

func main() {
	a := "hello"
	fmt.Println(strings.Contains(a, "ll"))
	fmt.Println(strings.Count(a, "l"))
	fmt.Println(strings.HasPrefix(a, "hel"))
	fmt.Println(strings.Index(a, "ll"))
	fmt.Println(strings.Join([]string{"he", "llo", "azz"}, ""))
	fmt.Println(strings.Repeat(a, 2))
}
