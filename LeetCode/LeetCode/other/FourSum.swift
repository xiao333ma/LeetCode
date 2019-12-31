//
//  FourSum.swift
//  LeetCode_Swift
//
//  Created by xiaoma on 5/23/17.
//  Copyright © 2017 xiaoma. All rights reserved.
//

import Cocoa

class FourSum: NSObject {
    func fourSum(_ nums: [Int], _ target: Int) -> [[Int]] {
        var res = [[Int]]()

        guard nums.count >= 4 else {
            return res
        }
        var sortedNums = nums.sorted(by: >);
        for i in 0 ..< sortedNums.count - 3 {
            let numI = sortedNums[i]
            if i == 0 || numI != sortedNums[i - 1] {
                
                
                let maxSum = sortedNums[i] + sortedNums[i + 1] + sortedNums[i + 2] + sortedNums[i + 3];
                if maxSum < target {
                    break
                } else if (maxSum == target) {
                    res .append([sortedNums[i], sortedNums[i + 1], sortedNums[i + 2], sortedNums[i + 3]])
                    break
                }
                
                let minSum = sortedNums[sortedNums.count - 1] + sortedNums[sortedNums.count - 2] + sortedNums[sortedNums.count - 3] + sortedNums[sortedNums.count - 4];

                if minSum > target {
                    continue
                } else if (maxSum == target) {
                    res .append([sortedNums[sortedNums.count - 1], sortedNums[sortedNums.count - 2], sortedNums[sortedNums.count - 3], sortedNums[sortedNums.count - 4]])
                    continue
                }
                
                
                
                for j in i + 1 ..< sortedNums.count - 2 {
                    let numJ = sortedNums[j]
                    if  j == i + 1 || numJ != sortedNums[j - 1] {
                        
                        
                        let maxSum = numI + numJ + sortedNums[j + 1] + sortedNums[j + 2];
                        if maxSum < target {
                            break
                        } else if (maxSum == target) {
                            res .append([numI, numJ, sortedNums[j + 1], sortedNums[j + 2]])
                            break
                        }
                        
                        let minSum = numI + numJ + sortedNums[sortedNums.count - 1] + sortedNums[sortedNums.count - 2];
                        
                        if minSum > target {
                            continue
                        } else if (maxSum == target) {
                            res .append([numI, numJ, sortedNums[sortedNums.count - 1], sortedNums[sortedNums.count - 2]])
                            continue
                        }
                        
                        
                        
                        
                        
                        
                        var left = j + 1
                        var right = sortedNums.count - 1
                        while left < right {
                            
                            let numLeft = sortedNums[left]
                            let numRight = sortedNums[right]
                            
                            let sum = numI + numJ + numLeft + numRight
                            if sum == target {
                                res .append([numI, numJ, numLeft, numRight])
                                repeat {
                                    left += 1
                                } while left < right && sortedNums[left] == numLeft
                                repeat {
                                    right -= 1
                                } while left < right && sortedNums[right] == numRight
                            } else if (sum > target) {
                                repeat {
                                    left += 1
                                } while left < right && sortedNums[left] == numLeft
                            } else {
                                repeat {
                                    right -= 1
                                } while left < right && sortedNums[right] == numRight
                            }
                        }
                    }
                    
                }
            }
            
        }
        return res
    }
}


