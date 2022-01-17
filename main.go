package main

import (
        "fmt"
        "os"
        "github.com/rs/zerolog"
        "github.com/rs/zerolog/log"
        "github.com/valyala/fasthttp"
)

func main() {
//	cfg := config.GetAPIConfig()
//	client := client.NewTradeMade(cfg.ClientURI)
//	ctrl := controller.NewController(client)
        cfg := Conf{}  
        client := &fasthttp.Client{
		ReadTimeout: 10000000,
	} 
        zerolog.SetGlobalLevel(zerolog.InfoLevel) 
       
        zerolog.SetGlobalLevel(zerolog.DebugLevel)
	
        fmt.Println("Test")
        
         
	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)
         
        logger := zerolog.New(os.Stderr).With().Timestamp().Logger()


//        logger := log.Logger.New() 
            
        cli := NewCli(&cfg,client,&logger)
       
        if cli != nil {
           fmt.Println("Tested")  
        }

      decimalvar := 1

      resp, _, _, _ := cli.GetQuotes("1min", "NASDAQ", "", "", "Common Stock", "", "America/New_York", decimalvar, []string{"AAPL"})

      fmt.Println(resp)
}
