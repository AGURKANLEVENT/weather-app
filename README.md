# 🌤 Hava Durumu Uygulaması

Bu proje, OpenWeatherMap API kullanarak herhangi bir şehrin güncel hava durumu bilgilerini gösteren bir web uygulamasıdır.

## Özellikler

- 🌍 Herhangi bir şehrin hava durumu bilgilerini görüntüleme
- 🌡️ Sıcaklık birimini Celsius/Fahrenheit olarak değiştirme
- 💧 Nem oranı
- 💨 Rüzgar hızı
- 🌡️ Hissedilen sıcaklık
- 📊 Atmosferik basınç
- 🌤️ Hava durumu açıklaması ve emoji
- 🌙 Karanlık mod desteği
- 📱 Mobil uyumlu tasarım

## Kurulum

1. Projeyi klonlayın:
```bash
git clone https://github.com/kullaniciadi/weather-app.git
cd weather-app
```

2. Gerekli bağımlılıkları yükleyin:
```bash
go mod download
```

3. `.env` dosyası oluşturun ve OpenWeatherMap API anahtarınızı ekleyin:
```
OPENWEATHER_API_KEY=your_api_key_here
```

4. Uygulamayı çalıştırın:
```bash
go build
./weather-app
```

5. Tarayıcınızda `http://localhost:8080` adresine gidin.

## API Anahtarı Alma

1. [OpenWeatherMap](https://openweathermap.org/) sitesine gidin
2. Ücretsiz hesap oluşturun
3. API anahtarınızı alın
4. `.env` dosyasına ekleyin

## Kullanım

- Arama kutusuna şehir adını yazın ve "Ara" butonuna tıklayın
- Sıcaklık birimini değiştirmek için "°C / °F" butonunu kullanın
- Son aranan şehir otomatik olarak kaydedilir

## Teknolojiler

- Go
- Gin Web Framework
- OpenWeatherMap API
- HTML5
- CSS3
- JavaScript
- Font Awesome

## Lisans

Bu proje MIT lisansı altında lisanslanmıştır. Detaylar için [LICENSE](LICENSE) dosyasına bakın.

## Katkıda Bulunma

1. Bu depoyu fork edin
2. Yeni bir özellik dalı oluşturun (`git checkout -b yeni-ozellik`)
3. Değişikliklerinizi commit edin (`git commit -am 'Yeni özellik eklendi'`)
4. Dalınıza push yapın (`git push origin yeni-ozellik`)
5. Bir Pull Request oluşturun
