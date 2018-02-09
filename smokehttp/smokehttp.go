package smokehttp

import (
"fmt"
"log"
"net/http"
"time"
"io/ioutil"
)

func SmokeHttp(url string, loopInterval int64) {

	for {
		time.Sleep(time.Duration(loopInterval)*time.Second)
		request, err := http.NewRequest("GET", url, nil)
		res, err := http.DefaultClient.Do(request)

		if err != nil {
			log.Fatal(err) //Something is wrong while sending request
		}

		if res.StatusCode != http.StatusOK {
			log.Fatalf("Success expected: %d", res.StatusCode) //Uh-oh this means our test failed
		} else {
			bodyBytes, _ := ioutil.ReadAll(res.Body)
			fmt.Println(string(bodyBytes))
		}

	}
}
