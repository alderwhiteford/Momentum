//
//  ContentView.swift
//  MomentumIOS
//
//  Created by Alder Whiteford on 8/13/24.
//

import SwiftUI
import GoogleSignIn

enum ViewType {
    case home
    case auth
}

struct MomentumAppView: View {  
    @EnvironmentObject var authManager: AuthManager
    @EnvironmentObject var errorManager: ErrorManager
        
    var body: some View {
        ZStack {
            Color("Background")
                .ignoresSafeArea()
            
            VStack {
                if authManager.session != nil {
                    HomeView()
                        .transition(.slide)
                } else {
                    AuthView()
                        .transition(.slide)
                }
            }
            .ignoresSafeArea()
            
            VStack {
                if errorManager.errorMessage != "" {
                    VStack {
                        Spacer().frame(height: 5)
                        ErrorBanner(props: ErrorBannerProps(message: errorManager.errorMessage))
                        Spacer()
                    }
                    .transition(.move(edge: .top).combined(with: .opacity))
                }
            }
            .animation(.easeInOut(duration: 1), value: errorManager.errorMessage)
            .zIndex(1)
        }
    }
}

#Preview {
    MomentumAppView()
}
