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
    
    static let appTitle = appFont(size: 45).weight(.heavy)
    static let lightAppTitle = appFont(size: 45).weight(.regular)
    
    static let lightAppSubheader = appFont(size: 32).weight(.medium)
    static let appSubheader = appFont(size: 32).weight(.regular)
    static let boldAppSubheader = appFont(size: 32).weight(.heavy)
    
    static let appBody = appFont(size: 24).weight(.medium)
    
    static let boldAppCaption = appFont(size: 16).weight(.heavy)
    static let appCaption = appFont(size: 16).weight(.regular)
    
    static let lightAppCaptionTwo = appFont(size: 14).weight(.regular)
    static let appCaptionTwo = appFont(size: 14).weight(.medium)
    static let boldAppCaptionTwo = appFont(size: 14).weight(.heavy)
    
    static let lightAppCaptionThree = appFont(size: 11).weight(.regular)
    static let appCaptionThree = appFont(size: 11).weight(.medium)
    static let boldAppCaptionThree = appFont(size: 11).weight(.heavy)
}
