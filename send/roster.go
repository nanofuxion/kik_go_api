package send

import (
	"strings"

	"../utils"
)

//Roster ...
func Roster() (string, string) {

	UUID := utils.GenUUID()

	str := `<iq type="get" id="{UUID}"><query xmlns="jabber:iq:roster" p="8"></query></iq>`

	str = strings.Replace(str, "{UUID}", UUID, 1)
	return str, UUID
}
