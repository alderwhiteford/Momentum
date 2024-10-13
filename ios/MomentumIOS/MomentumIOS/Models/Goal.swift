//
//  User.swift
//  MomentumIOS
//
//  Created by Alder Whiteford on 10/12/24.
//

import Foundation

enum SuccessIndicators: String {
    case target
    case satisfied
    case completion
}

struct Goal {
    let description: String;
    let estimatedCompletionAt: Date;
    let theWhy: String;
    let whenSuccess: [SuccessIndicators]
}
