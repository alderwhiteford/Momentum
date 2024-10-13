//
//  SSOButton.swift
//  MomentumIOS
//
//  Created by Alder Whiteford on 8/19/24.
//

import Foundation
import SwiftUI
import OSLog

struct SSOClientContent {
    let id: String
    let displayName: String
    let image: Image
    let handleSignIn: () async throws -> Void
}

struct SSOButtonStyling: ButtonStyle {
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

struct SSOButton: View {
    @EnvironmentObject var errorManager: ErrorManager
    var client: SSOClientContent
    @State var isLoading: Bool = false
    
    init(client: SSOClientContent) {
        self.client = client
    }
    
    var body: some View {
        VStack {
            Button(action: {
                Task {
                    do {
                        isLoading = true
                        try await client.handleSignIn()
                        isLoading = false
                    }
                    catch {
                        print(error.localizedDescription)
                        errorManager.setError(errorMessage: "Something went wrong...")
                        isLoading = false
                    }
                }
            }) {
                HStack {
                    if isLoading {
                        GifImage("loading")
                            .frame(width: 20, height: 20)
                    } else {
                        client.image
                            .resizable()
                            .scaledToFit()
                            .frame(width: 20, height: 20)
                        
                        Text("Continue with \(client.displayName)")
                            .foregroundColor(.white)
                            .font(.appCaption)
                    }
                }
                .frame(maxWidth: .infinity)
            }
            .buttonStyle(SSOButtonStyling())
            .frame(alignment: .bottom)
        }.id(client.id)
    }
}
