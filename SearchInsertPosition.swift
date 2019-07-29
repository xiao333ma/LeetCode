//
//  SearchInsertPosition.swift
//  LeetCode_Swift
//
//  Created by xiaoma on 6/5/17.
//  Copyright © 2017 xiaoma. All rights reserved.
//

import Cocoa

class SearchInsertPosition: NSObject {

    func searchInsert(_ nums: [Int], _ target: Int) -> Int {
        
        
        var left = 0
        var middle = 0
        var right = nums.count - 1
        
        while left <= right {
            middle = (left + right) / 2
            if nums[middle] == target {
                return middle
            } else if (nums[middle] < target) {
                left = middle + 1
            } else {
                right = middle - 1
            }
        }
        return left
    }
}
