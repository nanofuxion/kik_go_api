package send

import (
	"strings"

	"github.com/nanofuxion/kik_go_api/utils"
)

//Login ...
func Login(DEVICEID string, NODE string, USERNAME string, PASSOWRD string, VERSIONSTRU utils.Versionstruct) string {

	// tl := Tliteral{sType, jid, GJID, uuid, ts}
	TS := utils.Timestampf()
	TSS := TS.STR
	UUID := utils.GenUUID()
	JID := NODE + `@talk.kik.com`
	CV := utils.GenCV(VERSIONSTRU, TS.INT, JID)
	SIGN := utils.GenSig(VERSIONSTRU.VERSION, TS.INT, JID, UUID)
	PASSKEY := utils.PassKey(USERNAME, PASSOWRD)

	str := `<k cv="{CV}" n="1" p="{PASSKEY}" lang="en_US" to="talk.kik.com" conn="WIFI" sid="{UUID}" ts="{TS}" v="{VERSION}" signed="{SIGN}" from="{JID}/CAN{DEVICEID}">`

	str = strings.Replace(str, "{VERSION}", VERSIONSTRU.VERSION, 1)
	str = strings.Replace(str, "{DEVICEID}", DEVICEID, 1)
	str = strings.Replace(str, "{PASSKEY}", PASSKEY, 1)
	str = strings.Replace(str, "{SIGN}", SIGN, 1)
	str = strings.Replace(str, "{JID}", JID, 1)
	str = strings.Replace(str, "{UUID}", UUID, 1)
	str = strings.Replace(str, "{CV}", CV, 1)
	str = strings.Replace(str, "{TS}", TSS, 1)
	return str
}
