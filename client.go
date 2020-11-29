package kik_go_api

import (
	"encoding/hex"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/nanofuxion/kik_go_api/send"
	"github.com/nanofuxion/kik_go_api/utils"
)

type cl struct {
	Connect func(chan string)
	SetNode func(string)
	GetNode func()
	// Init     func()
	Captcha func(string)
	// Auth     func()
	SendRaw  func(string)
	SendMsg  func(string, string)
	Settings func(...string)
}

var cVersion utils.Versionstruct
var user string
var pass string

var session []byte
var deviceID string
var androidID string

var node string = ""
var nUUID = ""
var start = *&utils.Bridge

//Client ...
var Client = cl{
	Connect: func(r chan string) {
		msgr := make(chan string, 4)
		// start := *&utils.Bridge
		go start.Connect(msgr)

		intial := func() {
			if node == "" {
				init := send.Init()
				start.SendStanza(init)
			} else {
				LOGIN := send.Login(deviceID, node, user, pass, cVersion)
				start.SendStanza(LOGIN)
			}
		}
		intial()
		receivedOK := 0

		for msg := range msgr {
			if strings.Contains(msg, "jabber:iq:register") {
				re := regexp.MustCompile(`(<node>(.*)</node>)`)
				match := re.FindAllString(msg, 1)
				jid := strings.Replace(match[0], `<node>`, "", 1)
				node = strings.Replace(jid, `</node>`, "", 1)
				intial()
			}
			if strings.Contains(msg, "<k ok") && receivedOK == 0 {
				NODE := send.Node(deviceID, androidID, user, pass, cVersion)
				start.SendStanza(NODE)
				start.Refresh()
				// start = utils.Br{}

				start := *&utils.Bridge
				go start.Connect(msgr)
				receivedOK = 1
			}
			r <- msg //add parser here
		}

	},
	SetNode: func(node string) {

	},
	GetNode: func() {

	},
	// Init: func() {

	// },
	Captcha: func(response string) {

	},
	// Auth: func() {

	// },
	SendRaw: func(msg string) {
		go start.SendStanza(msg)
	},
	SendMsg: func(MSG string, JID string) { //use message and jid to send message
		TYPE := ""
		if strings.Contains(JID, "@groups.kik.com") {
			TYPE = "groupchat"
		} else {
			TYPE = "chat"
		}
		mess, _ := send.Send(MSG, TYPE, JID)
		go start.SendStanza(mess)
	},
	Settings: func(params ...string) { //set login params and version #
		if len(params) == 4 {
			cVersion.VERSION = params[2]
			cVersion.SHA1DIG = params[3]
		} else if len(params) == 3 {
			fmt.Fprintf(os.Stderr, "error: %v\n", "Settings params must be either 2 or 4 inputs")
			os.Exit(1)
		} else {
			cVersion.VERSION = "15.25.0.22493"
			cVersion.SHA1DIG = "pNtboj79GGFYk9w2RbZZTxLpZUY="
		}
		user = params[0]
		pass = params[1]
		sess := utils.GetInfo(pass)
		deviceID = hex.EncodeToString(sess[0:8]) + "e"
		androidID = hex.EncodeToString(sess[8:12])
	},
}
