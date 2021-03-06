//go:generate gobatis users.go
package gentest

type Users interface {
	Insert(u *User) (int64, error)

	Update(id int64, u *User) (int64, error)

	DeleteAll() (int64, error)

	Delete(id int64) (int64, error)

	Get(id int64) (*User, error)

	Count() (int64, error)

	GetName(id int64) (string, error)

	// @type select
	Roles(id int64) ([]Role, error)

	UpdateName(id int64, username string) (int64, error)

	InsertName(name string) (int64, error)

	// @filter id = 1
	Find1() ([]User, error)

	// @filter id = 1
	// @filter name = 'a'
	Find2() ([]User, error)

	// @filter id > #{id}
	Find3(id int64) ([]User, error)

	// @filter id > #{id}
	Find4(id int64, name string) ([]User, error)

	// @filter id > #{id}
	// @orderBy created_at ASC
	Find5(id int64, name string) ([]User, error)
}

type UserExDao interface {
	// @record_type User
	InsertName(name string) (int64, error)
}
