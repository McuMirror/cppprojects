package fetdata

import (
	"html/template"
	"time"

	"server/engine/utils"
	"server/engine/wrapper"
)

type BlogPost struct {
	wrap   *wrapper.Wrapper
	object *utils.Sql_blog_post

	user     *User
	category *BlogCategory
}

func (this *BlogPost) load() *BlogPost {
	return this
}

func (this *BlogPost) Id() int {
	if this == nil {
		return 0
	}
	return this.object.A_id
}

func (this *BlogPost) User() *User {
	if this == nil {
		return nil
	}
	if this.user != nil {
		return this.user
	}
	this.user = (&User{wrap: this.wrap}).load()
	this.user.loadById(this.object.A_user)
	return this.user
}

func (this *BlogPost) Name() string {
	if this == nil {
		return ""
	}
	return this.object.A_name
}

func (this *BlogPost) Alias() string {
	if this == nil {
		return ""
	}
	return this.object.A_alias
}

func (this *BlogPost) Category() *BlogCategory {
	if this == nil {
		return nil
	}
	if this.category != nil {
		return this.category
	}
	this.category = (&BlogCategory{wrap: this.wrap}).load(nil)
	this.category.loadById(this.object.A_category)
	return this.category
}

func (this *BlogPost) Briefly() template.HTML {
	if this == nil {
		return template.HTML("")
	}
	return template.HTML(this.object.A_briefly)
}

func (this *BlogPost) Content() template.HTML {
	if this == nil {
		return template.HTML("")
	}
	return template.HTML(this.object.A_content)
}

func (this *BlogPost) DateTimeUnix() int {
	if this == nil {
		return 0
	}
	return this.object.A_datetime
}

func (this *BlogPost) DateTimeFormat(format string) string {
	if this == nil {
		return ""
	}
	return time.Unix(int64(this.object.A_datetime), 0).Format(format)
}

func (this *BlogPost) Active() bool {
	if this == nil {
		return false
	}
	return this.object.A_active > 0
}

func (this *BlogPost) Permalink() string {
	if this == nil {
		return ""
	}
	return "/blog/" + this.object.A_alias + "/"
}
