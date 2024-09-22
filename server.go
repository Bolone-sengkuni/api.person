package main

import (
	"github.com/api.person/generate"
	"github.com/valyala/fasthttp"
	"log"
	"github.com/goccy/go-json"
)


func main()  {
	requestHandler := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
			case "/": 
				data := map[string]interface{}{
					"info": map[string]interface{}{
						"api_images": "localhost:8989/images",
						"api_data_orang": "localhost:8989/data_orang",
					},
				}
				jsonData, _ := json.Marshal(data)
				ctx.Response.Header.SetContentType("application/json")
				ctx.Response.SetStatusCode(fasthttp.StatusOK)
				ctx.Response.SetBody(jsonData)
			case "/images":
				log.Println("call localhost:8989: generate images")
				image := generate.GenerateImages()
				ctx.Response.Header.SetContentType("image/jpeg")
				ctx.Response.SetStatusCode(fasthttp.StatusOK)
				ctx.Response.SetBody(image)
			case "/data_orang": 
				result := generate.GetDataAll()
				jsonData, _ := json.Marshal(result)
				ctx.Response.Header.SetContentType("application/json")
				ctx.Response.SetStatusCode(fasthttp.StatusOK)
				ctx.Response.SetBody(jsonData)
			default:
				ctx.Error("Unsupported path", fasthttp.StatusNotFound)
				ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		}
	}


	addr := ":8989" 
	log.Printf("Starting server on %s\n", addr)
	if err := fasthttp.ListenAndServe(addr, requestHandler); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}










}