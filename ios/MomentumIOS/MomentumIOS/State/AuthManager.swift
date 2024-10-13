//
//  AuthListener.swift
//  MomentumIOS
//
//  Created by Alder Whiteford on 8/19/24.
//

import Foundation
import Combine
import Supabase
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

enum SSOClient: String {
    case Google
}

class AuthManager: ObservableObject {
    @Published var session: Session?
    @Published var googleClient: GoogleController
    
    init() {
        // Initialize OAuth Clients:
        self.googleClient = GoogleController()
        
        // Set initial session:
        self.session = supabase.auth.currentSession
        
        // Refresh session on start-up and listen for auth changes:
        if self.session != nil {
            Task {
                do {
                    try await supabase.auth.refreshSession()
                } catch {
                    print("failed to refresh session")
                }
            }
        }
        
        Task {
            await listenForAuthChanges()
        }
    }
    
    func signOut() async throws {
        try await supabase.auth.signOut();
    }
    
    private func listenForAuthChanges() async {
        for await (_, session) in supabase.auth.authStateChanges {
            let _ = print("Session updated: \(session)")
            DispatchQueue.main.async {
                self.session = session
            }
        }
    }
}
