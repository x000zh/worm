package proto


import (
    "testing"
    "encoding/json"
)

type testRespData struct {
    Str string `json:"str"`
}

func TestParseData(t *testing.T) {
    data := new(testRespData)
    data.Str = "x"
    resp := NewOK(data)
    resp.ErrNo = 110011
    payload, err := json.Marshal(resp)
    if nil != err {
        t.Error(err)
        return
    }
    t.Logf("encoded %s", payload)
    parsedData := new(testRespData)
    pResp, err := ParseData(payload, parsedData)
    if nil != err {
        t.Error(err)
        return
    }
    if pResp.ErrNo != resp.ErrNo {
        t.Errorf("parse Error %v", payload)
        return
    }
    if parsedData.Str != data.Str {
        t.Errorf("parse Error %v != %v", parsedData, data)
        return
    }
}
