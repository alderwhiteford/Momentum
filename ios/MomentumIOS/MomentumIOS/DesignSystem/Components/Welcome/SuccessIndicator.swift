//
//  SuccessIndicator.swift
//  MomentumIOS
//
//  Created by Alder Whiteford on 10/12/24.
//

import Foundation
import SwiftUI

let successIndicatorVariants: [SuccessIndicators: (image: Image, description: String)] = [
    .completion: (image: Image(.check), description: "When I complete or achieve something"),
    .target: (image: Image(.satisfy), description: "When I reach a target metric"),
    .satisfied: (image: Image(.target), description: "When I feel satisfied"),
]

struct SuccessIndicator: View {
    var successIndicator: SuccessIndicators
    var onClick: () -> Void
    
    init(successIndicator: SuccessIndicators, onClick: @escaping () -> Void) {
        self.successIndicator = successIndicator
        self.onClick = onClick
    }
        
    var body: some View {
        HStack(spacing: 16) {
            successIndicatorVariants[self.successIndicator]?.image
                .resizable()
                .scaledToFit()
                .aspectRatio(contentMode: .fit)
                .frame(width: 16)
            
            Text(successIndicatorVariants[self.successIndicator]?.description ?? "")
                .font(.appCaptionTwo)
                .foregroundColor(Color(.primaryContainerOn))
        }
        .frame(
            maxWidth: .infinity,
            maxHeight: 45,
            alignment: .leading
        )
        .padding(.horizontal)
        .background(Color(.primaryContainer))
        .cornerRadius(4.0)
        .onTapGesture {
            self.onClick()
        }
    }
}
