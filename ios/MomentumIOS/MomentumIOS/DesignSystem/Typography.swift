//
//  Typography.swift
//  MomentumIOS
//
//  Created by Alder Whiteford on 8/13/24.
//

import Foundation
import SwiftUI

extension Font {
    static func appFont(size: CGFloat) -> Font {
            return Font.custom("Work Sans", size: size)
        }
        
    static let appTitle = appFont(size: 32).weight(.bold)
    static let appBody = appFont(size: 24).weight(.medium)
    static let appCaption = appFont(size: 16).weight(.regular)
}
