package proto

import(
    "encoding/json"
)

//Resp - 返回值
type Resp struct  {
    ErrNo int `json:"errNo"`
    Data interface{} `json:"data"`
    ErrMsg string `json:"errMsg"`
}


//解析返回
func ParseResp(payload []byte, resp *Resp) error {
    return json.Unmarshal(payload, resp)
}

//解析data
func ParseData(payload []byte, data interface{}) (*Resp, error) {
    resp := new(Resp)
    resp.Data = data
    err := ParseResp(payload, resp)
    if nil != err {
        return nil, err
    }
    return resp, nil
}

//返回结果
func NewResp(errNo int, errMsg string, data interface{}) *Resp {
    resp := new(Resp)
    resp.ErrNo = errNo
    resp.ErrMsg = errMsg
    resp.Data = data
    return resp
}

//正常返回结果
func NewOK(data interface{}) *Resp {
    return NewResp(0, "ok", data)
}
