//
//  SettingOneView.swift
//  MomentumIOS
//
//  Created by Alder Whiteford on 8/21/24.
//

import Foundation
import SwiftUI

struct InitialWelcomeView: View {
    @Binding var state: Int
    
    let viewContent = WelcomePageOneContent
    
    var body: some View {
        GeometryReader { geometry in
            VStack (alignment: .leading) {
                VStack (alignment: .leading) {
                    Text(viewContent.title)
                        .font(.lightAppTitle)
                    
                    VStack(spacing: 24) {
                        ForEach(viewContent.stepTags) { stepTag in
                            StepTag(
                                image: Image(stepTag.icon),
                                title: stepTag.title,
                                description: stepTag.subTitle
                            )
                        }
                    }
                }
                .frame(maxHeight: .infinity)
                
                MomentumButton(props: MomentumButtonProps(
                    id: "welcome-button-1",
                    displayName: viewContent.buttonText,
                    onClick: {
                        state = state + 1
                    },
                    disabled: false,
                    icon: nil,
                    color: nil,
                    alignment: nil,
                    variant: MomentumButtonVariant.filled
                ))
            }
            .frame(
                maxWidth: .infinity,
                maxHeight: .infinity,
                alignment: .leading
            )
            .padding(geometry.size.width * 0.05)
            .padding(.bottom, geometry.size.width * 0.075)
            .background(Color("Primary-On"))
            .foregroundColor(.white)
            .ignoresSafeArea()
        }
    }
}
