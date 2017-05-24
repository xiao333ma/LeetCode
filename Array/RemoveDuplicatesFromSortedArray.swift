//
//  RemoveDuplicatesFromSortedArray.swift
//  LeetCode_Swift
//
//  Created by xiaoma on 5/24/17.
//  Copyright © 2017 xiaoma. All rights reserved.
//

import Cocoa

class RemoveDuplicatesFromSortedArray {
    func removeDuplicates(_ nums: inout [Int]) -> Int {
        if nums.count < 2 {
            return nums.count
        }
        var count = 1
        
        for i in 1 ..< nums.count {
            if nums[i] == nums[i - 1] {
                continue
            } else {
                nums[count] = nums[i]
                count += 1
            }
        }
        
        return count
    }
}
