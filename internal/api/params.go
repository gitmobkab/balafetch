package api

import "net/url"

func BuildParamsUrl(params map[string]string) string{
	params_values := url.Values{}
	
	for key,val := range params{
		params_values.Add(key,val)
	}

	var params_url string = "?" + params_values.Encode()
	return  params_url
}