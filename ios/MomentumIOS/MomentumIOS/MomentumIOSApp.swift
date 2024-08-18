//
//  MomentumIOSApp.swift
//  MomentumIOS
//
//  Created by Alder Whiteford on 8/13/24.
//

import SwiftUI
import GoogleSignIn

@main
struct MomentumIOSApp: App {
    var body: some Scene {
        WindowGroup {
            MomentumAppView()
            .onOpenURL(perform: { url in
                GIDSignIn.sharedInstance.handle(url)
            })
        }
    }
}
