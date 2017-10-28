package iapi

import jsval "github.com/lestrrat/go-jsval"

var HealthcheckSelfValidator *jsval.JSVal
var TodoCreateValidator *jsval.JSVal
var TodoInstancesValidator *jsval.JSVal
var TodoSelfValidator *jsval.JSVal
var UserCreateValidator *jsval.JSVal
var UserSelfValidator *jsval.JSVal
var M *jsval.ConstraintMap
var R0 jsval.Constraint
var R1 jsval.Constraint
var R2 jsval.Constraint
var R3 jsval.Constraint
var R4 jsval.Constraint
var R5 jsval.Constraint

func init() {
	M = &jsval.ConstraintMap{}
	R0 = jsval.String()
	R1 = jsval.Reference(M).RefersTo("#/definitions/user/definitions/id")
	R2 = jsval.String()
	R3 = jsval.String().Format("uuid")
	R4 = jsval.String()
	R5 = jsval.String()
	M.SetReference("#/definitions/todo/definitions/name", R0)
	M.SetReference("#/definitions/todo/definitions/userId", R1)
	M.SetReference("#/definitions/user/definitions/email", R2)
	M.SetReference("#/definitions/user/definitions/id", R3)
	M.SetReference("#/definitions/user/definitions/password", R4)
	M.SetReference("#/definitions/user/definitions/username", R5)
	HealthcheckSelfValidator = jsval.New().
		SetName("HealthcheckSelfValidator").
		SetConstraintMap(M).
		SetRoot(
			jsval.EmptyConstraint,
		)

	TodoCreateValidator = jsval.New().
		SetName("TodoCreateValidator").
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("name", "userId").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					"name",
					jsval.Reference(M).RefersTo("#/definitions/todo/definitions/name"),
				).
				AddProp(
					"userId",
					jsval.Reference(M).RefersTo("#/definitions/todo/definitions/userId"),
				),
		)

	TodoInstancesValidator = jsval.New().
		SetName("TodoInstancesValidator").
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					"limit",
					jsval.Integer(),
				).
				AddProp(
					"offset",
					jsval.Integer(),
				),
		)

	TodoSelfValidator = jsval.New().
		SetName("TodoSelfValidator").
		SetConstraintMap(M).
		SetRoot(
			jsval.EmptyConstraint,
		)

	UserCreateValidator = jsval.New().
		SetName("UserCreateValidator").
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("email", "password", "username").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					"email",
					jsval.Reference(M).RefersTo("#/definitions/user/definitions/email"),
				).
				AddProp(
					"password",
					jsval.Reference(M).RefersTo("#/definitions/user/definitions/password"),
				).
				AddProp(
					"username",
					jsval.Reference(M).RefersTo("#/definitions/user/definitions/username"),
				),
		)

	UserSelfValidator = jsval.New().
		SetName("UserSelfValidator").
		SetConstraintMap(M).
		SetRoot(
			jsval.EmptyConstraint,
		)

}
