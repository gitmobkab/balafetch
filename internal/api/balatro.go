package api



// GetFromBalatroApi takes a map of query parameters and a timeout value, constructs the full API URL, and performs a GET request to the Balatro API.
// It returns the response data as a byte slice or an error if the request fails.
func GetFromBalatroApi(params map[string]string, timeout int) ([]byte, error) {
	const API string = "https://balatrogame.fandom.com/api.php" 
	
	var params_url string = BuildParamsUrl(params)
	var query_url string = API+params_url

	return GetRequest(query_url, timeout)
}