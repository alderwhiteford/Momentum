//
//  GoogleView.swift
//  MomentumIOS
//
//  Created by Alder Whiteford on 8/13/24.
//

import Foundation
import SwiftUI
import OSLog

struct GoogleButtonStyle: ButtonStyle {
    func makeBody(configuration: Configuration) -> some View {
        configuration.label
            .padding()
            .background(Color("Surface-Container-3"))
            .clipShape(RoundedRectangle(cornerRadius: 10))
            .foregroundColor(.black)
            .scaleEffect(configuration.isPressed ? 0.95 : 1.0)
            .shadow(radius: 5)
    }
}

struct GoogleButton: View {
    @State private var googleError: String = ""
    var handleSignIn: () async throws -> Void
    
    var body: some View {
        VStack {
            Text(googleError)
                .foregroundColor(.red)
            Button(action: {
                Task {
                    do {
                        try await handleSignIn()
                    } catch {
                        googleError = error.localizedDescription
                    }
                }
            }) {
                HStack {
                    Image("google")
                        .resizable()
                        .scaledToFit()
                        .frame(width: 20, height: 20)
                    
                    Text("Continue with Google")
                        .foregroundColor(.white)
                        .font(.appCaption)
                }
                .frame(maxWidth: .infinity)
            }
            .buttonStyle(GoogleButtonStyle())
            .frame(alignment: .bottom)
        }
    }
}
