//
//  GoogleController.swift
//  MomentumIOS
//
//  Created by Alder Whiteford on 8/13/24.
//

import Foundation
import GoogleSignIn
import Supabase

struct SignInWithGoogleResponse {
    let idToken: String
}

enum GoogleAuthError: Error {
    case signInFailed
    case noIdToken
}

class GoogleController: ObservableObject {
    
    // Initialize the OAuth User Flow:
    func initializeSignInFlow(topVC: UIViewController) async throws -> SignInWithGoogleResponse {
        try await withCheckedThrowingContinuation { continuation in
            GIDSignIn.sharedInstance.signIn(withPresenting: topVC) { signInResult, error in
                if error != nil {
                    continuation.resume(throwing: GoogleAuthError.signInFailed)
                    return
                }
                
                guard let idToken = signInResult?.user.idToken else {
                    continuation.resume(throwing: GoogleAuthError.noIdToken)
                    return
                }
                
                let response = SignInWithGoogleResponse(idToken: idToken.tokenString)
                continuation.resume(returning: response)
            }
        }
    }
}
