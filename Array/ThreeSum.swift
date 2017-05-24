//
//  3Sum.swift
//  LeetCode_Swift
//
//  Created by xiaoma on 2017/5/20.
//  Copyright © 2017年 xiaoma. All rights reserved.
//

import Cocoa

class ThreeSum{
    func threeSum(_ nums: [Int]) -> [[Int]] {
        guard nums.count >= 3 else {
            return Array()
        }
        
        let sortedNums = nums.sorted(by: >);
        var result = [[Int]]();
        
        for i in 0 ... sortedNums.count - 3 {
            if i == 0 || sortedNums[i] != sortedNums[i - 1] {
                let numI = sortedNums[i];
                var left = i + 1
                var right = sortedNums.count - 1
                while left < right {
                    let leftNum = sortedNums[left]
                    let rightNum = sortedNums[right]
                    if leftNum + rightNum + numI == 0 {
                        let arr = [numI, leftNum, rightNum]
                        result.append(arr)
                        repeat {
                            left += 1
                        } while (left < right && sortedNums[left] == sortedNums[left - 1])
                        repeat {
                            right -= 1
                        } while (left < right && sortedNums[right] == sortedNums[right + 1])
                        
                        
                    } else if (leftNum + rightNum + numI < 0) {
                        repeat {
                            right -= 1
                        } while (left < right && sortedNums[right] == sortedNums[right + 1])
                    } else {
                        repeat {
                            left += 1
                        } while (left < right && sortedNums[left] == sortedNums[left - 1])
                    }
                }

            }
            
        }
        return result
    }
}
