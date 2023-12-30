package weather

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/joho/godotenv"
	"os"
)

type weatherData struct {
	Weather []struct{
		Description string
	}

	Main struct {
		Feels_like float64 
		Humidity int
	}
}

type coOrd []struct {
	Lat float64
	Lon float64
}

type aqi struct {
	List []struct {
		Main struct {
			Aqi int 
		} 
		Components struct {
			Co   float64 
			No   float64
			No2  float64 
			O3   float64 
			So2  float64 
			Pm25 float64 
			Pm10 float64 
			Nh3  float64 
		} 
	} 
}


func getApiKey(key string) string{
	
	if err:= godotenv.Load();err!=nil{
		panic(err)
	}

	return os.Getenv(key)
	
}

var apiKey = getApiKey("apiKey")

func GetCoOrd(city string) (coOrd,error){
	url := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s&limit=&appid=%s",city,apiKey)
	req,err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %s\n",err)
	}
	if req.StatusCode > 399 {
		return coOrd{} , fmt.Errorf("bad status code : %v",req.StatusCode)
	}

	defer req.Body.Close()
	var resp coOrd
	if err := json.NewDecoder(req.Body).Decode(&resp); err!=nil {
		return coOrd{},err
	}

	return resp,nil
	
}

func CurrentWeather(lat string,lon string)(weatherData , error){
	
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s",lat,lon,apiKey)

	req,err := http.Get(url)
	if err!= nil{
		return weatherData{},err	
	}
	defer req.Body.Close()

	var response weatherData
	if err:= json.NewDecoder(req.Body).Decode(&response); err!=nil {
		return weatherData{},err	
	}

	return response,nil
	
}
// 1day = 86400 seconds

func AqiValue(lat string,lon string)(aqi, error){
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/air_pollution?lat=%s&lon=%s&appid=%s",lat,lon,apiKey)
	req,err := http.Get(url)
	if err!= nil{
		return aqi{},err	
	}
	defer req.Body.Close()
	var response aqi
	if err:= json.NewDecoder(req.Body).Decode(&response); err!=nil {
		return aqi{},err	
	}

	return response,nil
}
	