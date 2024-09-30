//
//  Error.swift
//  MomentumIOS
//
//  Created by Alder Whiteford on 9/28/24.
//

import Foundation
import SwiftUI

struct ErrorBannerProps {
    var message: String
}

struct ErrorBanner: View {
    var props: ErrorBannerProps
    
    init(props: ErrorBannerProps) {
        self.props = props
        
        if props.message.count > 30 {
            var startIndex = self.props.message.startIndex
            var endIndex = self.props.message.index(startIndex, offsetBy: 30)
            self.props.message = String(self.props.message[startIndex..<endIndex]) + "..."
        }
    }
    
    
    var body: some View {
        VStack {
            Text(self.props.message)
                .foregroundColor(.white)
                .font(.appCaption)
                .padding(10)
                .frame(maxWidth: .infinity)
                .background(Color.red)
                .shadow(radius: 5)
                .clipShape(RoundedRectangle(cornerRadius: 4))
        }
        .padding(.horizontal, 16)
        .frame(maxWidth: .infinity, alignment: .top)
    }
}
