/* Complete the function solveMeFirst to compute the sum of two integers.

Example
a = 7
b = 3

Return 10.

Function Description

Complete the solveMeFirst function in the editor below.

solveMeFirst has the following parameters:

int a: the first value
int b: the second value
Returns
- int: the sum of a and b

Constraints
1 =< a,b =< 1000

Sample Input

a = 2
b = 3
Sample Output

5
Explanation
2 + 3 = 5
*/

package main

import "fmt"

func solveMeFirst(a uint32, b uint32) uint32 {
	if 1 <= a && b <= 1000 {
		return (a + b)
	}
	return 0
}

func main() {
	var a, b, res uint32
	fmt.Scanf("%v\n%v", &a, &b)
	res = solveMeFirst(a, b)
	fmt.Println(res)
}
