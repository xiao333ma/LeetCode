//
//  SearchForARange.swift
//  LeetCode_Swift
//
//  Created by xiaoma on 6/2/17.
//  Copyright © 2017 xiaoma. All rights reserved.
//

import Cocoa

class SearchForARange: NSObject {
    func searchRange(_ nums: [Int], _ target: Int) -> [Int] {
        var index = -1
        
        var left = 0
        var right = nums.count - 1
        var middle = 0;
        while left <= right {
            
            
            if nums[left] == target {
                index = left
                break
            }
            
            if nums[right] == target {
                index = right
                break
            }
            
            middle = right - (right-left)/2

            if nums[middle] == target {
                index = middle
                break
            }
            
            if (nums[middle] < target && target < nums[right]) {
                left = middle + 1
                right -= 1
            } else {
                right = middle - 1
                left += 1
            }
        }
        
        if index == -1 {
            return [-1, -1]
        } else {
            var j = index
            while j > 0 && nums[j - 1] == target {
                j -= 1
            }
            
            var i = index
            while  i < nums.count - 1 && nums[i + 1] == target{
                i += 1
            }
            return [j, i]
        }
    }
}
