package fetdata

import (
	"server/engine/utils"
	"server/engine/wrapper"
)

type User struct {
	wrap   *wrapper.Wrapper
	object *utils.Sql_user
}

func (this *User) load() *User {
	return this
}

func (this *User) loadById(id int) {
	if this == nil {
		return
	}
	if this.object != nil {
		return
	}
	this.object = &utils.Sql_user{}
	if err := this.wrap.DB.QueryRow(
		this.wrap.R.Context(),
		`SELECT
			id,
			first_name,
			last_name,
			email,
			admin,
			active,
			address,
			port
		FROM
			users
		WHERE
			id = ?
		LIMIT 1;`,
		id,
	).Scan(
		&this.object.A_id,
		&this.object.A_first_name,
		&this.object.A_last_name,
		&this.object.A_email,
		&this.object.A_admin,
		&this.object.A_active,
		&this.object.A_address,
		&this.object.A_port,
	); *this.wrap.LogCpError(&err) != nil {
		return
	}
}

func (this *User) Id() int {
	if this == nil {
		return 0
	}
	return this.object.A_id
}

func (this *User) FirstName() string {
	if this == nil {
		return ""
	}
	return this.object.A_first_name
}

func (this *User) LastName() string {
	if this == nil {
		return ""
	}
	return this.object.A_last_name
}

func (this *User) Email() string {
	if this == nil {
		return ""
	}
	return this.object.A_email
}

func (this *User) IsAdmin() bool {
	if this == nil {
		return false
	}
	return this.object.A_admin == 1
}

func (this *User) IsActive() bool {
	if this == nil {
		return false
	}
	return this.object.A_active == 1
}
