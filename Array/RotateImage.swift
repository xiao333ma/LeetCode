//
//  RotateImage.swift
//  LeetCode_Swift
//
//  Created by xiaoma on 6/10/17.
//  Copyright © 2017 xiaoma. All rights reserved.
//

import Cocoa

class RotateImage: NSObject {
    func rotate(_ matrix: inout [[Int]]) {
        let n = matrix.count
        
        for layer in 0 ..< n / 2 {
            let start = layer, end = n - layer - 1
            for i in start ..< end {
                let offset = i - start
                
                (matrix[start][i], matrix[i][end], matrix[end][end - offset], matrix[end - offset][start]) =
                (matrix[end - offset][start], matrix[start][i], matrix[i][end], matrix[end][end - offset])
            }
        }
    }
}
