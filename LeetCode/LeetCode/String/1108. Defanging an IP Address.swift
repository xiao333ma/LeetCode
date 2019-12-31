//
//  1108. Defanging an IP Address.swift
//  LeetCode
//
//  Created by xiaoma on 2019/12/30.
//  Copyright © 2019 xiaoma. All rights reserved.
//

import Foundation


class Solution {
    func defangIPaddr(_ address: String) -> String {
        return address.replacingOccurrences(of: ".", with: "[.]")
    }
}
