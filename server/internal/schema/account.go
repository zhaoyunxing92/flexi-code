package schema

type LoginReq struct {
	Email string `json:"email" validate:"required"`

	Password string `json:"pwd" validate:"required"`
}
type LoginResp struct {
	Token string `json:"token"`

	// 员工名称
	Name string `json:"name"`

	// 头像
	Avatar string `json:"avatar"`
}
