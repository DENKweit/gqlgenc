package modelgen

import "encoding/json"

type OptionalString struct {
	Value *string
}

func (o *OptionalString) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.Value)
}

type OptionalInt struct {
	Value *int
}

func (o *OptionalInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.Value)
}

type OptionalInt32 struct {
	Value *int32
}

func (o *OptionalInt32) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.Value)
}

type OptionalInt64 struct {
	Value *int64
}

func (o *OptionalInt64) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.Value)
}

type OptionalFloat32 struct {
	Value *float32
}

func (o *OptionalFloat32) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.Value)
}

type OptionalFloat64 struct {
	Value *float64
}

func (o *OptionalFloat64) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.Value)
}

type OptionalBool struct {
	Value *bool
}

func (o *OptionalBool) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.Value)
}
