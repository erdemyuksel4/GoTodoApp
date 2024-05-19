package helpers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func GetCookie[T any](cookieName string, w http.ResponseWriter, r *http.Request) *T {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprint(w, "information not found in cookie")
			return nil
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	decodedData, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil {
		fmt.Println("Base64 çözme hatası:", err)
		return nil
	}

	var value T
	err = json.Unmarshal(decodedData, &value)
	if err != nil {
		fmt.Println(err)
		fmt.Println(cookie)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}
	return &value
}

func SetCookie[T any](w http.ResponseWriter, r *http.Request, t *T, cookieName string) {

	tJSON, err := json.Marshal(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//characters := []string{`"`, `\`} // Kaçış karakteri eklemek istediğiniz karakterler
	//result := addEscapeCharacters(string(tJSON), characters)

	//fmt.Println(result)
	encodedData := base64.StdEncoding.EncodeToString([]byte(tJSON))
	fmt.Println(encodedData)
	cookie := &http.Cookie{
		Name:  cookieName,
		Value: encodedData,
		Path:  "/",
	}
	http.SetCookie(w, cookie)
}
func ParseCookies(r *http.Request) map[string]string {
	cookies := make(map[string]string)

	// HTTP başlıklarından "Cookie" başlığını al
	rawCookieHeaders := r.Header["Cookie"]
	fmt.Println(rawCookieHeaders)
	// Her bir "Cookie" başlığı için ayrıştırma işlemi yap
	for _, rawCookieHeader := range rawCookieHeaders {
		// Başlığı ";" karakterlerine göre bölerek ayrıştır
		parts := strings.Split(rawCookieHeader, ";")

		// Her bir ayrıştırılmış parçayı işle
		for _, part := range parts {
			// Boşlukları temizle ve parçayı "=" karakterlerine göre böl
			parts := strings.Split(strings.TrimSpace(part), "=")

			// Eğer parçaların sayısı 2 ise, haritaya ekle
			if len(parts) == 2 {
				fmt.Println(parts[1])
				cookies[parts[0]] = parts[1]
			}
		}
	}

	return cookies
}

func addEscapeCharacters(s string, characters []string) string {
	for _, c := range characters {
		s = strings.ReplaceAll(s, c, "\\"+c)
	}
	return s
}
