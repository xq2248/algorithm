/*
 * Copyright xq2248
 * mail: xq2248@163.com
 * github: xq2248
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package polygon

import (
	"fmt"
	"math"
)

//From https://blog.csdn.net/DevolperFront/article/details/99688144
//字节跳动公司算法面试题
//问题1：
//有一个n边型（P0，P1，P2)，每一条边结尾垂直或者水平的线段，先给定数字k，
// 以P0为起点将n便携性分成k端，每段的长度相同，打印等分点

//实现要点：算法必须与正确的数据结构相匹配
//问题点：浮点数比较问题
//进一步优化：可以采用三角函数来求中间点，这样每个边不一定需要是垂直或者水平

//Point 定义点
type Point struct {
	X, Y float64
}

//Line 定义边
type Line struct {
	from, to int     //边在Point切片中的位置
	length   float64 //边的长度
	xDirect  float64 //边的水平方向: 取值-1,0,1
	yDirect  float64 //边的垂直平方向: 取值-1,0,1
}

//equalsZero 浮点数比较，注意不能采用==比较运算符
func equalsZero(num float64) bool {
	var eps float64 = 0.00000001
	if math.Abs(num) < eps {
		return true
	}
	return false
}

//getLines 根据数据点切片返回边的数组
func getLines(ps []Point) []Line {
	count := len(ps)
	ls := make([]Line, count)
	for i := 0; i < count; i++ {
		var line Line
		line.from = i
		if i == count-1 {
			line.to = 0 //最后一条边的终点是第1个点
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

//getDirect 返回边的水平方向和垂直方向
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

//distOf 计算两点之间的距离
func distOf(a, b *Point) float64 {
	lenx := math.Abs(a.X - b.X)
	leny := math.Abs(a.Y - b.Y)
	return lenx + leny
}

//totalLen 计算整个多边形的总周长
func totalLen(lines []Line) (total float64) {
	total = 0
	for _, line := range lines {
		total += line.length
	}
	return
}

//getPoint 在指定的边上获取离起始点delta距离的点
func getPoint(line *Line, delta float64, ps []Point) (point Point) {
	src := &ps[line.from]
	point.X = src.X + line.xDirect*delta
	point.Y = src.Y + line.yDirect*delta
	return
}

//Walk 算法入口
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

lineLoop:
	for _, line := range ls {
		lineRes := line.length
		for {
			if lineRes < curRes { //当前边长小于要找的下一点的剩余距离
				curRes -= lineRes
				break
			} else if lineRes == curRes { //当前边长等于下一点的剩余距离，那么下一个点就是to
				lineRes = 0
				newPoint := getPoint(&line, curRes, ps)
				fmt.Println("new point1", newPoint)
				curNo++
				curRes = average
				if curNo == count {
					break lineLoop
				}
			} else { //当前边长很长，下一个点在当前的边上
				lineRes -= curRes
				newPoint := getPoint(&line, curRes, ps)
				fmt.Println("new point2", newPoint)
				curNo++
				curRes = average
				if curNo == count {
					goto lineLoop
				}
			}

		}
	}
}
