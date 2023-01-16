package main

import (
    "fmt"
    "sort"
)

/*
自动售货机进出货管理系统

*/
/*

type vendingMachine struct {
   val []productor
   lens int
   perCal int
}
type productor struct {
   id int
   productIdList []int
}

func VendingMachineSystem(trayNum int, trayCapacity int) vendingMachine{
   v := vendingMachine{
      val : make([]productor, 0),
      lens: trayNum,
      perCal: trayCapacity,
   }
   return v
}

func (v *vendingMachine) addProduct( brandId int,  productIdList []int) bool{
   if v.lens == 0{
      return false
   }
   flag := false
   for i:=0 ;i < len(v.val); i++{
      if v.val[i].id == brandId{
         flag = true
         curLens := len(v.val[i].productIdList)
         if v.perCal -  curLens >= len(productIdList){
            for j:=0; j<len(productIdList); j++{
               v.val[i].productIdList = append(v.val[i].productIdList, productIdList[j])
            }
            return true
         }
         return false
      }
   }


   if !flag && len(v.val) <= v.lens && v.perCal >= len(productIdList){
      if len(v.val) == v.lens {
         for i:=0;i<len(v.val);i++{
            if len(v.val[i].productIdList) == 0{
               v.val[i].id = brandId
               v.val[i].productIdList = productIdList
               return true
            }
         }
         return false
      }

      v.val = append(v.val, productor{id:brandId, productIdList:productIdList})
      return true
   }



   return false

}

func (v *vendingMachine) buyProduct(brandId, num int) []int {

   for i:=0 ;i < len(v.val); i++{
      if v.val[i].id == brandId{
         curLens := len(v.val[i].productIdList)
         if curLens >= num{
            temp := v.val[i].productIdList[:num]
            v.val[i].productIdList = v.val[i].productIdList[num:]
            return temp
         }
         return nil
      }
   }
   return nil

}

func (v *vendingMachine) queryProduct() []int {


   if len(v.val) == 0{
      return nil
   }

   temp := make([][]int, 0)
   for i:=0;i<len(v.val);i++{
      if len(v.val[i].productIdList) > 0{
         temp = append(temp, []int{v.val[i].id, v.val[i].productIdList[0]})
      }
   }
   sort.Slice(temp, func(i, j int) bool {
      return temp[i][0] < temp[j][0]
   })
   res := make([]int, 0)
   for i:=0;i<len(temp);i++{
         res = append(res, temp[i][1])

   }
   return res

}

func main() {
   //obj := VendingMachineSystem(2,5)
   //a1:= obj.addProduct(3,[]int{3,5,4,6,2})
   //fmt.Println(a1)
   //a2:=obj.buyProduct(3,3)
   //fmt.Println(a2)
   //a3:= obj.queryProduct()
   //fmt.Println(a3)
   //
   //["VendingMachineSystem","addProduct","addProduct","addProduct","addProduct","buyProduct","buyProduct","buyProduct",
   //"queryProduct","addProduct","addProduct","buyProduct","addProduct"]
   //[[2,5],[3,[3,5,4,6,2,1]],[3,[3,5,4,6,2]],[1,[9]],[2,[7]],[3,3],[9,3],[3,3],
   //[],[3,[10,20,13,14]],[3,[10,20,13]],[3,5],[9,[5,4]]]
   //[null,false,true,true,false,[3,5,4],[],[],[9,6],false,true,[6,2,10,20,13],true]

   obj := VendingMachineSystem(2,5)
   a1:= obj.addProduct(3,[]int{3,5,4,6,2,1})
   fmt.Println(a1)
   a2:= obj.addProduct(3,[]int{3,5,4,6,2})
   fmt.Println(a2)
   a3:= obj.addProduct(1,[]int{9})
   fmt.Println(a3)
   a4:= obj.addProduct(2,[]int{7})
   fmt.Println(a4)
   a5:=obj.buyProduct(3,3)
   fmt.Println(a5)
   a6:=obj.buyProduct(9,3)
   fmt.Println(a6)
   a7:=obj.buyProduct(3,3)
   fmt.Println(a7)
   a8:= obj.queryProduct()
   fmt.Println(a8)
   a10:= obj.addProduct(3,[]int{10,20,13,14})
   fmt.Println(a10)
   a11:= obj.addProduct(3,[]int{10,20,13})
   fmt.Println(a11)
   a12:=obj.buyProduct(3,5)
   fmt.Println(a12)
   a13:= obj.addProduct(9,[]int{5,4})
   fmt.Println(a13)
}
*/
type VendingMachineSystem struct {
    trayNum      int       //商品轨道
    trayCapacity int       //轨道容量
    tray         []product //轨道上具体信息
}
type product struct {
    brandID int //品牌
    //num       int   //所在轨道
    productId []int //商品编号列表
}

