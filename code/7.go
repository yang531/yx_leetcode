// 整数反转
package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	res:=0
	for x!=0 {
		res=res*10+x%10
		if(res<math.MinInt32||res>math.MaxInt32){
			return 0
		}
		x=x/10
	}
	return res
}

func main()  {
	res := reverse(-7800)

	fmt.Println(res)
}

