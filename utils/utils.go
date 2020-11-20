package main

import (
	"errors"
	"io"
	"net/http"
	"strings"
)

func GetExternalIP() (string,error) {
	response, err := http.Get("http://ip.cip.cc")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	res := ""
	for {
		tmp := make([]byte,32)
		n, err := response.Body.Read(tmp)
		if err != nil {
			if err != io.EOF {
				return "",errors.New("external IP fetch failed,detail:" + err.Error())
			}
			res += string(tmp[:n])
			break
		}
		res += string(tmp[:n])
	}
	return strings.TrimSpace(res),nil
}