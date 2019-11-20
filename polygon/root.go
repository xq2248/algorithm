package polygon

//作者：大熊
//mail: xq2248@163.com

import (
	"fmt"
	"math"
)

//From https://blog.csdn.net/DevolperFront/article/details/99688144
//问题1：
//有一个n边型（P0，P1，P2)，每一条边结尾垂直或者水平的线段，先给定数字k，
// 以P0为起点将n便携性分成k端，每段的长度相同，打印等分点

//实现要点：算法必须与正确的数据结构相匹配
//问题点：浮点数比较问题
//进一步优化：可以采用三角函数来求中间点，这样每个边不一定需要是垂直或者水平
type Point struct {
	X, Y float64
}

type Line struct {
	from, to         int     //边在Point切片中的位置
	length           float64 //边的长度
	xDirect  float64         //边的水平方向: 取值-1,0,1
	yDirect  float64         //边的垂直平方向: 取值-1,0,1
}

func equalsZero(num float64) bool {
	var eps float64 = 0.00000001
	if math.Abs(num) < eps {
		return true
	}
	return false
}

func getLines(ps []Point) []Line {
	count := len(ps)
	ls := make([]Line, count)
	for i := 0; i < count; i++ {
		var line Line
		line.from = i
		if i == count-1 {
			line.to = 0
		} else {
			line.to = i + 1
		}
		src := &ps[line.from]
		dst := &ps[line.to]
		line.length = distOf(src, dst)

		getDirect(dst, src, &line)

		ls = append(ls, line)
	}
	return ls
}

func getDirect(dst *Point, src *Point, line *Line) {
	xDirect := dst.X - src.X
	yDirect := dst.Y - src.Y
	xBool := equalsZero(xDirect)
	yBool := equalsZero(yDirect)
	fmt.Println(*src, *dst, xDirect, yDirect)
	if xBool == true && yBool == false {
		line.xDirect = 0
		if yDirect > 0 {
			line.yDirect = 1
		} else {
			line.yDirect = -1
		}
	} else if xBool == false && yBool == true {
		line.yDirect = 0
		if xDirect > 0 {
			line.xDirect = 1
		} else {
			line.xDirect = -1
		}
	} else {
		fmt.Println("bad line:", *src, *dst, xDirect, yDirect)
		panic("bad line ")
	}
}

//计算两点之间的距离
func distOf(a, b *Point) float64 {
	lenx := math.Abs(a.X - b.X)
	leny := math.Abs(a.Y - b.Y)
	return lenx + leny
}

//计算总周长
func totalLen(lines []Line) (total float64) {
	total = 0
	for _, line := range lines {
		total += line.length
	}
	return
}

func getPoint(line *Line, delta float64, ps []Point) (point Point) {
	src := &ps[line.from]
	point.X = src.X + line.xDirect*delta
	point.Y = src.Y + line.yDirect*delta
	return
}

func Walk(ps []Point, k int) {
	if k < 0 {
		panic("bad division")
	}
	count := len(ps)
	if count < 2 {
		panic("too little points")
	}

	ls := getLines(ps)
	average := totalLen(ls) / float64(k)

	var curNo int = 0            //要找的当前点序号
	var curRes float64 = average //要找的当前点剩余长度

	for _, line := range ls {
		lineRes := line.length
		for {
			if lineRes < curRes {
				curRes -= lineRes
				goto NextLine
			} else if lineRes == curRes {
				lineRes = 0
				newPoint := getPoint(&line, curRes, ps)
				fmt.Println("new point1", newPoint)
				curNo++
				curRes = average
				if curNo == count {
					goto Finish
				}
			} else {
				lineRes -= curRes
				newPoint := getPoint(&line, curRes, ps)
				fmt.Println("new point2", newPoint)
				curNo++
				curRes = average
				if curNo == count {
					goto Finish
				}
			}

		}
	NextLine:
	}
Finish:
}
