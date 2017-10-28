package iapi

//go:generate gom exec prmdg struct --file=../apischema/schema/schema.json --package=iapi --output=./struct.go
//go:generate gom exec prmdg jsval --file=../apischema/schema/schema.json --package=iapi --output=./validator.go
