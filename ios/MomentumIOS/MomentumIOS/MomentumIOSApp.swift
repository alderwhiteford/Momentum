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
    @StateObject private var authManager = AuthManager()
    @StateObject private var errorManager = ErrorManager()

    var body: some Scene {
        WindowGroup {
            MomentumAppView()
            .onOpenURL(perform: { url in
                GIDSignIn.sharedInstance.handle(url)
            })
            .environmentObject(authManager)
            .environmentObject(errorManager)
        }
    }
}