func newVendingMachineSystem(trayNum, trayCapacity int) VendingMachineSystem {
    //初始化没有使用的轨道 此时轨道cap存在，但是没有len 代表最大长度为cap，上面的货物为0
    tray := make([]product, 0, trayNum)
    return VendingMachineSystem{
        tray:         tray,
        trayCapacity: trayCapacity,
        trayNum:      trayNum,
    }
}

func (v *VendingMachineSystem) addProduct(brandId int, productIdList []int) bool {
    for i := range v.tray {

        //存在对应轨道
        if v.tray[i].brandID == brandId {
            //剩余长度充足
            if v.trayCapacity-len(v.tray[i].productId) >= len(productIdList) {
                v.tray[i].productId = append(v.tray[i].productId, productIdList...)
                return true
            }
        }

    }
    //不存在时判断是否有空轨道
    if len(v.tray) < v.trayNum {
        //判断空间是否充足
        if len(productIdList) <= v.trayCapacity {
            v.tray = append(v.tray, product{
                brandID:   brandId,
                productId: productIdList,
            })
            return true
        }
    }
    return false
}

func (v *VendingMachineSystem) buyProduct(brandId, num int) []int {
    for i := range v.tray {
        //判断有没有这个种类的商品
        if v.tray[i].brandID == brandId {
            if len(v.tray[i].productId) >= num {
                tmp := v.tray[i].productId[:num]
                v.tray[i].productId = v.tray[i].productId[num:]
                //没货的时候重新置空
                if len(v.tray[i].productId) == 0 {
                    v.tray = append(v.tray[:i], v.tray[i+1:]...)
                }
                return tmp
            }
        }
    }
    return []int{}
}

func (v *VendingMachineSystem) queryProduct() []int {
    res := []int{}
    if len(v.tray) == 0 {
        return res
    }
    //按品牌id排序
    sort.Slice(v.tray, func(i, j int) bool {
        return v.tray[i].brandID < v.tray[j].brandID
    })
    for i := range v.tray {
        res = append(res, v.tray[i].productId[0])
    }
    return res
}
func main() {
    //obj := newVendingMachineSystem(2, 5)
    //fmt.Println(obj)
    //a1 := obj.addProduct(3, []int{3, 5, 4, 6, 2})
    //fmt.Println(a1)
    //a2 := obj.buyProduct(3, 3)
    //fmt.Println(a2)
    //a3 := obj.queryProduct()
    //fmt.Println(a3)
    obj := newVendingMachineSystem(2, 5)
    a1 := obj.addProduct(3, []int{3, 5, 4, 6, 2, 1})
    fmt.Println(a1)
    a2 := obj.addProduct(3, []int{3, 5, 4, 6, 2})
    fmt.Println(a2)
    a3 := obj.addProduct(1, []int{9})
    fmt.Println(a3)
    a4 := obj.addProduct(2, []int{7})
    fmt.Println(a4)
    a5 := obj.buyProduct(3, 3)
    fmt.Println(a5)
    a6 := obj.buyProduct(9, 3)
    fmt.Println(a6)
    a7 := obj.buyProduct(3, 3)
    fmt.Println(a7)
    a8 := obj.queryProduct()
    fmt.Println(a8)
    a10 := obj.addProduct(3, []int{10, 20, 13, 14})
    fmt.Println(a10)
    a11 := obj.addProduct(3, []int{10, 20, 13})
    fmt.Println(a11)
    a12 := obj.buyProduct(3, 5)
    fmt.Println(a12)
    a13 := obj.addProduct(9, []int{5, 4})
    fmt.Println(a13)
}
