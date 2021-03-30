package stats
import  (
	"github.com/anonimous-arn/bank/pkg/bank/types"
	"fmt"
	 
)



func ExampleAvg() {

	res := Avg([]types.Payment{
		{ID: 1, Amount: 10},
		{ID: 2, Amount: 20},
		{ID: 3, Amount: 30},
	})
	fmt.Println(res)
	//Output:
	//20
}

func ExampleTotalInCategory() {

	res := TotalInCategory([]types.Payment{
		{ID: 1, Amount: 10, Category: "internet"},
		{ID: 2, Amount: 20, Category: "cafe"},
		{ID: 3, Amount: 30, Category: "cafe"},
	}, "cafe")
	fmt.Println(res)
	//Output:
	//50
}