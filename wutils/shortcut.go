package wutils

import "encoding/json"




func JsonEncode(v interface{}) string {
 	b,_ := json.Marshal(v)
	return string(b)
}
