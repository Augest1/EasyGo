package convert

import (
	"EasyGo/services/model/userModel"
	"EasyGo/services/rpc/user/userRpcModel"
)

// goverter:converter
// goverter:matchIgnoreCase
// goverter:extend SQLStringToPString
// goverter:extend Int64To32
// goverter:extend SQLTimeToString
// goverter:extend SQLNullTimeToString
// goverter:extend SQLNullInt64ToInt64
// goverter:extend SQLNullFloat64ToFloat64
// goverter:extend StringToSQLNullString
// goverter:extend StringToSQLTime
// goverter:extend StringToSQLNullTime
type Converter interface {
	// goverter:ignore state
	// goverter:ignore sizeCache
	// goverter:ignore unknownFields
	UserS2R(source *userModel.Users) *userRpcModel.User
	// goverter:ignore state
	// goverter:ignore sizeCache
	// goverter:ignore unknownFields
	UserR2S(source *userRpcModel.User) *userModel.Users
	//请在以下加其他的转换方法
	UsersS2R(source []*userModel.Users) []*userRpcModel.User
}
