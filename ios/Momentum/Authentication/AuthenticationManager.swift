//
//  AuthenticationManager.swift
//  Momentum
//
//  Created by Alder Whiteford on 6/25/24.
//

import Foundation
import FirebaseAuth

// Storing the result of firebase authentication:
struct AuthDataModelResult {
    let uid: String
    let email: String?
    let photoUrl: URL?
    
    init(user: User) {
        self.uid = user.uid
        self.email = user.email
        self.photoUrl = user.photoURL
    }
}

// Managing all authentication related tasks:
final class AuthenticationManager {
    // Creates a singleton of the auth manager:
    static let shared = AuthenticationManager()
    private init() { }
    
    // Retrieve the current authenticated user:
    func getAuthenticatedUser() throws -> AuthDataModelResult {
        guard let user = Auth.auth().currentUser else {
            throw URLError(.badServerResponse)
        }
        
        return AuthDataModelResult(user: user)
    }
}
