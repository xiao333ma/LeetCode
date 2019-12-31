//
//  MaximumSubarray.swift
//  LeetCode_Swift
//
//  Created by xiaoma on 6/10/17.
//  Copyright © 2017 xiaoma. All rights reserved.
//

import Cocoa

class MaximumSubarray: NSObject {
    //DP
    func maxSubArray(_ nums: [Int]) -> Int {
    
        var max_current = nums[0]
        var max_global = nums[0]
        
        for i in 1 ..< nums.count {
            max_current = max(max_current + nums[i], nums[i])
            max_global = max(max_global, max_current)
        }
        return max_global
    }
}
