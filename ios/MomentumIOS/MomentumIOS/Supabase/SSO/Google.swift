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
        
    // Handle Google Sign In:
    func handleGoogleSignIn() async throws -> Void {
        let topVC = await MainActor.run {
            return retrieveRootViewController()
        }
        
        // Handle the OAuth flow:
        let googleFlowResponse = try await initializeSignInFlow(topVC: topVC);
        
        // Close out the flow if the user cancelled the request:
        if googleFlowResponse.idToken == "" {
            return
        }
        
        // Initiate supabase auth:
        try await supabase.auth.signInWithIdToken(credentials: .init(provider: .google, idToken: googleFlowResponse.idToken))
    }
    
    // Initialize the OAuth User Flow:
    private func initializeSignInFlow(topVC: UIViewController) async throws -> SignInWithGoogleResponse {
        try await withCheckedThrowingContinuation { continuation in
            GIDSignIn.sharedInstance.signIn(withPresenting: topVC) { signInResult, error in
                if error != nil {
                    continuation.resume(returning: SignInWithGoogleResponse(idToken: ""))
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
    
    // Retrieve the root view of the controller:
    func retrieveRootViewController() -> UIViewController {
        if let rootViewController = UIApplication.shared.rootViewController {
            return rootViewController
        } else {
            fatalError("Failed to find the root view controller.")
        }
    }
}
