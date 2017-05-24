//
//  RemoveElement.swift
//  LeetCode_Swift
//
//  Created by xiaoma on 5/24/17.
//  Copyright © 2017 xiaoma. All rights reserved.
//

import Cocoa

class RemoveElement {

    func removeElement(_ nums: inout [Int], _ val: Int) -> Int {
        
        
        let length = nums.count
        var count = 0
        for i in 0 ..< length  {
            let numI = nums[i]
            if numI != val  {
                nums[count] = numI
                count += 1
            }
        }
        return count
    }
}
