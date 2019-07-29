//
//  SearchInRotatedSortedArray.swift
//  LeetCode_Swift
//
//  Created by xiaoma on 6/2/17.
//  Copyright © 2017 xiaoma. All rights reserved.
//

import Cocoa

class SearchInRotatedSortedArray: NSObject {

    func search(_ nums: [Int], _ target: Int) -> Int {
        var min = 0
        var mid = 0
        var max = nums.count - 1
        while min <= max {
            mid = (max + min) / 2

            if nums[mid] == target {
                return mid
            }
            
            if nums[min] <= nums[mid] {
                if nums[min] <= target && nums[mid] > target {
                    max = mid - 1
                } else {
                    min = mid + 1
                }
                
            } else {
                if nums[mid] < target && nums[max] >= target {
                    min = mid + 1
                } else {
                    max = mid - 1
                }
                
            }
        }
        return -1
    }
}
