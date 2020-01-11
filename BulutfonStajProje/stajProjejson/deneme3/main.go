package main

import ("net/http"
	"html/template"
	"fmt"
	//	"bufio"
	//	"os"
	"regexp"
	strconv "strconv"
	"strings"
	"encoding/json"

)
var result = "result" //global
var price int
var priceKDV int

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))


}


func main() {
	http.HandleFunc("/",index)
	http.HandleFunc("/process",processor)

	http.ListenAndServe(":3002",nil)


}



func index(w http.ResponseWriter, r*http.Request){
	tpl.ExecuteTemplate(w, "index.html",nil)

}

func  processor(w http.ResponseWriter, r *http.Request)  {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	fname := r.FormValue("firster")
	//lname := r.FormValue("laster")
	fnumber := "+90 850 885 "+fname


	//priceS := strconv.FormatInt(int64(price), 10)
	//priceKDVS := strconv.FormatInt(int64(priceKDV), 10)


	phonoNumber(fname)
	type Profile struct{
		Control []string
		Price int
		TotalPrice int

	}

	profile  := Profile{[]string{fnumber +" "+ result}, price, priceKDV}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)


}




func phonoNumber  (phoneNumber string)  {
	var tryAgain = 1  //global

	for tryAgain == 1 {
		tryNumber :=0
		tryNumber = tryNumber +1

		if (tryNumber ==400){
			result="you try 400 once"
			tryAgain = 0
			break
		}




		//	reader := bufio.NewReader(os.Stdin)
		//	fmt.Print("Enter text: ")
		//	text, _ := reader.ReadString('\n')


		text := phoneNumber


		match, _ := regexp.MatchString("([a-z]+)", text)
		if  !match {

			match, _ := regexp.MatchString("([0-3]+[0-9]+[0-9]+[0-9$])",text)
			//  fmt.Println(match) match kullanmazsak hata verir match tanımladıktan sonra

			// text :="90850885"+text


			if match {



				//acess char
				firstC := (string([]rune(text)[0])) // UTF-8
				secondC :=  (string([]rune(text)[1]))
				thirdC := (string([]rune(text)[2]))
				fourthC := (string([]rune(text)[3]))

				//int çevirme
				allNumber, _ := strconv.ParseInt(strings.TrimSpace(text), 10, 64)

				firstNumber, _ := strconv.ParseInt(strings.TrimSpace(firstC), 10, 64)
				secondNumber, _ := strconv.ParseInt(strings.TrimSpace(secondC), 10, 64)
				thirdNumber, _ := strconv.ParseInt(strings.TrimSpace(thirdC), 10, 64)
				fourthNumber, _ := strconv.ParseInt(strings.TrimSpace(fourthC), 10, 64)
				//firstTwoNumber, _ := strconv.ParseInt(strings.TrimSpace(firstC+secondC), 10, 64)
				//lastTwoNumber, _ := strconv.ParseInt(strings.TrimSpace(firstC+secondC), 10, 64)
				//lastThirdNumber, _ := strconv.ParseInt(strings.TrimSpace(secondC+thirdC+fourthC), 10, 64)

				fmt.Println(allNumber )






				if  "855"==secondC+thirdC+fourthC {
					result = "Ön prefix ile simetrik numaralardan"
					price = 1000
					priceKDV =kdv(price)


					tryAgain =0
					break

				}


				if firstC == secondC && secondC == thirdC && thirdC == fourthC {


					result = "Son 4 rakamı aynı olan numaralardan"


					price = 1000
					priceKDV =kdv(price)

					break

				}

				if  secondC == thirdC &&  thirdC == fourthC && secondC != firstC {

					result = "Son 3 rakamı aynı olan numaralardan  "
					price = 500
					priceKDV =kdv(price)

					tryAgain =0
					break


				}


				if firstC + secondC == thirdC + fourthC   {
					if firstC == secondC && secondC == thirdC && thirdC == fourthC {
						temp :=0
						if temp ==1{

							result = "Son 4 rakamı 2’şer simetrik olanlardan"
							price = 500
							priceKDV =kdv(price)

							tryAgain =0
							break

						}
					}
				}

				if (firstC == secondC && secondC != thirdC ) || (thirdC == fourthC && thirdC != secondC) || (secondC == thirdC && thirdC != fourthC && firstC != secondC  )  {

					fmt.Println("Son 4 rakamı 2’şer olarak aynı olanlardan")
					price = 250
					priceKDV =kdv(price)


					tryAgain =0
					break


				}

				if firstNumber <= 6 && firstNumber +3 == secondNumber+2 &&secondNumber+2 == thirdNumber+1 &&thirdNumber+1 ==fourthNumber{
					result = "Son dört rakamı ardaşık olan "
					price = 100
					priceKDV =kdv(price)


					tryAgain =0
					break
				}


				if secondNumber==0 && thirdNumber<=10 && fourthNumber ==0 {
					result = "Son 3 rakamı 010, 020, 030, 040 gibi artanlardan"
					price = 100
					priceKDV =kdv(price)


				}

				if(allNumber >= 1000) && (2018 >= allNumber)  {

					allNumber := strconv.FormatInt(int64(allNumber), 10)

					result = allNumber+"Bu numara tarihi bir anlam taşımaktadır"
					price = 50
					priceKDV =kdv(50)
					tryAgain =0
					break

				}




				//fmt.Println(text)

				//int çevirme
				//n, err := strconv.ParseInt(text, 10, 64)
				//if err == nil {
				//	fmt.Printf("%d of type %T", n, n)
				//}

				//        fmt.Println(text)

				result = "özel bir numara değildir numaranız uygundur"

				tryAgain = 0

				break



			}
			if !match {
				result = "Lütfen 90 850 885 0000 ile  90 850 885 3999 arasında bir sayı giriniz"
				tryAgain = 0
				break
			}


		}

		if match {
			result = "Lütfen geçerli bir sayı giriniz"

		}
	}
}



func kdv(number int) int{

	f := float64(number)
	f = f*0.18 + f
	var k int = int(f)

	return k
}