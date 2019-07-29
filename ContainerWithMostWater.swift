
//  ContainerWithMostWater.swift
//  LeetCode_Swift
//
//  Created by xiaoma on 2017/5/19.
//  Copyright © 2017年 xiaoma. All rights reserved.
//

import Cocoa

class ContainerWithMostWater: NSObject {
    func maxArea(_ height: [Int]) -> Int {
        var sum = 0
        var left = 0
        var right = height.count - 1
        
        while left < right {
            let leftHeight = height[left]
            let rightHeight = height[right];
            sum = max((right - left) * min(leftHeight, rightHeight), sum)
            if leftHeight < rightHeight {
                left += 1
            } else {
                right -= 1
            }
        }
        return sum
        
    }
}
