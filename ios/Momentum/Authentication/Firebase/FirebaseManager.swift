//
//  FirebaseManager.swift
//  Momentum
//
//  Created by Alder Whiteford on 6/25/24.
//

import Foundation
import GoogleSignIn
import GoogleSignInSwift
import FirebaseAuth

// Storing the result of firebase authentication:
struct FirebaseAuthModelResult {
    let uid: String
    let email: String?
    let photoUrl: URL?
    
    init(user: User) {
        self.uid = user.uid
        self.email = user.email
        self.photoUrl = user.photoURL
    }
}

@MainActor
final class FirebaseManager {
    static let shared = FirebaseManager()
    private init() { }
    
    // Retrieve the current authenticated user:
    func getAuthenticatedUser() throws -> AuthDataModelResult {
        guard let user = Auth.auth().currentUser else {
            throw URLError(.badServerResponse)
        }
        
        return AuthDataModelResult(user: user)
    }
    
    
}

// Google SSO:
extension FirebaseManager {
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
