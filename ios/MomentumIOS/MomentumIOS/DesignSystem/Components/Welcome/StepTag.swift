//
//  StepTag.swift
//  MomentumIOS
//
//  Created by Alder Whiteford on 8/19/24.
//

import Foundation
import SwiftUI

struct StepTag: View {
    private var image: Image
    private var title: String
    private var description: String
    
    init(image: Image, title: String, description: String) {
        self.image = image
        self.title = title
        self.description = description
    }
    
    var body: some View {
        HStack (spacing: 16) {
            ZStack {
                VStack(){}
                    .frame(width: 64, height: 64, alignment: .center)
                self.image
                    .resizable()
                    .scaledToFit()
                    .aspectRatio(contentMode: .fit)
                    .frame(width: 25)
            }
            .background(Color("Primary-Container"))
            .clipShape(RoundedRectangle(cornerRadius: 4))
            .frame(alignment: .leading)
                                    
            VStack(alignment: .leading) {
                Text(self.title)
                    .textCase(.uppercase)
                    .font(.boldAppCaption)
                    .foregroundColor(Color("Background-On"))
                Text(self.description)
                    .font(.lightAppCaptionTwo)
                    .foregroundColor(Color("Background-On-Dark"))
            }
            .frame(maxWidth: .infinity, alignment: .leading)
        }
        .frame(
            maxWidth: .infinity,
            maxHeight: 65
        )
    }
}
