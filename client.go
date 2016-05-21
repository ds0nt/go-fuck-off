// Package fuck provides a simple interface for interacting with the
// Fuck Off As A Service (FOAAS) API.
// See http://foaas.com for more details on what fucks are given.
package fuck

import (
	"io/ioutil"
	"net/http"
	"strings"
)

const baseURL = "http://foaas.com/"

// ClientVersion is the mot recent FOAAS API version supported
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

// Version will return content with the current FOAAS version number.
func Version() (string, error) {
	return makeRequest("version")
}

// Operations will return a JSON list of operations with names and fields. Note: JSON Only
func Operations() (string, error) {
	return makeRequest("operations")
}

// Thing will return content of the form 'Fuck :thing. - :from',
func Thing(thing, from string) (string, error) {
	return makeRequest(thing, from)
}

// Off will return content of the form 'Fuck off, :name. - :from', e.g. /off/Tom/Chris will return 'Fuck off, Tom - Chris'
func Off(name, from string) (string, error) {
	return makeRequest("off", name, from)
}

// You will return content of the form 'Fuck you, :name. - :from', e.g. /you/Tom/Chris will return 'Fuck you, Tom - Chris'
func You(name, from string) (string, error) {
	return makeRequest("you", name, from)
}

// This will return content of the form 'Fuck this - :from', e.g. /this/Chris will return 'Fuck this. - Chris'
func This(from string) (string, error) {
	return makeRequest("this", from)
}

// That will return content of the form 'Fuck that. - :from', e.g. /that/Chris will return 'Fuck that. - Chris'
func That(from string) (string, error) {
	return makeRequest("that", from)
}

// Everything will return content of the form 'Fuck everything. - :from', e.g. /everything/Chris will return 'Fuck everything. - Chris'
func Everything(from string) (string, error) {
	return makeRequest("everything", from)
}

// Everyone will return content of the form 'Everyone can go and fuck off. - :name', e.g. /everyone/Tom will return 'Everyone can go and fuck off. - Tom'
func Everyone(from string) (string, error) {
	return makeRequest("everyone", from)
}

// Donut will return content of the form ':name, go and take a flying fuck at a rolling donut. - :from', e.g. /donut/Tom/Chris will return 'Tom, go and take a flying fuck at a rolling donut. - Chris'
func Donut(name, from string) (string, error) {
	return makeRequest("donut", name, from)
}

// Shakespeare will return content of the form ':name, Thou clay-brained guts, thou knotty-pated fool, thou whoreson obscene greasy tallow-catch! - :from', e.g. /shakespeare/Falstaff/Prince%20Henry will return 'Falstaff, Thou clay-brained guts, thou knotty-pated fool, thou whoreson obscene greasy tallow-catch! - Prince Henry
func Shakespeare(name, from string) (string, error) {
	return makeRequest("shakespeare", name, from)
}

// Linus will return content of the form ':name, there aren't enough swear-words in the English language, so now I'll have to call you perkeleen vittupää just to express my disgust and frustration with this crap. - :from'. e.g. /linus/Tom/Chris
func Linus(name, from string) (string, error) {
	return makeRequest("linus", name, from)
}

// King will return content of the form 'Oh fuck off, just really fuck off you total dickface. Christ :name, you are fucking thick. - :from'. e.g. /king/Tom/Chris
func King(name, from string) (string, error) {
	return makeRequest("king", name, from)
}

// Pink will return content of the form 'Well, Fuck me pink. - :from'. e.g. /pink/Tom
func Pink(from string) (string, error) {
	return makeRequest("pink", from)
}

// Life will return content of the form 'Fuck my life. - :from', e.g. /life/Phil will return 'Fuck my life. - Phil'.
func Life(from string) (string, error) {
	return makeRequest("life", from)
}

// Chainsaw will return content of the form 'Fuck me gently with a chainsaw, :name. Do I look like Mother Teresa? - :from', e.g. /chainsaw/Chris/Heather will return 'Fuck me gently with a chainsaw, Chris. Do I look like Mother Teresa? - Heather'.
func Chainsaw(name, from string) (string, error) {
	return makeRequest("chainsaw", name, from)
}

// Outside will return content of the form ':name, why don't you go outside and play hide-and-go-fuck-yourself? - :from', e.g. /outside/BigBrother/TheWorld will return 'BigBrother, why don't you go outside and play hide-and-go-fuck-yourself? - TheWorld'.
func Outside(name, from string) (string, error) {
	return makeRequest("outside", name, from)
}

// Thanks will return content of the form 'Fuck you very much. - :from'.
func Thanks(from string) (string, error) {
	return makeRequest("thanks", from)
}

// Flying will return content of the form 'I don't give a flying fuck. - :from'. e.g. /flying/batman
func Flying(from string) (string, error) {
	return makeRequest("flying", from)
}

// Fascinating will return content of the form 'Fascinating story, in what chapter do you shut the fuck up? - :from'.
func Fascinating(from string) (string, error) {
	return makeRequest("fascinating", from)
}

// Madison will return content of the form 'What you've just said is one of the most insanely idiotic things I have ever heard, :name. At no point in your rambling, incoherent response were you even close to anything that could be considered a rational thought. Everyone in this room is now dumber for having listened to it. I award you no points :name, and may God have mercy on your soul. - :from'
func Madison(name, from string) (string, error) {
	return makeRequest("madison", name, from)
}

