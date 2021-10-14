package main

import (
	"fmt"
)

type Response struct {
	Success bool   `json:"success"`
	Rate    string `json:"rate"`
	Amount  string `json:"amount"`
	Price   string `json:"price"`
}


func main(){
	//JPY:=0.00487
	//res,err:=getNowJPY(JPY)
	//if err !=nil{
	//	log.Println(err)
	//	return
	//}

	//v := fmt.Sprintf("%.5f",JPY)
	//resStr := fmt.Sprintf("%.2f",res)
	//fmt.Println(v+"="+resStr+"円")

}

func f() {
	fmt.Println("関数fの処理")
}

func getNowJPY(buyBTC float64)(Response,error){
	var response Response
	//var nowJPY float64
	var err error
	//resp, _ := http.Get("https://coincheck.com/api/exchange/orders/rate/?order_type=sell&pair=btc_jpy&amount="+strconv.FormatFloat(buyBTC, 'f', 4, 64))
	//defer resp.Body.Close()
	//byteArray, _ := ioutil.ReadAll(resp.Body)
	//err = json.Unmarshal(byteArray, &response)
	//if err !=nil{
	//	log.Println(err)
	//	return nowJPY, err
	//}
	//nowJPY,err=strconv.ParseFloat(response.Price, 64)
	//if err !=nil{
	//	log.Println(err)
	//	return nowJPY, err
	//}
	response.Rate="2387638976"
	response.Amount="1"
	response.Price="22"
	response.Success=true
	return response,err
}