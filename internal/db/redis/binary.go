package redis

import (
	"encoding/json"
)

type RedisMarshaller struct{}

// MarshalBinary TODO
func (m RedisMarshaller) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

// UnmarshalBinary ...
func (m *RedisMarshaller) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}
