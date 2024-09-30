//
//  ErrorManager.swift
//  MomentumIOS
//
//  Created by Alder Whiteford on 9/28/24.
//

import Foundation

class ErrorManager: ObservableObject {
    @Published var errorMessage: String
    
    init() {
        self.errorMessage = ""
    }
    
    func setError(errorMessage: String) -> Void {
        // Update the error message:
        self.errorMessage = errorMessage
        
        // Clear the error after 4 seconds:
        DispatchQueue.main.asyncAfter(deadline: .now() + 4) {
            self.errorMessage = ""
        }
    }
}
