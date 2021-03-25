package resourceserver

import (
	"fmt"
	"net/http"

	"io/ioutil"

	"golang.org/x/oauth2/clientcredentials"
)

type session struct {
	User string
}

func ProtectedEndpoint(c clientcredentials.Config) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		r, err := http.NewRequest(http.MethodGet, "http://localhost:8088/api/v1/projects", nil)
		if err != nil {
			fmt.Fprintf(rw, "<h1>An error occurred!</h1><p>Could not perform introspection request: %v</p>", err)
			return
		}

		r.Header.Add("Authorization", "Bearer "+req.URL.Query().Get("token"))
		resp, err := http.DefaultClient.Do(r)
		if err != nil {
			fmt.Fprintf(rw, "<h1>An error occurred!</h1><p>Could not perform introspection request: %v</p>", err)
			return
		}

		raw, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(rw, "<h1>An error occurred!</h1><p>Could not perform introspection request: %v</p>", err)
			return
		}

		rw.Write(raw)

		//		resp, err := c.Client(context.Background()).PostForm("https://localhost:8088/api/v1/projects", url.Values{"token": []string{req.URL.Query().Get("token")}, "scope": []string{req.URL.Query().Get("scope")}})
		//		if err != nil {
		//			fmt.Fprintf(rw, "<h1>An error occurred!</h1><p>Could not perform introspection request: %v</p>", err)
		//			return
		//		}
		//		defer resp.Body.Close()
		//
		//		var introspection = struct {
		//			Active bool `json:"active"`
		//		}{}
		//		out, _ := ioutil.ReadAll(resp.Body)
		//		if err := json.Unmarshal(out, &introspection); err != nil {
		//			fmt.Fprintf(rw, "<h1>An error occurred!</h1>%s\n%s", err.Error(), out)
		//			return
		//		}
		//
		//		if !introspection.Active {
		//			fmt.Fprint(rw, `<h1>Request could not be authorized.</h1>
		//<a href="/">return</a>`)
		//			return
		//		}
		//
		//		fmt.Fprintf(rw, `<h1>Request authorized!</h1>
		//<code>%s</code><br>
		//<hr>
		//<a href="/">return</a>
		//`,
		//			out,
		//		)
	}
}
