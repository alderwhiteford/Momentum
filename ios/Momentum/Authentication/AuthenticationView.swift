//
//  AuthenticationView.swift
//  Momentum
//
//  Created by Alder Whiteford on 6/25/24.
//

import SwiftUI
import GoogleSignIn
import GoogleSignInSwift
import FirebaseAuth

@MainActor
final class AuthenticationViewModel: ObservableObject {
    func googleSignIn async throws {
        // Retrieve the top view controller
        guard let topVC = Utilities.shared.topViewController() else {
            throw(URLError(.cannotFindHost))
        }
        
        // Initialize sign-in with google
        let googleSignInResult = try await GIDSignIn.sharedInstance.signIn(withPresenting: topVC)
        
        // Parse the idToken
        guard let idToken = googleSignInResult.user.idToken?.tokenString else {
            throw(URLError(.badServerResponse))
        }
        
        // Parse the accessToken
        let accessToken = googleSignInResult.user.accessToken.tokenString
        
        // Retrieve the credentials from google
        let credential = GoogleAuthProvider.credential(withIDToken: idToken, accessToken: accessToken)
    }
}

struct AuthenticationView: View {
    @StateObject private var viewModel = AuthenticationViewModel()
    
    var body: some View {
        VStack {
            GoogleSignInButton(viewModel: GoogleSignInButtonViewModel(scheme: .dark, style: .wide, state: .normal)) {}
        }
        .padding()
        .navigationTitle("Sign-In")
    }
}

struct AuthenticationView_Previews: PreviewProvider {
    static var previews: some View {
        NavigationStack {
            AuthenticationView()
        }
    }
}
