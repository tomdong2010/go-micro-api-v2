package helper

import jsoniter "github.com/json-iterator/go"

var json = jsoniter.Config{
	EscapeHTML:                    true,
	SortMapKeys:                   true,
	ValidateJsonRawMessage:        true,
	CaseSensitive:                 true,
}.Froze()

func JsonEncode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func JsonDecode(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
