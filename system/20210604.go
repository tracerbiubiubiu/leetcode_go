package main

import (
    "fmt"
    "sort"
)

/*
地铁闸机
*/
//俩思路优先队列与两个队列 划窗没想到咋搞
//两种队列 0出 1进

//[]int{下标,时间，方向}
func main() {
    arrTime := []int{0, 0, 1, 5}
    direction := []int{0, 1, 1, 0}

    fmt.Println(getTimes(arrTime, direction)) //[2,0,1,5]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
func buildQueue(arrTime, direction []int) (inQueue [][]int, outQueue [][]int) {
    for i := range arrTime {
        //出 //[]int{下标,到达时间,通过闸机时间}
        if direction[i] == 0 {
            outQueue = append(outQueue, []int{i, arrTime[i], -1})
        } else {
            inQueue = append(inQueue, []int{i, arrTime[i], -1})
        }
    }
    //先把队列排序
    //sort.Slice(inQueue, func(i, j int) bool {
    //    if inQueue[i][1] < inQueue[j][1] {
    //        return true
    //    }
    //    if inQueue[i][0] < inQueue[j][0] {
    //        return true
    //    }
    //    return false
    //})
    fmt.Println(inQueue)
    fmt.Println(outQueue)
    return
}
func getTimes(arrTime, direction []int) []int {
    inQueue, outQueue := buildQueue(arrTime, direction)
    //下次闸机对选择方向的时间
    nextTime := 0

    i, j := 0, 0
    for i < len(inQueue) && j < len(outQueue) {
        //存在乘客先进站的场景 闸机初始默认向进站方向开放
        if inQueue[i][1] <= outQueue[j][1] {
            //出站时间为进站时间与闸机时间最大值
            inQueue[i][2] = max(nextTime, inQueue[i][1])
            nextTime = inQueue[i][2] + 1
            i++
            //开始寻找下一位进站乘客 在下次闸机选择方向前，进站的人全部进站
            for i < len(inQueue) && inQueue[i][1] <= nextTime {
                inQueue[i][2] = max(nextTime, inQueue[i][1])
                nextTime = inQueue[i][2] + 1
                i++
            }
        } else {
            outQueue[j][2] = max(nextTime, outQueue[j][1])
            nextTime = outQueue[j][2] + 1
            j++
            for j < len(outQueue) && outQueue[j][1] <= nextTime {
                outQueue[j][2] = max(nextTime, outQueue[j][1])
                nextTime = outQueue[j][2] + 1
                j++
            }
        }
    }
    //还存在剩余队列
    for i < len(inQueue) {
        inQueue[i][2] = max(nextTime, inQueue[i][1])
        nextTime = inQueue[i][2] + 1
        i++
    }
    for j < len(outQueue) {
        outQueue[j][2] = max(nextTime, outQueue[j][1])
        nextTime = outQueue[j][2] + 1
        j++
    }
    return getResult(inQueue, outQueue)
}
func getResult(inQueue [][]int, outQueue [][]int) []int {
    res := []int{}
    inQueue = append(inQueue, outQueue...)
    sort.Slice(inQueue, func(i, j int) bool {
        return inQueue[i][0] < inQueue[j][0]
    })
    for i := range inQueue {
        res = append(res, inQueue[i][2])
    }
    return res
}

/*


func main() {
   arrTime := []int{2, 2, 2, 2, 3, 3, 5, 5, 20, 20}
   direction := []int{0, 1, 1, 0, 0, 1, 1, 0, 0, 1}

   fmt.Println(getTimes(arrTime, direction))
}

func getQueue(arrTimes []int, direction []int) ([][]int, [][]int) {
   // 进站乘客 {到达时间，下标，通过闸机时刻}
   var inQueue [][]int
   // 出站乘客 {达到时间，下标，通过闸机时刻}
   var outQueue [][]int
   for i, val := range arrTimes {
      if direction[i] == 0 {
         // 出栈
         outQueue = append(outQueue, []int{val, i, 0})
      } else {
         // 进栈
         inQueue = append(inQueue, []int{val, i, 0})
      }
   }
   return inQueue, outQueue
}

func getTimes(arrTimes []int, direction []int) []int {
   inQueue, outQueue := getQueue(arrTimes, direction)
   // 初始化闸机时刻
   nextTime := arrTimes[0]
   // 初始化进站、出站指针
   in, out := 0, 0
   for in < len(inQueue) && out < len(outQueue) {
      if inQueue[in][0] <= outQueue[out][0] {
         inQueue[in][2] = max(nextTime, inQueue[in][0])
         nextTime = inQueue[in][2] + 1
         in++

         // 下一位乘客陆续进站
         for in < len(inQueue) && inQueue[in][0] <= nextTime {
            inQueue[in][2] = max(nextTime, inQueue[in][0])
            nextTime = inQueue[in][2] + 1
            in++
         }
      } else {
         outQueue[out][2] = max(nextTime, outQueue[out][0])
         nextTime = outQueue[out][2] + 1
         out++

         for out < len(outQueue) && outQueue[out][0] <= nextTime {
            outQueue[out][2] = max(nextTime, outQueue[out][0])
            nextTime = outQueue[out][2] + 1
            out++
         }
      }
   }

   // 处理剩余队列
   for in < len(inQueue) {
      inQueue[in][2] = max(nextTime, inQueue[in][0])
      nextTime = inQueue[in][2] + 1
      in++
   }

   for out < len(outQueue) {
      outQueue[out][2] = max(nextTime, outQueue[out][0])
      nextTime = outQueue[out][2] + 1
      out++
   }
   return getResult(inQueue, outQueue)
}

func max(a, b int) int {
   if a > b {
      return a
   } else {
      return b
   }
}

func getResult(inQueue, outQueue [][]int) []int {
   inQueue = append(inQueue, outQueue...)
   sort.Slice(inQueue, func(i, j int) bool {
      return inQueue[i][1] < inQueue[j][1]
   })
   var ans []int
   for _, val := range inQueue {
      ans = append(ans, val[2])
   }
   return ans
}
*/
