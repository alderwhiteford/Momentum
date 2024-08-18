//
//  ContentView.swift
//  MomentumIOS
//
//  Created by Alder Whiteford on 8/13/24.
//

import SwiftUI
import GoogleSignIn

struct MomentumAppView: View {  

    var body: some View {
        ZStack {
            Color("Background")
                .ignoresSafeArea()
            
            VStack {
                AuthView()
            }
            .padding(10)
        }
    }
}

#Preview {
    MomentumAppView()
}
