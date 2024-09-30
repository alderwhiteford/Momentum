//
//  AuthView.swift
//  MomentumIOS
//
//  Created by Alder Whiteford on 8/13/24.
//

import Foundation
import SwiftUI
import Supabase

struct AuthView: View {
    @EnvironmentObject var authManager: AuthManager
    @EnvironmentObject var errorManager: ErrorManager
    
    var body: some View {
        GeometryReader { geometry in
            VStack (spacing: 50) {
                VStack(spacing: 30) {
                    Image("Momentum")
                        .resizable()
                        .scaledToFit()
                        .aspectRatio(contentMode: .fit)
                        .frame(width: UIScreen.main.bounds.width - 225)
                    
                    VStack(spacing: 10) {
                        Text("Welcome to Momentum")
                            .font(.appBody)
                            .foregroundColor(Color("MomentumPrimary"))
                            .multilineTextAlignment(.center)
                        
                        Text("Please select a provider below to continue")
                            .font(.appCaption)
                            .foregroundColor(.white)
                            .multilineTextAlignment(.center)
                    }
                }
                .frame(maxHeight: .infinity, alignment: .center)
                
                SSOButton(client: SSOClientContent(
                    id: "google",
                    displayName: "Google",
                    image: Image("google"),
                    handleSignIn: authManager.googleClient.handleGoogleSignIn
                ))
            }
            .padding(geometry.size.width * 0.05)
            .padding(.bottom, geometry.size.width * 0.075)
            .padding(.top, geometry.size.height * 0.05)
            .frame(maxHeight: .infinity)
        }
      }
}

#Preview {
    AuthView()
}
