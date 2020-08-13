package main





import "fmt"

var(
coins = 50
users = []string{
"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
}
distribution = make(map[string]int, len(users))
)

func main() {
	left := dispatchCoin(users)
	fmt.Println("剩下：", left)
}
func dispatchCoin(s []string) (left int) {
	for _, v := range s {
		coin:=0
		for _, a := range v {
			switch a {
			case 'e', 'E': coin += 1
			case 'i', 'I': coin += 2
			case 'o', 'O': coin += 3
			case 'u', 'U': coin += 4
			}
		}
		distribution[v]=coin
		coins-=coin
	}
	fmt.Print(distribution)
	return coins
}
