// An open source light community based on wetalk.
package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/beego/social-auth"
	"github.com/missdeer/kelly/cache"
	"github.com/missdeer/kelly/modules/utils"
	"github.com/missdeer/kelly/routers/admin"
	"github.com/missdeer/kelly/routers/api"
	"github.com/missdeer/kelly/routers/article"
	"github.com/missdeer/kelly/routers/attachment"
	"github.com/missdeer/kelly/routers/auth"
	"github.com/missdeer/kelly/routers/base"
	"github.com/missdeer/kelly/routers/pay"
	"github.com/missdeer/kelly/routers/post"
	"github.com/missdeer/kelly/setting"

	_ "github.com/go-sql-driver/mysql"
)

// We have to call a initialize function manully
// because we use `bee bale` to pack static resources
// and we cannot make sure that which init() execute first.
func initialize() {
	flag.StringVar(&setting.AppConfPath, "rc", "conf/app.ini", "configuration file")
	flag.Parse()

	setting.LoadConfig()
	setting.LoadAds()
	setting.LoadContacts()

	if err := utils.InitSphinxPools(); err != nil {
		beego.Error(fmt.Sprint("sphinx init pool", err))
	}

	setting.SocialAuth = social.NewSocial("/login/", auth.SocialAuther)
	setting.SocialAuth.ConnectSuccessURL = "/settings/profile"
	setting.SocialAuth.ConnectFailedURL = "/settings/profile"
	setting.SocialAuth.ConnectRegisterURL = "/register/connect"
	setting.SocialAuth.LoginURL = "/login"
}

func unauthorized_handler(rw http.ResponseWriter, r *http.Request) {
	response := []byte("<html><body><meta http-equiv=\"refresh\" content=\"0;url=/401\"></body></html>")
	rw.Write(response)
}

func forbidden_handler(rw http.ResponseWriter, r *http.Request) {
	response := []byte("<html><body><meta http-equiv=\"refresh\" content=\"0;url=/403\"></body></html>")
	rw.Write(response)
}

func not_found_handler(rw http.ResponseWriter, r *http.Request) {
	response := []byte("<html><body><meta http-equiv=\"refresh\" content=\"0;url=/404\"></body></html>")
	rw.Write(response)
}

func internal_server_error_handler(rw http.ResponseWriter, r *http.Request) {
	response := []byte("<html><body><meta http-equiv=\"refresh\" content=\"0;url=/500\"></body></html>")
	rw.Write(response)
}

func service_unavailable_handler(rw http.ResponseWriter, r *http.Request) {
	response := []byte("<html><body><meta http-equiv=\"refresh\" content=\"0;url=/503\"></body></html>")
	rw.Write(response)
}

