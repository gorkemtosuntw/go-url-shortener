# Go URL Kısaltıcı

## Özellikler
- PostgreSQL ile URL depolama
- URL doğrulama
- Kısa URL oluşturma
- Tıklanma sayısı takibi
- Yönlendirme

## Gereksinimler
- Go 1.21+
- PostgreSQL
- Gerekli paketler: 
  - github.com/lib/pq
  - github.com/gorilla/mux
  - github.com/google/uuid

## Kurulum
1. Veritabanı ayarlarını `main.go` içinde yapılandırın
2. Gerekli paketleri yükleyin: `go mod tidy`
3. Uygulamayı çalıştırın: `go run .`

## API Kullanımı
### URL Kısaltma
`POST /shorten`
- İstek Gövdesi: `{"original_url": "https://example.com"}`
- Yanıt: `{"short_url": "http://localhost:8080/kisa-kod"}`

### Yönlendirme
`GET /{short_code}` otomatik olarak orijinal URL'ye yönlendirir

## Lisans
MIT Lisansı
