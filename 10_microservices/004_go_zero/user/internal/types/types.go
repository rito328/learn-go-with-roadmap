// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.6

package types

type GetUserRequest struct {
	UserId string `path:"userId"`
}

type GetUserResponse struct {
	Message string `json:"message"`
}