// Cool 'Cool story, Bro - :from'
func Cool(from string) (string, error) {
	return makeRequest("cool", from)
}

// Field will return content of the form 'And :name said unto :from, "Verily, cast thine eyes upon the field in which I grow my fucks", and :from gave witness unto the field, and saw that it was barren. - :reference'
func Field(name, from, reference string) (string, error) {
	return makeRequest("field", name, from, reference)
}

// Nugget will return content of the form 'Well :name, aren't you a shining example of a rancid fuck-nugget. - :from'
func Nugget(name, from string) (string, error) {
	return makeRequest("nugget", name, from)
}

// Yoda will return content of the form 'Fuck off, you must, :name. - :from'.
func Yoda(name, from string) (string, error) {
	return makeRequest("yoda", name, from)
}

// Ballmer will return content of the form 'Fucking :name is a fucking pussy. I'm going to fucking bury that guy, I have done it before, and I will do it again. I'm going to fucking kill :company. - :from'
func Ballmer(name, company, from string) (string, error) {
	return makeRequest("ballmer", name, company, from)
}

// What will return content of the form 'What the fuck?!. - :from".
func What(from string) (string, error) {
	return makeRequest("what", from)
}

// Because will return content of the form 'Why? Because Fuck you, that's why. - :from'.
func Because(from string) (string, error) {
	return makeRequest("because", from)
}

// CanIUse will return content of the form 'Can you use :tool? Fuck no! - :from'.
func CanIUse(tool, from string) (string, error) {
	return makeRequest("caniuse", tool, from)
}

// Bye will return content of the form 'Fuckity bye! - :from'.
func Bye(from string) (string, error) {
	return makeRequest("bye", from)
}

// Diabetes will return content of the form 'I'd love to stop and chat to you but I'd rather have type 2 diabetes. - :from'.
func Diabetes(from string) (string, error) {
	return makeRequest("diabetes", from)
}

// Bus will return content of the form 'Christ on a bendy-bus, :name, don't be such a fucking faff-arse. - :from'.
func Bus(name, from string) (string, error) {
	return makeRequest("bus", name, from)
}

// Xmas will return content of the form 'Merry Fucking Christmas, :name. - :from'.
func Xmas(name, from string) (string, error) {
	return makeRequest("xmas", name, from)
}

// Bday will return content of the form 'Happy Fucking Birthday, :name. - :from'.
func Bday(name, from string) (string, error) {
	return makeRequest("bday", name, from)
}

// Awesome will return content of the form 'This is Fucking Awesome. - :from', , e.g. /awesome/Macklemore will return 'This is Fucking Awesome. - Macklemore'
func Awesome(from string) (string, error) {
	return makeRequest("awesome", from)
}

// Tucker will return content of the form 'Come the fuck in or fuck the fuck off. - :from', e.g. /tucker/Malcolm%20Tucker will return 'Come the fuck in or fuck the fuck off. - Malcolm Tucker'
func Tucker(from string) (string, error) {
	return makeRequest("tucker", from)
}

// Bucket will return content of the form 'Please choke on a bucket of cocks. - :from'.
func Bucket(from string) (string, error) {
	return makeRequest("bucket", from)
}

// Family will return content of the form 'Fuck you, your whole family, your pets, and your feces. - :from'.
func Family(from string) (string, error) {
	return makeRequest("family", from)
}

// Shutup will return content of the form ':name, shut the fuck up. - :from'.
func Shutup(name, from string) (string, error) {
	return makeRequest("shutup", name, from)
}

// Zayn will return content of the form 'Ask me if I give a motherfuck ?!! - :from'.
func Zayn(from string) (string, error) {
	return makeRequest("zayn", from)
}

// KeepCalm will return content of the form 'Keep the fuck calm and :reaction! - :from'.
func KeepCalm(reaction, from string) (string, error) {
	return makeRequest("keepcalm", reaction, from)
}

// DoSomething will return content of the form ':do the fucking :something! - :from'.
func DoSomething(do, something, from string) (string, error) {
	return makeRequest("dosomething", do, something, from)
}

// Mornin will return content of the form "Happy fuckin' Mornin'! - :from".
func Mornin(from string) (string, error) {
	return makeRequest("mornin", from)
}

// Thumbs will return content of the form 'Who has two thumbs and doesn't give a fuck? :subject. - :from', e.g. /thumbs/This%20Guy/Kob%20Belso will return 'Who has two thumbs and doesn't give a fuck? This Guy. - Kob Belso'
func Thumbs(subject, from string) (string, error) {
	return makeRequest("thumbs", subject, from)
}

// Retard will return content of the form 'You Fucktard! - :from'.
func Retard(from string) (string, error) {
	return makeRequest("retard", from)
}

// Greed will return content of the form 'The point is, ladies and gentleman, that :noun -- for lack of a better word -- is good. :noun is right. :noun works. :noun clarifies, cuts through, and captures the essence of the evolutionary spirit. :noun, in all of its forms -- :noun for life, for money, for love, knowledge -- has marked the upward surge of mankind' - :from.
func Greed(noun, from string) (string, error) {
	return makeRequest("greed", noun, from)
}
