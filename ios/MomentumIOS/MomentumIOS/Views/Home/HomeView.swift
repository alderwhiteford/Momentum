//
//  Home.swift
//  MomentumIOS
//
//  Created by Alder Whiteford on 8/19/24.
//

import Foundation
import SwiftUI

struct HomeView: View {
    @EnvironmentObject var authManager: AuthManager
    
    var body: some View {
//        VStack {
//            Text("Welcome to the home view")
//                .font(.appCaption)
//                .foregroundColor(.white)
//                .multilineTextAlignment(.center)
//            
//            Button(action: {
//                Task {
//                    do {
//                        try await authManager.signOut()
//                    }
//                    catch {
//                        print("Failed to signout!")
//                    }
//                }
//            }) {
//                HStack {
//                    Text("Sign out")
//                        .foregroundColor(.white)
//                        .font(.appCaption)
//                }
//                .frame(maxWidth: .infinity)
//            }
//            .buttonStyle(SSOButtonStyling())
//            .frame(alignment: .bottom)
//        }.frame(maxHeight: .infinity)
        WelcomeView()
    }
}

#Preview {
    HomeView()
}
