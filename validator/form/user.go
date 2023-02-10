package form

// CreateUserRequest 创建菜单request
type CreateUserRequest struct {
	Nickname  string `json:"nickname"  validate:"required"`  // 昵称
	Username  string `json:"username"  validate:"required"`  // 账号
	Telephone string `json:"telephone"  validate:"required"` // 手机号
	Role      string `json:"role"  validate:"required"`      // 关联角色的uid
	Password  string `json:"password"  validate:"required"`  // 密码
	Email     string `json:"email"`                          // 邮箱
	Sex       string `json:"sex"`                            // 性别
	Status    int    `json:"status"`                         // 状态
}

// DelUserRequest 删除菜单request
type DelUserRequest struct {
	Uid string `json:"uid" validate:"required"` // uid
}

// UpdateUserRequest 编辑菜单request
type UpdateUserRequest struct {
	Nickname  string `json:"nickname"  validate:"required"`  // 昵称
	Username  string `json:"username"  validate:"required"`  // 账号
	Telephone string `json:"telephone"  validate:"required"` // 手机号
	Role      string `json:"role"  validate:"required"`      // 关联角色的uid
	Email     string `json:"email"`                          // 邮箱
	Sex       string `json:"sex"`                            // 性别
	Status    int    `json:"status"`                         // 状态
	Uid       string `json:"uid" validate:"required"`        // uid
}
