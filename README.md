# 🌤 Hava Durumu Uygulaması

Bu uygulama, OpenWeatherMap API'sini kullanarak belirtilen şehrin hava durumu bilgilerini gösterir.

## ✨ Özellikler

- 🏙 Şehir bazlı hava durumu bilgisi
- 🌡 Sıcaklık (Celsius/Fahrenheit)
- 💧 Nem oranı
- 💨 Rüzgar hızı
- 📊 Hava basıncı
- 🌤 Hava durumu açıklaması ve emoji gösterimi
- 📝 JSON formatında çıktı alma
- 🎨 Renkli ve emoji destekli terminal çıktısı
- 🔧 Komut satırı parametreleri ile özelleştirme

## 🚀 Kurulum

1. Projeyi klonlayın:
   ```bash
   git clone https://github.com/kullaniciadi/weather-app.git
   cd weather-app
   ```

2. OpenWeatherMap'ten bir API anahtarı alın (https://openweathermap.org/api)

3. `.env` dosyası oluşturun ve API anahtarınızı ekleyin:
   ```
   OPENWEATHER_API_KEY=your_api_key_here
   ```

4. Uygulamayı derleyin:
   ```bash
   go build
   ```

5. Çalıştırma izinlerini ayarlayın:
   ```bash
   chmod +x weather-app
   ```

## 💻 Kullanım

Temel kullanım:
```bash
./weather-app
```

Şehir belirterek kullanım:
```bash
./weather-app -city="Ankara"
```

Fahrenheit cinsinden sıcaklık için:
```bash
./weather-app -unit=f
```

JSON formatında çıktı almak için:
```bash
./weather-app -json
```

## ⚙️ Parametreler

- `-city`: Hava durumu bilgisi alınacak şehir (varsayılan: Istanbul)
- `-unit`: Sıcaklık birimi (c: Celsius, f: Fahrenheit) (varsayılan: c)
- `-json`: JSON formatında çıktı ver (varsayılan: false)

## 📋 Gereksinimler

- Go 1.16 veya üzeri
- OpenWeatherMap API anahtarı
- Terminal emoji desteği (opsiyonel)

## 🔍 Örnek Çıktı

```
☀️ Istanbul Hava Durumu ☀️
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
🌡 Sıcaklık: 22.5°C
🤔 Hissedilen: 21.8°C
💧 Nem: 65%
💨 Rüzgar: 3.2 m/s
📊 Basınç: 1015 hPa
📝 Durum: açık
🕒 Son Güncelleme: 15:30:45
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
```

## 📝 Lisans

Bu proje MIT lisansı altında lisanslanmıştır. Detaylar için [LICENSE](LICENSE) dosyasına bakın. 