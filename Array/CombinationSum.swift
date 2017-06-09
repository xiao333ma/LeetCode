//
//  CombinationSum.swift
//  LeetCode_Swift
//
//  Created by xiaoma on 6/5/17.
//  Copyright © 2017 xiaoma. All rights reserved.
//

import Cocoa

class CombinationSum: NSObject {
    
    func combinationSum(_ candidates: [Int], _ target: Int) -> [[Int]] {
        var res = [[Int]]()
        var path = [Int]()
        
        dfs(candidates.sorted(by: {$0 < $1}), &res, &path, target, 0);
        return res;
    }
    private func dfs (_ candidates: [Int], _ res: inout [[Int]], _ path: inout [Int], _ target: Int, _ index: Int ) {
        
        if target == 0 {
            res.append(Array(path));
            return
        }
        
        for i in index ..< candidates.count {
            guard candidates[i] <= target else {
                break
            }
            
            path.append(candidates[i])
            dfs(candidates, &res, &path, target - candidates[i], i)
            path.removeLast()
        }
    }
}
