package main
import "math"

//812. 最大三角形面积
func largestTriangleArea(points [][]int) float64 {
    var result float64
    for i:=0; i<len(points); i++{
        x1, y1 := points[i][0], points[i][1]
        for j:=i+1; j<len(points); j++{
            x2, y2 := points[j][0], points[j][1]
            for k:=j+1; k<len(points); k++{
                x3, y3 := points[k][0], points[k][1]
                //求三条边长
                a := math.Sqrt(float64((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)))
                b := math.Sqrt(float64((x2-x3)*(x2-x3) + (y2-y3)*(y2-y3)))
                c := math.Sqrt(float64((x3-x1)*(x3-x1) + (y3-y1)*(y3-y1)))
                //海伦公式求面积
                p := (a+b+c)/2
                area := math.Sqrt(float64(p*(p-a)*(p-b)*(p-c)))
                if area > result{
                    result = area
                }
            }
        }
    }
    return result
}
