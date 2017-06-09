//
//  CombinationSumIII.swift
//  LeetCode_Swift
//
//  Created by xiaoma on 6/9/17.
//  Copyright © 2017 xiaoma. All rights reserved.
//

import Cocoa

class CombinationSumIII: NSObject {

    func combinationSum3(_ k: Int, _ n: Int) -> [[Int]] {
        
        let arr = Array((1 ... 9))
        var res = [[Int]]()
        var path = [Int]()
        dfs(&res, &path, arr, n, 0, k)
        return res
    }
    
    private func dfs(_ res: inout [[Int]], _ path: inout [Int], _ arr: [Int], _ target: Int, _ index: Int, _ size: Int) {
        
        if target == 0 && path.count == size {
            res.append(Array(path))
            return
        }
        
        for i in index ..< arr.count {
            guard arr[i] <= target else {
                break
            }
            
            path.append(arr[i])
            dfs(&res, &path, arr, target - arr[i], i + 1, size)
            path.removeLast()
        }
        
    }
    
    
}
