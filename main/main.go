/*
 * @Author       : Motto motto@hortorgames.com
 * @Description  :
 * @Date         : 2023-06-06 10:04:30
 * @LastEditors  : Motto Yin yjxxtgs@126.com
 * @LastEditTime : 2023-11-06 00:16:32
 * @FilePath     : \go-i18n\main\main.go
 */
package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"time"

	"github.com/go-playground/pure/v5"
	middleware "github.com/go-playground/pure/v5/_examples/middleware/logging-recovery"
	"github.com/motto0808/go-i18n/i18n"
)

var (
	tmpls    *template.Template
	transKey = struct {
		name string
	}{
		name: "transKey",
	}
)

func main() {
	_, path, _, _ := runtime.Caller(0)
	_ = os.Chdir(filepath.Dir(path))
	i18n.InitAll("en")
	var err error
	tmpls, err = template.ParseFiles("index.tpl")
	if err != nil {
		log.Fatalf("Error parsing templates: %s", err)
	}

	r := pure.New()
	r.Use(middleware.LoggingAndRecovery(true), translatorMiddleware)
	r.Get("/", index)

	log.Println("Running on Port :8080")
	log.Println("Try me with URL http://localhost:8080/?locale=en")
	time.AfterFunc(1*time.Second, func() {
		open("http://localhost:8080/?locale=en")
	})
	http.ListenAndServe(":8080", r.Serve())
}

func translatorMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// there are many ways to check, this is just checking for query param &
		// Accept-Language header but can be expanded to Cookie's etc....

		params := r.URL.Query()

		locale := params.Get("locale")
		var t i18n.Translator

		if len(locale) > 0 {
			t = i18n.GetTranslator(locale)
		}
		if t == nil {
			// get and parse the "Accept-Language" http header and return an array
			t = i18n.GetTranslator(pure.AcceptedLanguages(r)...)
		}

		// I would normally wrap ut.Translator with one with my own functions in order
		// to handle errors and be able to use all functions from translator within the templates.
		r = r.WithContext(context.WithValue(r.Context(), transKey, t))

		next(w, r)
	}
}

func index(w http.ResponseWriter, r *http.Request) {

	// get locale translator ( could be wrapped into a helper function )
	t := r.Context().Value(transKey).(i18n.Translator)

	locales := i18n.AvaliableLocales()
	sort.Strings(locales)
	s := struct {
		Trans       i18n.Translator
		Locales     []string
		Now         time.Time
		PositiveNum float64
		NegativeNum float64
		Percent     float64
	}{
		Trans:       t,
		Locales:     locales,
		Now:         time.Now(),
		PositiveNum: 1234576.45,
		NegativeNum: -35900394.34,
		Percent:     96.76,
	}

	if err := tmpls.ExecuteTemplate(w, "index", s); err != nil {
		log.Fatal(err)
	}
}

var commands = map[string]string{
	"windows": "cmd /c start",
	"darwin":  "open",
	"linux":   "xdg-open",
}

var Version = "0.1.0"

// Open calls the OS default program for uri
func open(uri string) error {
	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}

	cmd := exec.Command(run, uri)
	return cmd.Start()
}
