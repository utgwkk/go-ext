package jsonext

import "encoding/json"

// NonNullArray is the same type as a slice except that
// it serializes nil as an empty array when serializing to JSON.
type NonNullArray[T any] []T

// MarshalJSON implements json.Marshaler
func (a NonNullArray[T]) MarshalJSON() ([]byte, error) {
	if len(a) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal([]T(a))
}

// UnmarshalJSON implements json.Unmarshaler
func (a *NonNullArray[T]) UnmarshalJSON(b []byte) error {
	var xs []T
	if err := json.Unmarshal(b, &xs); err != nil {
		return err
	}
	*a = xs
	return nil
}
