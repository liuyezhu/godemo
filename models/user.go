package models

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey;column:id;type:int unsigned auto_increment;comment:ID"`                    // ID
	Username  string `json:"username" gorm:"column:username;type:varchar(32);not null;uniqueIndex:uk_username;comment:用户名"` // 用户名
	Password  string `json:"password" gorm:"column:password;type:varchar(32);not null;comment:密码"`                          // 密码
	Salt      string `json:"salt" gorm:"column:salt;type:varchar(30);not null;comment:密码盐"`                                 // 密码盐
	Email     string `json:"email" gorm:"column:email;type:varchar(100);not null;uniqueIndex:uk_email;comment:电子邮箱"`        // 电子邮箱
	Mobile    string `json:"mobile" gorm:"column:mobile;type:varchar(11);not null;comment:手机号"`                             // 手机号
	CreatedAt uint   `json:"created_at" gorm:"column:created_at;type:int unsigned;not null;default:0;comment:创建时间"`         // 创建时间
	UpdatedAt uint   `json:"updated_at" gorm:"column:updated_at;type:int unsigned;not null;default:0;comment:更新时间"`         // 更新时间
	Status    int32  `json:"status" gorm:"column:status;type:tinyint;not null;default:0;comment:状态:0=黑名单;1=正常;2=注销"`        // 状态:0=黑名单;1=正常;2=注销
}

// TableName returns the table name of the User model
func (u *User) TableName() string {
	return "user"
}
