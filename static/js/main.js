let currentUnit = 'c';
let lastSearchedCity = '';

// Sayfa yüklendiğinde son aranan şehri localStorage'dan al
document.addEventListener('DOMContentLoaded', function() {
    lastSearchedCity = localStorage.getItem('lastSearchedCity') || 'Istanbul';
    document.getElementById('cityInput').value = lastSearchedCity;
    getWeather();
});

async function getWeather() {
    const cityInput = document.getElementById('cityInput');
    const weatherInfo = document.getElementById('weatherInfo');
    const loading = document.getElementById('loading');
    const city = cityInput.value.trim();

    if (!city) {
        showError('Lütfen bir şehir adı girin');
        return;
    }

    // Arama geçmişini kaydet
    lastSearchedCity = city;
    localStorage.setItem('lastSearchedCity', city);

    loading.style.display = 'block';
    weatherInfo.innerHTML = '';

    try {
        const response = await fetch(`/weather?city=${encodeURIComponent(city)}&unit=${currentUnit}`);
        const data = await response.json();

        if (response.ok) {
            displayWeather(data);
            // Başarılı aramada hata mesajını temizle
            clearError();
        } else {
            showError(data.message || 'Bir hata oluştu');
        }
    } catch (error) {
        showError('Sunucu ile iletişim kurulamadı');
    } finally {
        loading.style.display = 'none';
    }
}

function displayWeather(data) {
    const weatherInfo = document.getElementById('weatherInfo');
    const unit = currentUnit === 'c' ? '°C' : '°F';

    weatherInfo.innerHTML = `
        <h2>${data.city}</h2>
        <div class="weather-icon">${getWeatherEmoji(data.description)}</div>
        <div class="weather-details">
            <p><i class="fas fa-temperature-high"></i> Sıcaklık: ${data.temperature}${unit}</p>
            <p><i class="fas fa-thermometer-half"></i> Hissedilen: ${data.feels_like}${unit}</p>
            <p><i class="fas fa-tint"></i> Nem: ${data.humidity}%</p>
            <p><i class="fas fa-wind"></i> Rüzgar: ${data.wind_speed} m/s</p>
            <p><i class="fas fa-tachometer-alt"></i> Basınç: ${data.pressure} hPa</p>
            <p><i class="fas fa-cloud"></i> Durum: ${data.description}</p>
            <p><i class="fas fa-clock"></i> Son Güncelleme: ${new Date(data.timestamp).toLocaleTimeString()}</p>
        </div>
    `;

    // Sayfa başlığını güncelle
    document.title = `${data.city} Hava Durumu - ${data.temperature}${unit}`;
}

function getWeatherEmoji(description) {
    const emojis = {
        'açık': '☀️',
        'parçalı bulutlu': '⛅️',
        'bulutlu': '☁️',
        'yağmurlu': '🌧',
        'gök gürültülü': '⛈',
        'kar yağışlı': '🌨',
        'sisli': '🌫',
        'rüzgarlı': '💨'
    };

    for (const [key, emoji] of Object.entries(emojis)) {
        if (description.toLowerCase().includes(key)) {
            return emoji;
        }
    }
    return '🌤';
}

function toggleUnit() {
    currentUnit = currentUnit === 'c' ? 'f' : 'c';
    const unitToggle = document.getElementById('unitToggle');
    unitToggle.innerHTML = currentUnit === 'c' 
        ? '<i class="fas fa-temperature-high"></i> °C / °F'
        : '<i class="fas fa-temperature-high"></i> °F / °C';
    
    if (lastSearchedCity) {
        getWeather();
    }
}

function showError(message) {
    const weatherInfo = document.getElementById('weatherInfo');
    weatherInfo.innerHTML = `
        <div class="error">
            <i class="fas fa-exclamation-circle"></i>
            ${message}
        </div>
    `;
}

function clearError() {
    const errorElement = document.querySelector('.error');
    if (errorElement) {
        errorElement.remove();
    }
}

// Enter tuşu ile arama yapma
document.getElementById('cityInput').addEventListener('keypress', function(e) {
    if (e.key === 'Enter') {
        getWeather();
    }
});

// Input değiştiğinde hata mesajını temizle
document.getElementById('cityInput').addEventListener('input', clearError);

// Sayfa görünürlüğü değiştiğinde hava durumunu güncelle
document.addEventListener('visibilitychange', function() {
    if (!document.hidden && lastSearchedCity) {
        getWeather();
    }
}); 