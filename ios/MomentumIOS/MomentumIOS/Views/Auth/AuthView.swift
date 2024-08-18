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
    let authController = AuthViewModel()
        
    var body: some View {
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
                
            GoogleButton(handleSignIn: authController.handleGoogleSignIn)
                .frame(alignment: .bottom)
        }
        .frame(maxHeight: .infinity)
      }
}

#Preview {
    AuthView()
}
