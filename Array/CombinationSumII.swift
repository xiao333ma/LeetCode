//
//  CombinationSumII.swift
//  LeetCode_Swift
//
//  Created by xiaoma on 6/9/17.
//  Copyright © 2017 xiaoma. All rights reserved.
//

import Cocoa

class CombinationSumII: NSObject {

    func combinationSum2(_ candidates: [Int], _ target: Int) -> [[Int]] {
        
        var res = [[Int]]()
        var path = [Int]()
        dfs(&res, &path, candidates.sorted(by: {$0 < $1}), target, 0)
        return res
    }
    
    private func dfs(_ res: inout [[Int]], _ path: inout [Int], _ candidates: [Int], _ target: Int, _ index: Int) {
    
        if target == 0 {
            res.append(Array(path))
            return
        }
        
        for i in index ..< candidates.count  {
            
            guard candidates[i] <= target else {
                break
            }
            
            if i > 0 && candidates[i] == candidates[i - 1] && i != index {
                continue
            }
            
            path.append(candidates[i])
            dfs(&res, &path, candidates, target - candidates[i], i+1)
            path.removeLast()
            
        }
    }
}
