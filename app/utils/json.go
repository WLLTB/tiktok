package utils

import (
	"encoding/json"
	"io"
)

// Marshal 将对象序列化为 JSON 字符串
func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// Unmarshal 将 JSON 字符串反序列化为对象
func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// NewEncoder 返回一个 JSON 编码器
func NewEncoder(w io.Writer) *json.Encoder {
	return json.NewEncoder(w)
}

// NewDecoder 返回一个 JSON 解码器
func NewDecoder(r io.Reader) *json.Decoder {
	return json.NewDecoder(r)
}
