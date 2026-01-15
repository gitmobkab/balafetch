package api



// is this a doc ?
func GetFromBalatroApi(params map[string]string) ([]byte, error) {
	const API string = "https://balatrogame.fandom.com/api.php" 
	
	var params_url string = BuildParamsUrl(params)
	var query_url string = API+params_url

	return GetRequest(query_url)
}