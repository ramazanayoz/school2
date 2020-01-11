package main
import ("fmt"
"regexp"
"bufio"
"os"
"strconv"
"strings"

)
func kdv(number int) float64{

	f := float64(number)
	f = f*0.18 + f
	return f
}

var tryAgain = 1  //global

func main() {


	for tryAgain == 1 {
		tryNumber :=0
		tryNumber = tryNumber +1

		if (tryNumber ==400){
			fmt.Println("you try 400 once")
			tryAgain = 0
			break
		}




		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')


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






				if  "850"==secondC+thirdC+fourthC ||"855"==secondC+thirdC+fourthC {
					fmt.Println("Ön prefix ile simetrik numaralardan 1000")
					fmt.Println(kdv(1000))

					tryAgain =0
					break

				}


				if firstC == secondC && secondC == thirdC && thirdC == fourthC {


					fmt.Println("Son 4 rakamı aynı olan numaralardan 1000")


					tryAgain =0
					break

				}

				if  secondC == thirdC &&  thirdC == fourthC && secondC != firstC {

					fmt.Println("Son 3 rakamı aynı olan numaralardan 500")
					tryAgain =0
					break


				}


				if firstC + secondC == thirdC + fourthC   {
					if firstC == secondC && secondC == thirdC && thirdC == fourthC {
						temp :=0
						if temp ==1{

							fmt.Println("Son 4 rakamı 2’şer simetrik olanlardan 500")
							tryAgain =0
							break

						}
					}
				}

				if (firstC == secondC && secondC != thirdC ) || (thirdC == fourthC && thirdC != secondC) || (secondC == thirdC && thirdC != fourthC && firstC != secondC  )  {

					fmt.Println("Son 4 rakamı 2’şer olarak aynı olanlardan 250 ek")
					tryAgain =0
					break


				}

				if firstNumber <= 6 && firstNumber +3 == secondNumber+2 &&secondNumber+2 == thirdNumber+1 &&thirdNumber+1 ==fourthNumber{
					fmt.Println("Son dört rakamı ardışık olan 100 ek")
					tryAgain =0
					break
				}


				if secondNumber==0 && thirdNumber<=10 && fourthNumber ==0 {
					fmt.Println(" Son 3 rakamı 010, 020, 030, 040 gibi artanlardan 100 ek")
				}




				//fmt.Println(text)

				//int çevirme
				//n, err := strconv.ParseInt(text, 10, 64)
				//if err == nil {
				//	fmt.Printf("%d of type %T", n, n)
				//}

				//        fmt.Println(text)

				fmt.Println("numaranız uygundur")
				tryAgain = 0

				break



			}
			if !match {
				fmt.Println("Lütfen 90 850 885 0000 ile  90 850 885 3999 arasında bir sayı giriniz")

			}


		}

		if match {
			fmt.Println("Lütfen geçerli bir sayı giriniz")

		}
	}
}


