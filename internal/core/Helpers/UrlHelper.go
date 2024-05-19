package helpers

func URLParse(url string, urlPattern []string) map[string]string {
	parsedURL := TextParse(url, '/')[1:]
	//parsedURL[0] = url
	mappedURL := make(map[string]string)

	if len(parsedURL) > len(urlPattern) {
		ArrayPrint(parsedURL)
		return nil
	}
	for i := 0; i < len(urlPattern); i++ {
		if len(parsedURL) > i {
			mappedURL[urlPattern[i]] = parsedURL[i]

		} else {
			mappedURL[urlPattern[i]] = ""
		}
	}
	return mappedURL
}
