//
//  AuthViewController.swift
//  MomentumIOS
//
//  Created by Alder Whiteford on 8/13/24.
//

import Foundation
import GoogleSignIn
import UIKit

// EXTENSION OF UI APPLICATION TO RETRIEVE ROOT VIEW CONTROLLER
extension UIApplication {
    var rootViewController: UIViewController? {
        // Get the connected scenes
        return self.connectedScenes
            // Keep only active scenes, onscreen and visible to the user
            .filter { $0.activationState == .foregroundActive }
            // Keep only the first UIWindowScene
            .compactMap { $0 as? UIWindowScene }
            // Get its associated windows
            .first?.windows
            // Finally, get the rootViewController of the first window
            .first(where: { $0.isKeyWindow })?.rootViewController
    }
}

class AuthViewModel: ObservableObject {
    @Published var googleClient: GoogleController
    // ADD NEW CLIENTS HERE
    
    init() {
        self.googleClient = GoogleController();
    }
    
    // Handle Google Sign In:
    func handleGoogleSignIn() async throws -> Void {
        let topVC = await MainActor.run {
            return retrieveRootViewController()
        }
        
        // Handle the OAuth flow:
        let googleFlowResponse = try await self.googleClient.initializeSignInFlow(topVC: topVC);
        
        // Initiate supabase auth:
        try await supabase.auth.signInWithIdToken(credentials: .init(provider: .google, idToken: googleFlowResponse.idToken))
    }
    
    func generateNonce() -> String {
        let nonceLength = 32
        let characters = "0123456789ABCDEFGHIJKLMNOPQRSTUVXYZabcdefghijklmnopqrstuvwxyz-._"
        var nonce = ""
        
        for _ in 0..<nonceLength {
            let randomIndex = Int.random(in: 0..<characters.count)
            nonce.append(characters[characters.index(characters.startIndex, offsetBy: randomIndex)])
        }
        
        return nonce
    }
    
    // Retrieve the root view of the controller:
    func retrieveRootViewController() -> UIViewController {
        if let rootViewController = UIApplication.shared.rootViewController {
            return rootViewController
        } else {
            fatalError("Failed to find the root view controller.")
        }
    }
}
