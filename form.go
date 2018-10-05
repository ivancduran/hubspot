package hubspot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type Form struct {
	Data     map[string]interface{}
	PortalID string
	FormGUID string
	HsContext
}

// hubspotutk

type HsContext struct {
	Hutk     string `json:"hutk"`
	IP       string `json:"ipAddress"`
	PageUrl  string `json:"pageUrl"`
	PageName string `json:"pageName"`
}

// 'email': req.body.email,
// 'firstname': req.body.firstname,
// 'lastname': req.body.lastname,

// 'hs_context': JSON.stringify({
// 	"hutk": req.cookies.hubspotutk,
// 	"ipAddress": req.headers['x-forwarded-for'] || req.connection.remoteAddress,
// 	"pageUrl": "http://www.example.com/form-page",
// 	"pageName": "Example Title"
// })

func NewForm(ID, GUID string) *Form {
	return &Form{
		PortalID: ID,
		FormGUID: GUID,
	}
}

func (f Form) Send(data map[string]interface{}, hs HsContext) bool {
	const (
		hubspotUrl = "https://forms.hubspot.com/uploads/form/v2/%s/%s"
	)

	urlPost := fmt.Sprintf(hubspotUrl, f.PortalID, f.FormGUID)

	jhs, _ := json.Marshal(hs)
	data["hs_context"] = string(jhs)
	f.Data = data

	values := url.Values{}
	for k, v := range data {
		// fmt.Printf("key: %s, value: %s \n", k, v)
		values.Add(k, v.(string))
	}

	fmt.Println(values.Encode())
	countlength := strconv.Itoa(len(values.Encode()))
	// bytes.NewBuffer(b)

	req, err := http.NewRequest("POST", urlPost, bytes.NewBuffer([]byte(values.Encode())))
	req.Host = "forms.hubspot.com"
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Length", countlength)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// x, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("Hubspot body", string(x), resp)

	if resp.StatusCode == 204 {
		return true
	}
	return false
}

// Es un request una peticion POST , la configuracion es la siguiente:
// En la api recibes los datos del formulario y creas el body data por ejemplo:

// var postData = querystring.stringify({
//     'email': req.body.email,
//     'firstname': req.body.firstname,
//     'lastname': req.body.lastname,
// });

// enseguida creas los parametros del request

// var options = {
// 	hostname: 'forms.hubspot.com',
// 	path: '/uploads/form/v2/4021398/014913b6-47e3-4f86-8b51-9ead39eec12b',
// 	method: 'POST',
// 	headers: {
// 		'Content-Type': 'application/x-www-form-urlencoded',
// 		'Content-Length': postData.length
// 	}
// }

// pasas esas configuraciones y los datos y ejecutas el post por ejemplo en nodejs

// var request = https.request(options)
// request.write(postData);
// request.end();
