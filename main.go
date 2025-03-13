package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// WeatherResponse OpenWeatherMap API'den gelen yanıtı temsil eder
type WeatherResponse struct {
	Main struct {
		Temp      float64 `json:"temp"`
		Humidity  int     `json:"humidity"`
		Pressure  int     `json:"pressure"`
		FeelsLike float64 `json:"feels_like"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	Name    string `json:"name"`
	Cod     int    `json:"cod"`
	Message string `json:"message"`
}

func getWeatherEmoji(icon string) string {
	switch icon {
	case "01d", "01n":
		return "☀️" // açık
	case "02d", "02n":
		return "⛅️" // az bulutlu
	case "03d", "03n":
		return "☁️" // bulutlu
	case "04d", "04n":
		return "☁️" // çok bulutlu
	case "09d", "09n":
		return "🌧" // yağmurlu
	case "10d":
		return "🌦" // güneşli ve yağmurlu
	case "10n":
		return "🌧" // gece yağmurlu
	case "11d", "11n":
		return "⛈" // gök gürültülü
	case "13d", "13n":
		return "🌨" // karlı
	case "50d", "50n":
		return "🌫" // sisli
	default:
		return "❓"
	}
}

func main() {
	// .env dosyasını yükle
	if err := godotenv.Load(); err != nil {
		fmt.Println("❌ Hata: .env dosyası yüklenemedi!")
		os.Exit(1)
	}

	// Komut satırı parametrelerini tanımla
	city := flag.String("city", "Istanbul", "Hava durumu bilgisi alınacak şehir")
	unit := flag.String("unit", "c", "Sıcaklık birimi (c: Celsius, f: Fahrenheit)")
	jsonOutput := flag.Bool("json", false, "JSON formatında çıktı ver")
	flag.Parse()

	// API anahtarını .env dosyasından oku
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	if apiKey == "" {
		fmt.Println("❌ Hata: OPENWEATHER_API_KEY çevre değişkeni tanımlanmamış!")
		fmt.Println("Lütfen .env dosyasını düzenleyin ve API anahtarınızı ekleyin.")
		os.Exit(1)
	}

	// API URL'sini oluştur
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric&lang=tr", *city, apiKey)

	// API'ye istek gönder
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("❌ Hata: API isteği başarısız oldu: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Yanıtı oku
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("❌ Hata: Yanıt okunamadı: %v\n", err)
		os.Exit(1)
	}

	// Debug için API yanıtını göster
	if resp.StatusCode != 200 {
		fmt.Printf("❌ API Hatası (HTTP %d):\n%s\n", resp.StatusCode, string(body))
		os.Exit(1)
	}

	// JSON'ı parse et
	var weather WeatherResponse
	if err := json.Unmarshal(body, &weather); err != nil {
		fmt.Printf("❌ Hata: JSON ayrıştırılamadı: %v\n", err)
		fmt.Printf("API Yanıtı:\n%s\n", string(body))
		os.Exit(1)
	}

	// API hata kontrolü
	if weather.Cod != 200 {
		fmt.Printf("❌ API Hatası: %s\n", weather.Message)
		os.Exit(1)
	}

	// Weather array kontrolü
	if len(weather.Weather) == 0 {
		fmt.Println("❌ Hata: Hava durumu verisi bulunamadı!")
		os.Exit(1)
	}

	// Sıcaklığı dönüştür
	temp := weather.Main.Temp
	feelsLike := weather.Main.FeelsLike
	if *unit == "f" {
		temp = (temp * 9 / 5) + 32
		feelsLike = (feelsLike * 9 / 5) + 32
	}

	if *jsonOutput {
		// JSON çıktısı
		output := map[string]interface{}{
			"city":        weather.Name,
			"temperature": temp,
			"feels_like":  feelsLike,
			"humidity":    weather.Main.Humidity,
			"pressure":    weather.Main.Pressure,
			"wind_speed":  weather.Wind.Speed,
			"description": weather.Weather[0].Description,
			"icon":        weather.Weather[0].Icon,
			"timestamp":   time.Now().Format(time.RFC3339),
		}
		jsonData, _ := json.MarshalIndent(output, "", "  ")
		fmt.Println(string(jsonData))
		return
	}

	// Normal çıktı
	emoji := getWeatherEmoji(weather.Weather[0].Icon)
	fmt.Printf("\n%s %s Hava Durumu %s\n", emoji, weather.Name, emoji)
	fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
	fmt.Printf("🌡 Sıcaklık: %.1f°%s\n", temp, getUnitSymbol(*unit))
	fmt.Printf("🤔 Hissedilen: %.1f°%s\n", feelsLike, getUnitSymbol(*unit))
	fmt.Printf("💧 Nem: %d%%\n", weather.Main.Humidity)
	fmt.Printf("💨 Rüzgar: %.1f m/s\n", weather.Wind.Speed)
	fmt.Printf("📊 Basınç: %d hPa\n", weather.Main.Pressure)
	fmt.Printf("📝 Durum: %s\n", weather.Weather[0].Description)
	fmt.Printf("🕒 Son Güncelleme: %s\n", time.Now().Format("15:04:05"))
	fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
}

func getUnitSymbol(unit string) string {
	if unit == "f" {
		return "F"
	}
	return "C"
}
