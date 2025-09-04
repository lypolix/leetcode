package main 

func main() {
	generate(5)
}

func generate(numRows int) [][]int {
    pt := make([][]int, numRows)
    pt[0] = []int{1}

    lenArray := 1
    for i := 1; i < numRows; i ++ {
        lenArray += 1
        num := 0
        array := make ([]int, 0)
        for j := 0; j < lenArray; j++ {
            if j - 1 < 0 {
                num = pt[i - 1][0]
            }else if j == lenArray - 1{
                num = pt[i -1][lenArray - 2]
            }else {
                num = pt[i - 1][j - 1] + pt[i - 1][j]
            }
            array = append(array, num)
        }
        pt[i] = array
    }
	return pt
}
