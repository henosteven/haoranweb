package henocurl

import (
    "net/http"
    "io/ioutil"
    "strings"
    "time"
    "context"
    "log"
)

func Get(url string, timeout string) string{
    var params = make(map[string]string)
    td, err := time.ParseDuration(timeout)
    if err != nil {
        td = time.Duration(1) * time.Second
    }
    return httpDo(url, "GET", params, td)
}

func GetSSL() {

}

func Post(url string, argv map[string]string, timeout string) string{
    td, err := time.ParseDuration(timeout)
    if err != nil {
        td = time.Duration(1) * time.Second
    }
    return httpDo(url, "POST", argv, td)
}

func PostSSL() {

}

func httpDo(url, method string, argv map[string]string, td time.Duration) string {
    
    tr := &http.Transport{}
    client := &http.Client{Transport: tr}

    req, err := http.NewRequest(method, url, strings.NewReader(""))
    if err != nil {
        //handle error
    }

    if method == "POST" {
        req.Header.Set("Content-Type", "application/x-www-form-urlencodeed")
    }
    
    var ch  = make(chan *http.Response)
    ctx, cancel := context.WithTimeout(context.Background(), td)
    defer cancel()

    go func() {
        resp, err := client.Do(req)
        if err != nil {
            //handle error
        }
        ch <- resp
    }()

    select {
        case resp := <-ch:
            defer resp.Body.Close()
            body, err := ioutil.ReadAll(resp.Body)
            if err != nil {
                //handle err
            }
            return string(body)
        case <-ctx.Done():
            tr.CancelRequest(req)
            <-ch
            log.Println(ctx.Err())
            return ""
    }
}