/*
 struct Component {
	var index1 : Int
	var index2 : Int
	init(_ i1: Int, _ i2: Int) {
 self.index1 = i1
 self.index2 = i2
	}
 }
 
 infix operator  ⊖ { associativity left precedence 140 }
 func ⊖(left: Component, right: Component) -> Bool {
	let set = Set(arrayLiteral:left.index1, left.index2, right.index1, right.index2)
	return set.count == 4
 }
 
 class Solution {
	func fourSum(_ nums: [Int], _ target: Int) -> [[Int]] {
 guard nums.count > 0 else {
 return []
 }
 var suitDict = [ Int : [Component] ]()
 for i in 0..<(nums.count - 1) {
 for j in (i + 1)..<nums.count {
 let add = nums[i] + nums[j]
 if suitDict[add] == nil {
 suitDict[add] = [Component(i, j)]
 }else {
 suitDict[add]!.append(Component(i, j))
 }
 }
 }
 var result = [[Int]]()
 var alreadyMatchList = [Int]()
 for dict in suitDict {
 guard !alreadyMatchList.contains(dict.key) else {
 continue
 }
 alreadyMatchList.append(dict.key)
 let num2 = target - dict.key
 if let suitArray = suitDict[num2] {
 // 匹配上了
 alreadyMatchList.append(num2)
 for c1 in dict.value {
 for c2 in suitArray {
 if c1 ⊖ c2 {
 let r = [nums[c1.index1], nums[c1.index2], nums[c2.index1], nums[c2.index2]].sorted()
 var flag = false
 for index in 0..<result.count {
 if result[index] == r {
 flag = true
 break
 }
 }
 if !flag {
 result.append(r)
 }
 }
 }
 }
 }
 }
 return result
	}
 }
 */
class Solution {
    func fourSum(_ nums: [Int], _ target: Int) -> [[Int]] {
        guard nums.count >= 4 else {
            return []
        }
        
        var sortedNums = nums.sorted(by:{$0 < $1})
        var result : [[Int]] = []
        
        for i in 0 ... sortedNums.count - 4 {
            if i > 0 && sortedNums[i] == sortedNums[i - 1] {
                continue
            }
            
            // pruning
            let minSum = sortedNums[i] + sortedNums[i + 1] + sortedNums[i + 2] + sortedNums[i + 3]
            if minSum > target {
                break
            } else if minSum == target {
                result.append([sortedNums[i], sortedNums[i + 1], sortedNums[i + 2], sortedNums[i + 3]])
                break
            }
            
            let maxSum = sortedNums[i] + sortedNums[sortedNums.count - 1] + sortedNums[sortedNums.count - 2] + sortedNums[sortedNums.count - 3]
            if maxSum < target {
                continue
            } else if maxSum == target {
                result.append([sortedNums[i], sortedNums[sortedNums.count - 1], sortedNums[sortedNums.count - 2], sortedNums[sortedNums.count - 3]])
                continue
            }
            
            for j in i + 1 ... sortedNums.count - 3 {
                if j > i + 1 && sortedNums[j] == sortedNums[j - 1] {
                    continue
                }
                
                let partialTarget = target - sortedNums[i] - sortedNums[j]
                
                // pruning
                let minSum = sortedNums[j + 1] + sortedNums[j + 2]
                if minSum > partialTarget {
                    break
                } else if minSum == partialTarget {
                    result.append([sortedNums[i], sortedNums[j], sortedNums[j + 1], sortedNums[j + 2]])
                    break
                }
                
                let maxSum = sortedNums[sortedNums.count - 1] + sortedNums[sortedNums.count - 2]
                if maxSum < partialTarget {
                    continue
                } else if maxSum == partialTarget {
                    result.append([sortedNums[i], sortedNums[j], sortedNums[sortedNums.count - 1], sortedNums[sortedNums.count - 2]])
                    continue
                }
                
                var left = j + 1
                var right = sortedNums.count - 1
                
                while left < right {
                    let sum = sortedNums[left] + sortedNums[right]
                    if sum == partialTarget {
                        result.append([sortedNums[i], sortedNums[j], sortedNums[left], sortedNums[right]])
                    }
                    
                    if sum < partialTarget {
                        repeat {
                            left += 1
                        } while (left < right && sortedNums[left] == sortedNums[left - 1])
                    } else {
                        repeat {
                            right -= 1
                        } while (left < right && sortedNums[right] == sortedNums[right + 1])
                    }
                }
            }
        }
        
        return result
    }
}
