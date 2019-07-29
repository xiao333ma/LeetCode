//
//  ThreeSumClosest.swift
//  LeetCode_Swift
//
//  Created by xiaoma on 5/22/17.
//  Copyright © 2017 xiaoma. All rights reserved.
//

import Foundation

class ThreeSumClosest {
    
    func threeSumClosest(_ nums: [Int], _ target: Int) -> Int {
        
        guard nums.count >= 3 else {
            return 0
        }
        
        var sortedNums = nums.sorted(by: >);
        var sum = nums[0] + nums[1] + nums[2]
        
        for i in 0 ..< sortedNums.count - 2 {
            
            var left = i + 1
            var right = sortedNums.count - 1
            let iNum = sortedNums[i]
            while left < right {
                
                let currSum = sortedNums[left] + sortedNums[right] + iNum
                
                if abs(target - currSum) < abs(target - sum) {
                    sum = currSum
                    if (currSum == target){
                        return sum;
                    }
                    
                }
                
                if (currSum > target){
                    left += 1
                } else {
                    right -= 1
                }
            }
            
        }
        return sum
    }
}
