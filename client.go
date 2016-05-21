package fuck

import (
	"io/ioutil"
	"net/http"
	"strings"
)

const baseURL = "http://foaas.com/"
const ClientVersion = "0.1.9"

func makeRequest(params ...string) (string, error) {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, baseURL+strings.Join(params, "/"), nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "text/plain")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}

func Version() (string, error) {
	return makeRequest("version")
}

func Thing(thing, from string) (string, error) {
	return makeRequest(thing, from)
}

func Operations() (string, error) {
	return makeRequest("operations")
}

func Off(name, from string) (string, error) {
	return makeRequest("off", name, from)
}

func You(name, from string) (string, error) {
	return makeRequest("you", name, from)
}

func This(from string) (string, error) {
	return makeRequest("this", from)
}

func That(from string) (string, error) {
	return makeRequest("that", from)
}

func Everything(from string) (string, error) {
	return makeRequest("everything", from)
}

func Everyone(from string) (string, error) {
	return makeRequest("everyone", from)
}

func Donut(name, from string) (string, error) {
	return makeRequest("donut", name, from)
}

func Shakespeare(name, from string) (string, error) {
	return makeRequest("shakespeare", name, from)
}

func Linus(name, from string) (string, error) {
	return makeRequest("linus", name, from)
}

func King(name, from string) (string, error) {
	return makeRequest("king", name, from)
}

func Pink(from string) (string, error) {
	return makeRequest("pink", from)
}
func Life(from string) (string, error) {
	return makeRequest("life", from)
}
func Chainsaw(name, from string) (string, error) {
	return makeRequest("chainsaw", name, from)
}

func Outside(name, from string) (string, error) {
	return makeRequest("outside", name, from)
}

func Thanks(from string) (string, error) {
	return makeRequest("thanks", from)
}
func Flying(from string) (string, error) {
	return makeRequest("flying", from)
}
func Fascinating(from string) (string, error) {
	return makeRequest("fascinating", from)
}
func Madison(name, from string) (string, error) {
	return makeRequest("madison", name, from)
}

func Cool(from string) (string, error) {
	return makeRequest("cool", from)
}

func Field(name, from, reference string) (string, error) {
	return makeRequest("field", name, from, reference)
}

func Nugget(name, from string) (string, error) {
	return makeRequest("nugget", name, from)
}

func Yoda(name, from string) (string, error) {
	return makeRequest("yoda", name, from)
}

func Ballmer(name, company, from string) (string, error) {
	return makeRequest("ballmer", name, company, from)
}

func What(from string) (string, error) {
	return makeRequest("what", from)
}
func Because(from string) (string, error) {
	return makeRequest("because", from)
}
func CanIUse(tool, from string) (string, error) {
	return makeRequest("caniuse", tool, from)
}

func Bye(from string) (string, error) {
	return makeRequest("bye", from)
}
func Diabetes(from string) (string, error) {
	return makeRequest("diabetes", from)
}
func Bus(name, from string) (string, error) {
	return makeRequest("bus", name, from)
}

func Xmas(name, from string) (string, error) {
	return makeRequest("xmas", name, from)
}

func Bday(name, from string) (string, error) {
	return makeRequest("bday", name, from)
}

func Awesome(from string) (string, error) {
	return makeRequest("awesome", from)
}
func Tucker(from string) (string, error) {
	return makeRequest("tucker", from)
}
func Bucket(from string) (string, error) {
	return makeRequest("bucket", from)
}
func Family(from string) (string, error) {
	return makeRequest("family", from)
}
func Shutup(name, from string) (string, error) {
	return makeRequest("shutup", name, from)
}

func Zayn(from string) (string, error) {
	return makeRequest("zayn", from)
}
func KeepCalm(reaction, from string) (string, error) {
	return makeRequest("keepcalm", reaction, from)
}

func DoSomething(do, something, from string) (string, error) {
	return makeRequest("dosomething", do, something, from)
}

func Mornin(from string) (string, error) {
	return makeRequest("mornin", from)
}
func Thumbs(subject, from string) (string, error) {
	return makeRequest("thumbs", subject, from)
}

func Retard(from string) (string, error) {
	return makeRequest("retard", from)
}
func Greed(noun, from string) (string, error) {
	return makeRequest("greed", noun, from)
}
