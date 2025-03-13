package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// WeatherResponse OpenWeatherMap API'den gelen yanıtı temsil eder
type WeatherResponse struct {
	Weather []struct {
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		Humidity  int     `json:"humidity"`
		Pressure  int     `json:"pressure"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	Name string `json:"name"`
	Cod  int    `json:"cod"`
}

type WeatherData struct {
	City        string    `json:"city"`
	Description string    `json:"description"`
	Temperature float64   `json:"temperature"`
	FeelsLike   float64   `json:"feels_like"`
	Humidity    int       `json:"humidity"`
	WindSpeed   float64   `json:"wind_speed"`
	Pressure    int       `json:"pressure"`
	Icon        string    `json:"icon"`
	Timestamp   time.Time `json:"timestamp"`
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
		fmt.Println("❌ .env dosyası yüklenemedi:", err)
		return
	}

	// API anahtarını al
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	if apiKey == "" {
		fmt.Println("❌ OPENWEATHER_API_KEY bulunamadı")
		return
	}

	// Gin web sunucusunu başlat
	r := gin.Default()

	// Statik dosyaları sun
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	// Ana sayfa
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Hava durumu API endpoint'i
	r.GET("/weather", func(c *gin.Context) {
		city := c.DefaultQuery("city", "Istanbul")
		unit := c.DefaultQuery("unit", "c")

		// API URL'sini oluştur
		url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric&lang=tr", city, apiKey)

		// API'ye istek gönder
		resp, err := http.Get(url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "API isteği başarısız oldu"})
			return
		}
		defer resp.Body.Close()

		// API yanıtını oku
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "API yanıtı okunamadı"})
			return
		}

		// API yanıtını parse et
		var weatherResp WeatherResponse
		if err := json.Unmarshal(body, &weatherResp); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "API yanıtı işlenemedi"})
			return
		}

		// API hata kontrolü
		if weatherResp.Cod != 200 {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Şehir bulunamadı"})
			return
		}

		// Hava durumu verilerini hazırla
		weatherData := WeatherData{
			City:        weatherResp.Name,
			Description: weatherResp.Weather[0].Description,
			Temperature: weatherResp.Main.Temp,
			FeelsLike:   weatherResp.Main.FeelsLike,
			Humidity:    weatherResp.Main.Humidity,
			WindSpeed:   weatherResp.Wind.Speed,
			Pressure:    weatherResp.Main.Pressure,
			Icon:        weatherResp.Weather[0].Icon,
			Timestamp:   time.Now(),
		}

		// Fahrenheit'a çevir
		if unit == "f" {
			weatherData.Temperature = (weatherData.Temperature * 9 / 5) + 32
			weatherData.FeelsLike = (weatherData.FeelsLike * 9 / 5) + 32
		}

		c.JSON(http.StatusOK, weatherData)
	})

	// Sunucuyu başlat
	fmt.Println("🌐 Web sunucusu başlatılıyor... http://localhost:8080")
	r.Run(":8080")
}

func getUnitSymbol(unit string) string {
	if unit == "f" {
		return "F"
	}
	return "C"
}
