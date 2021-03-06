package article

import (
	"github.com/astaxie/beego/validation"
	"github.com/missdeer/kelly/modules/models"
	"github.com/missdeer/kelly/modules/utils"
)

type ArticleAdminForm struct {
	Create      bool   `form:"-"`
	User        int    `form:"attr(rel,select2-admin-model);attr(data-model,User)" valid:"Required"`
	LastAuthor  int    `form:"attr(rel,select2-admin-model);attr(data-model,User)" valid:""`
	URI         string `valid:"Required;MaxSize(60);Match(/[0-9a-z-./]+/)"`
	Title       string `valid:"Required;MaxSize(60)"`
	Content     string `form:"type(textarea,markdown)" valid:"Required"`
	TitleZhCn   string `valid:"Required;MaxSize(60)"`
	ContentZhCn string `form:"type(textarea,markdown)" valid:"Required"`
	IsPublish   bool   ``
}

func (form *ArticleAdminForm) Valid(v *validation.Validation) {
	user := models.User{Id: form.User}
	if user.Read() != nil {
		v.SetError("User", "admin.not_found_by_id")
	}
}

func (form *ArticleAdminForm) SetFromArticle(article *models.Article) {
	utils.SetFormValues(article, form)

	if article.User != nil {
		form.User = article.User.Id
	}

	if article.LastAuthor != nil {
		form.LastAuthor = article.LastAuthor.Id
	}
}

func (form *ArticleAdminForm) SetToArticle(article *models.Article) {
	utils.SetFormValues(form, article)

	if article.User == nil {
		article.User = &models.User{}
	}
	article.User.Id = form.User

	if article.LastAuthor == nil {
		article.LastAuthor = &models.User{}
	}
	article.LastAuthor.Id = form.LastAuthor

	article.ContentCache = utils.Render(article.Content)
	article.ContentCacheZhCn = utils.Render(article.ContentZhCn)
}
