# Go Daily Algorithms

每日算法练习仓库，使用 Go 实现。

## 题目列表

### Day 01: 两数之和
- 问题描述：在数组中找到两个数使它们的和等于目标值
- 解法：哈希表
- 运行示例：
  ```bash
  cd examples
  go run day01_main.go
  ```

# Daily Coding Challenges

## 题目列表

### Day 02: 找出单独的数字
- **问题描述**：给定一个非空整数数组，其中某个元素只出现一次，其余每个元素均出现两次，找出那个唯一出现的数字
- **解法**：哈希表计数法
- **时间复杂度**：O(n)
- **空间复杂度**：O(n)
- **运行示例**：
  ```bash
  cd day02_single_number
  go test -v  # 运行测试用例