func main() {
	initialize()

	beego.Info("AppPath:", beego.AppPath)

	if setting.IsProMode {
		beego.Info("Product mode enabled")
	} else {
		beego.Info("Development mode enabled")
	}
	beego.Info(beego.BConfig.AppName, setting.APP_VER, setting.AppUrl)

	if !setting.IsProMode {
		beego.SetStaticPath("/static_source", "static_source")
		beego.BConfig.WebConfig.DirectoryIndex = true
	}

	// Add Filters
	beego.InsertFilter("/img/*", beego.BeforeRouter, attachment.ImageFilter)

	beego.InsertFilter("/captcha/*", beego.BeforeRouter, setting.Captcha.Handler)

	// Register routers.
	posts := new(post.PostListRouter)
	beego.Router("/", posts, "get:Home")
	beego.Router("/orca.txt", posts, "get:ORCA;head:ORCA")
	beego.Router("/:slug(recent|best|cold|favs|follow)", posts, "get:Navs")
	beego.Router("/category/:slug", posts, "get:Category")
	beego.Router("/topic/:slug", posts, "get:Topic;post:TopicSubmit")

	auxiliaryR := new(post.AuxiliaryRouter)
	beego.Router("/faq", auxiliaryR, "get:FAQ")
	beego.Router("/contact", auxiliaryR, "get:Contact")
	beego.Router("/about", auxiliaryR, "get:About")
	beego.Router("/401", auxiliaryR, "get:Err401")
	beego.Router("/403", auxiliaryR, "get:Err403")
	beego.Router("/404", auxiliaryR, "get:Err404")
	beego.Router("/500", auxiliaryR, "get:Err500")
	beego.Router("/503", auxiliaryR, "get:Err503")

	beego.ErrorHandler("401", unauthorized_handler)
	beego.ErrorHandler("403", forbidden_handler)
	beego.ErrorHandler("404", not_found_handler)
	beego.ErrorHandler("500", internal_server_error_handler)
	beego.ErrorHandler("503", service_unavailable_handler)

	postR := new(post.PostRouter)
	beego.Router("/new", postR, "get:New;post:NewSubmit")
	beego.Router("/post/:post([0-9]+)", postR, "get:Single;post:SingleSubmit")
	beego.Router("/post/:post([0-9]+)/edit", postR, "get:Edit;post:EditSubmit")
	beego.Router("/post/:post([0-9]+)/append", postR, "get:Append;post:AppendSubmit")

	if setting.NativeSearch || setting.SphinxEnabled {
		searchR := new(post.SearchRouter)
		beego.Router("/search", searchR, "get:Get")
	}

	user := new(auth.UserRouter)
	beego.Router("/user/:username/comments", user, "get:Comments")
	beego.Router("/user/:username/posts", user, "get:Posts")
	beego.Router("/user/:username/following", user, "get:Following")
	beego.Router("/user/:username/followers", user, "get:Followers")
	beego.Router("/user/:username/favs", user, "get:Favs")
	beego.Router("/user/:username", user, "get:Home")

	login := new(auth.LoginRouter)
	beego.Router("/login", login, "get:Get;post:Login")
	beego.Router("/logout", login, "get:Logout")

	//beego.InsertFilter("/login/*/access", beego.BeforeRouter, auth.OAuthAccess)
	//beego.InsertFilter("/login/*", beego.BeforeRouter, auth.OAuthRedirect)

	socialR := new(auth.SocialAuthRouter)
	beego.Router("/register/connect", socialR, "get:Connect;post:ConnectPost")

	register := new(auth.RegisterRouter)
	beego.Router("/register", register, "get:Get;post:Register")
	beego.Router("/active/success", register, "get:ActiveSuccess")
	beego.Router("/active/:code([0-9a-zA-Z]+)", register, "get:Active")

	settings := new(auth.SettingsRouter)
	beego.Router("/settings/profile", settings, "get:Profile;post:ProfileSave")

	forgot := new(auth.ForgotRouter)
	beego.Router("/forgot", forgot)
	beego.Router("/reset/:code([0-9a-zA-Z]+)", forgot, "get:Reset;post:ResetPost")

	upload := new(attachment.UploadRouter)
	beego.Router("/upload", upload, "post:Post")

	apiR := new(api.ApiRouter)
	beego.Router("/api/user", apiR, "post:Users")
	beego.Router("/api/md", apiR, "post:Markdown")
	beego.Router("/api/post", apiR, "post:PostToggle")

	adminDashboard := new(admin.AdminDashboardRouter)
	beego.Router("/admin", adminDashboard)

	adminR := new(admin.AdminRouter)
	beego.Router("/admin/model/get", adminR, "post:ModelGet")
	beego.Router("/admin/model/select", adminR, "post:ModelSelect")

	forwarderR := new(post.ForwarderRouter)
	beego.Router("/i/:id([0-9]+)", forwarderR, "get:TaobaoItem")

	alipayR := new(alipay.AlipayRouter)
	beego.Router("/alipay", alipayR, "get:Pay;post:Pay")

	routes := map[string]beego.ControllerInterface{
		"user":     new(admin.UserAdminRouter),
		"post":     new(admin.PostAdminRouter),
		"comment":  new(admin.CommentAdminRouter),
		"topic":    new(admin.TopicAdminRouter),
		"category": new(admin.CategoryAdminRouter),
		"article":  new(admin.ArticleAdminRouter),
	}
	for name, router := range routes {
		beego.Router(fmt.Sprintf("/admin/:model(%s)", name), router, "get:List")
		beego.Router(fmt.Sprintf("/admin/:model(%s)/:id(new)", name), router, "get:Create;post:Save")
		beego.Router(fmt.Sprintf("/admin/:model(%s)/:id([0-9]+)", name), router, "get:Edit;post:Update")
		beego.Router(fmt.Sprintf("/admin/:model(%s)/:id([0-9]+)/:action(delete)", name), router, "get:Confirm;post:Delete")
	}

	// "robots.txt"
	beego.Router("/robots.txt", &base.RobotRouter{})

	articleR := new(article.ArticleRouter)
	beego.Router("/a/:slug", articleR, "get:Show")

	if beego.BConfig.RunMode == "dev" {
		beego.Router("/test/:tmpl(mail/.*)", new(base.TestRouter))
	}

	cache.Init()

	go api.ClearTodayReplys()
	// For all unknown pages.
	beego.Run()
}
