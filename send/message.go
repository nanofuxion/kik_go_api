package send

import (
	"fmt"
	"strings"

	"../utils"
)

//Send ...
func Send(PAYLOAD string, TYPE string, JID string) (string, string) {

	// tl := Tliteral{sType, jid, GJID, uuid, ts}
	TS := utils.Timestampf()
	TSS := TS.STR
	UUID := utils.GenUUID()

	str := "<message cts=\"{TS}\" xmlns=\"jabber:client\" id=\"{UUID}\" to=\"{JID}\" type=\"{TYPE}\"><body>{PAYLOAD}</body><pb/><preview>{PAYLOAD}</preview><ri/></message>"

	fmt.Println(JID)

	str = strings.Replace(str, "{PAYLOAD}", PAYLOAD, -1)
	str = strings.Replace(str, "{TYPE}", TYPE, 1)
	str = strings.Replace(str, "{JID}", JID, -1)
	str = strings.Replace(str, "{UUID}", UUID, -1)
	str = strings.Replace(str, "{TS}", TSS, -1)
	return str, UUID
}
