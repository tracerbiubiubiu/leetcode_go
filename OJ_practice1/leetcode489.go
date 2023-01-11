package main

/*
房间（用格栅表示）中有一个扫地机器人。格栅中的每一个格子有空和障碍物两种可能。

扫地机器人提供4个API，可以向前进，向左转或者向右转。每次转弯90度。

当扫地机器人试图进入障碍物格子时，它的碰撞传感器会探测出障碍物，使它停留在原地。

请利用提供的4个API编写让机器人清理整个房间的算法。

interface Robot {
  // 若下一个方格为空，则返回true，并移动至该方格
  // 若下一个方格为障碍物，则返回false，并停留在原地
  boolean move();

  // 在调用turnLeft/turnRight后机器人会停留在原位置
  // 每次转弯90度
  void turnLeft();
  void turnRight();

  // 清理所在方格
  void clean();
}
*/
/**
 * // This is the robot's control interface.
 * // You should not implement it, or speculate about its implementation
 * type Robot struct {
 * }
 *
 * // Returns true if the cell in front is open and robot moves into the cell.
 * // Returns false if the cell in front is blocked and robot stays in the current cell.
 * func (robot *Robot) Move() bool {}
 *
 * // Robot will stay in the same cell after calling TurnLeft/TurnRight.
 * // Each turn will be 90 degrees.
 * func (robot *Robot) TurnLeft() {}
 * func (robot *Robot) TurnRight() {}
 *
 * // Clean the current cell.
 * func (robot *Robot) Clean() {}
 */

type Robot struct {
}

func (robot *Robot) TurnLeft()  {}
func (robot *Robot) TurnRight() {}
func (robot *Robot) Clean()     {}
func (robot *Robot) Move() bool {
    return false
}

var direction = [][]int{
    {0, 1},  //up
    {0, -1}, //down
    {-1, 0}, //left
    {1, 0},  //right
}
var visited = make(map[int]map[int]bool)

func cleanRoom(robot *Robot) {

}

//没路的场景下回退至前一格
func goBack(robot *Robot) {
    robot.TurnRight()
    robot.TurnRight()
    robot.Move()
    robot.TurnRight()
    robot.TurnRight()
}
func visit(i, l int, dir []int) {
    if _, ok := visited[i][l]; !ok {
        visited[i] = make(map[int]bool)
    }
    visited[i][l] = true
}
