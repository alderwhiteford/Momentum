//
//  WelcomeView.swift
//  MomentumIOS
//
//  Created by Alder Whiteford on 8/19/24.
//

import Foundation
import SwiftUI


struct WelcomeView: View {
    @State private var settingState: Int = 0
    
    var body: some View {
        if settingState == 0 {
            InitialWelcomeView(state: $settingState)
        } else {
            WelcomeFormView(state: $settingState)
        }
    }
}

#Preview {
    WelcomeView()
}

