//
//  BestTimeToBuyAndSellStockII.swift
//  LeetCode_Swift
//
//  Created by xiaoma on 2017/5/19.
//  Copyright © 2017年 xiaoma. All rights reserved.
//

import Cocoa

class BestTimeToBuyAndSellStockII {

    func maxProfit(_ prices: [Int]) -> Int {
        var sum = 0
        guard prices.count > 1 else {
            return sum
        }
        for i in 1 ..< prices.count where prices[i] > prices[i - 1] {
            sum += prices[i] - prices[i - 1]
        }
        return sum
    }
    
}
