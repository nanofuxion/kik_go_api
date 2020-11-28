package send

import (
	"strings"

	"../utils"
)

//Node ...
func Node(DEVICEID string, ANDROIDID string, USERNAME string, PASSOWRD string, VERSIONSTRU utils.Versionstruct) string {

	UUID := utils.GenUUID()
	PASSKEY := utils.PassKey(USERNAME, PASSOWRD)
	VERSIONNO := VERSIONSTRU.VERSION

	str := `<iq type="set" id="{UUID}"><query xmlns="jabber:iq:register"><username>{USERNAME}</username><passkey-u>{PASSKEY}</passkey-u><device-id>{DEVICEID}</device-id><install-referrer>utm_source=google-play&amp;utm_medium=organic</install-referrer><operator>310260</operator><install-date>1595173064</install-date><device-type>android</device-type><brand>Google</brand><logins-since-install>1</logins-since-install><version>{VERSIONNO}</version><lang>en_US</lang><android-sdk>29</android-sdk><registrations-since-install>0</registrations-since-install><prefix>CAN</prefix><android-id>{ANDROIDID}</android-id><model>Google Pixel 4 XL - 10 - API 29 - 1440x2960</model><challenge><response>undefined</response></challenge></query></iq>`

	str = strings.Replace(str, "{DEVICEID}", DEVICEID, 1)
	str = strings.Replace(str, "{ANDROIDID}", ANDROIDID, 1)
	str = strings.Replace(str, "{VERSIONNO}", VERSIONNO, 1)
	str = strings.Replace(str, "{PASSKEY}", PASSKEY, 1)
	str = strings.Replace(str, "{USERNAME}", USERNAME, 1)
	str = strings.Replace(str, "{UUID}", UUID, 1)
	return str
}
