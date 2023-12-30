package cmdParser

import (
	"fmt"
	"os"
	"weatherCli/internal/weather"
	"log"
	"github.com/fatih/color"
	

)

type Command struct{
	name string
	description string
	CmdFuncMain func(string)
	CmdFunc func()
	hasInput bool
	
}



func GetCommands() map[string]Command{
	
	return map[string]Command{
		
		"cw" :		{
							name:   		"cw",
							description: 	"shows the current weather of the given city ",
							CmdFuncMain: 	getWeather,
							hasInput: 		true,
							
						},
						
		"aqi" : 		{
							name:   		"aqi",
							description: 	"shows the Air Pollution Index of the given city and air composition",
							CmdFuncMain: 	getAQI,
							hasInput: 		true,
							
						},
						
		"help": {
					name:			"help",
					description: 	"Prints all the command",
					CmdFunc: 		helpFunc,
					hasInput: 		false,
					
				},
			
		"exit": {
					name: 			"exit",
					description: 	"Stops the Command line",
					CmdFunc:		exitFunc,
					hasInput: 		false,
					
				},
	}
	
}

func getCoOrdinates(city string) (lat string,lon string){
	coOrdinates , err := weather.GetCoOrd(city)
	defer func(){
		if r:=recover();r!=nil{
			color.Red("invalid city name")
			os.Exit(1)
		}
	}()
	
	if err!= nil{
		log.Fatal(err)
	}
	lat = fmt.Sprintf("%f",coOrdinates[0].Lat)
	lon = fmt.Sprintf("%f",coOrdinates[0].Lon)
	return lat,lon
}

func getWeather(cityName string){

	lat,lon := getCoOrdinates(cityName)

	weatherInfo, err := weather.CurrentWeather(lat,lon)
	if err!= nil {
		log.Fatal(err)
	}


	fmt.Printf("=> weather condition -->   ")
	color.HiWhite("%s",weatherInfo.Weather[0].Description)
	fmt.Printf("=> Temperature       -->   ")
	
	temp := weatherInfo.Main.Feels_like-273.15
	if temp > 30{
		color.Red("%.1f",temp)
		fmt.Printf("=> its hot out outside!\n")
	}else{
		color.Cyan("%.1f",temp)
		fmt.Printf("=> its cold out outside!\n")
	}
	
	color.HiMagenta("=> Humidity         -->      %d",weatherInfo.Main.Humidity)

	fmt.Println("--------------------------------------------------------------------------------------------------------------------")
}

func getAQI(cityName string){

	lat,lon := getCoOrdinates(cityName)

	AqiInfo,err := weather.AqiValue(lat,lon)

	if err!= nil {
		log.Fatal(err)
	}

	aqi:= AqiInfo.List[0].Main.Aqi

	switch aqi{
		case 1:
			fmt.Printf("=> AQI   -->  ")
			color.HiWhite("%d",aqi)
			color.HiWhite("=> Clean air\n")

		case 2:
			fmt.Printf("=> AQI   -->  ")
			color.Cyan("%d",aqi)
			color.HiCyan("=> Good quality air\n")
		
		case 3:
			fmt.Printf("=> AQI   -->  ")
			color.Yellow("%d",aqi)
			color.HiYellow("=> A bit polluted\n")
		
		case 4:
			fmt.Printf("=> AQI   -->  ")
			color.Red("%d",aqi)
			color.Red("=> Time to put on your mask\n")
		
		case 5:
			fmt.Printf("=> AQI   -->  ")
			color.HiRed("%d",aqi)
			color.HiRed("=> LEAVE THE CITY")
		default:
			fmt.Printf("=> welp somethings wrong")
	}
	
	color.Green("------------------------------------------------ AIR COMPOSITION ---------------------------------------------------")
	fmt.Printf("=> Сoncentration of CO (Carbon monoxide)                 ---->       ")
	color.HiRed("%.2f μg/m3",AqiInfo.List[0].Components.Co)
	fmt.Printf("=> Сoncentration of NO (Nitrogen monoxide)               ---->       ")
	color.Green("%.2f μg/m3",AqiInfo.List[0].Components.No)
	fmt.Printf("=> Сoncentration of NO2 (Nitrogen dioxide),              ---->       ")
	color.HiGreen("%.2f μg/m3",AqiInfo.List[0].Components.No2)
	fmt.Printf("=> Сoncentration of O3 (Ozone)                           ---->       ")
	color.HiCyan("%.2f μg/m3",AqiInfo.List[0].Components.O3)
	fmt.Printf("=> Сoncentration of SO2 (Sulphur dioxide)                ---->       ")
	color.HiYellow("%.2f μg/m3",AqiInfo.List[0].Components.So2)
	fmt.Printf("=> Сoncentration of PM2.5 (Fine particles matter)        ---->       ")
	color.Cyan("%.2f μg/m3",AqiInfo.List[0].Components.Pm25)
	fmt.Printf("=> Сoncentration of PM10 (Coarse particulate matter)     ---->       ")
	color.HiCyan("%.2f μg/m3",AqiInfo.List[0].Components.Pm10)
	fmt.Printf("=> Сoncentration of NH3 (Ammonia)                        ---->       ")
	color.Yellow("%.2f μg/m3",AqiInfo.List[0].Components.Nh3)
	
}

func helpFunc(){
	allCommands := GetCommands()
	color.HiGreen("------------------------------------------------- All Commands -------------------------------------------------------")
	for _,cmd:=range allCommands {
		fmt.Printf("=> %s   : %s\n",cmd.name,cmd.description)
	}
}

func exitFunc(){
	color.HiRed("=> Exiting...")
	os.Exit(0)
}

