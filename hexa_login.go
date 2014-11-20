package main

import (
    "fmt"
    "os"

    "github.com/headzoo/surf"
    "github.com/PuerkitoBio/goquery"
)

func main() {
    bow := surf.NewBrowser()

    bow.AddRequestHeader("Accept", "text/html")
    bow.AddRequestHeader("Accept-Charset", "utf8")

    err := bow.Open("http://hexa.perl.sh")
    if err != nil { panic(err) }
    fmt.Println("[*] Before login : " + bow.Title())

    fm, err := bow.Form("form#log_form")
    if err != nil { panic(err) }
    fm.Input("id", "carpedm20")
    fm.Input("pw", "carpedm20")
    err = fm.Submit()
    if err != nil { panic(err) }

    fmt.Println("[*] Login success : " + bow.Title())

    err = bow.Click("a:contains('carpedm20')")
    if err != nil { panic(err) }
    fmt.Println(bow.Title())

    //bow.Dom().Find("li a").Each(func(_ int, s *goquery.Selection) {
    //    fmt.Println(s.Text() + " : " + s.Attr("href"))
    //})

    bow.Find("li a").Each(func(_ int, s *goquery.Selection) {
        href, exist := s.Attr("href")

        if exist {
            fmt.Printf("%s [%s]\n", s.Text(), href)
        } else {
            fmt.Println(s.Text())
        }
    })

    file, err := os.Create("hexa.php")
    if err != nil { panic(err) }
    defer file.Close()
    bow.Download(file)
}
