package send

import (
	"strings"

	"github.com/nanofuxion/kik_go_api/utils"
)

//Login ...
func Raw(STANZA string, TYPE string, JID string) (string, string) {

	// tl := Tliteral{sType, jid, GJID, uuid, ts}
	TS := utils.Timestampf()
	TSS := TS.STR
	UUID := utils.GenUUID()

	STANZA= strings.Replace(STANZA, "{JID}", JID, -1)
	STANZA= strings.Replace(STANZA, "{TYPE}", TYPE, -1)
	STANZA= strings.Replace(STANZA, "{UUID}", UUID, -1)
	STANZA= strings.Replace(STANZA, "{UUID2}", utils.GenUUID(), -1)
	STANZA= strings.Replace(STANZA, "{TS}", TSS, -1)
	return STANZA, UUID
}
