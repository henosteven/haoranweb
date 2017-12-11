package henocurl

import (
    "net/http"
    "io/ioutil"
    "strings"
)

func Get(url string) string{
    var params = make(map[string]string)
    return httpDo(url, "GET", params)
}

func GetSSL() {

}

func Post(url string, argv map[string]string) string{
    return httpDo(url, "POST", argv)
}

func PostSSL() {

}

func httpDo(url, method string, argv map[string]string) string{
    
    client := &http.Client{}

    req, err := http.NewRequest(method, url, strings.NewReader(""))
    if err != nil {
        //handle error
    }

    if method == "POST" {
        req.Header.Set("Content-Type", "application/x-www-form-urlencodeed")
    }

    resp, err := client.Do(req)
    defer resp.Body.Close()
    if err != nil {
        //handle error
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        //handle err
    }
    
    return string(body)
}
