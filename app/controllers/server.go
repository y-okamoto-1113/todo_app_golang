package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"todo_app_golang/app/models"
	"todo_app_golang/config"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	fmt.Println("=================================================================")
	fmt.Println("filenames in generateHTML() =>", filenames)

	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	fmt.Println("files in generateHTML() =>", files)

	// Must は引数のファイルをあらかじめキャッシュする。エラーの場合は Panic になる。
	templates := template.Must(template.ParseFiles(files...))
	fmt.Println("templates in generateHTML() =>", templates)
	// defineを使って定義したテンプレートを使う場合は、テンプレート名を明示的に指定する必要がある。
	// 第3引数はテンプレートに渡す（展開する）値を渡す。
	templates.ExecuteTemplate(w, "layout", data)
}

func getSession(w http.ResponseWriter, r *http.Request) (s models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		s = models.Session{UUID: cookie.Value}
		if ok, _ := s.CheckSession(); !ok {
			err = fmt.Errorf("Invalid Session!")
		}
	}
	return s, err
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	fmt.Println("files =>", files) // => &{app/views}
	// localhost/static/css/bootstrap.min.css みたいに `/static/` にリクエストが来ると、`static` を無視して `app/views/` 以下の `css/bootstrap.min.css` を見るようにする。
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", TopHandler)
	http.HandleFunc("/signup", SignupHandler)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/authenticate", AuthenticateHandler)
	http.HandleFunc("/logout", LogoutHandler)
	http.HandleFunc("/todos", IndexHandler)
	http.HandleFunc("/todos/new", TodoNewHandler)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
