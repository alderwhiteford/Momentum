//
//  Button.swift
//  MomentumIOS
//
//  Created by Alder Whiteford on 8/21/24.
//

import Foundation
import SwiftUI

enum MomentumButtonVariant {
    case filled
    case outline
    case none
}

struct MomentumButtonProps {
    let id: String
    let displayName: String
    let onClick: () -> Void
    let disabled: Bool
    let icon: Image?
    let variant: MomentumButtonVariant
}

struct MomentumButtonStyling: ButtonStyle {
    var disabled: Bool
    var variant: MomentumButtonVariant = MomentumButtonVariant.filled
    
    func makeBody(configuration: Configuration) -> some View {
        configuration.label
            .padding(variant == MomentumButtonVariant.none ? 0 : 16)
            .background(Color(
                variant == MomentumButtonVariant.filled ?
                disabled ? "Disabled-Container" : "MomentumPrimary"
                : ""
            ))
            .border(Color(
                variant == MomentumButtonVariant.outline ?
                disabled ? "Disabled-Container" : "MomentumPrimary"
                : ""
            ), width: 1)
            .clipShape(RoundedRectangle(cornerRadius: 4))
            .scaleEffect(configuration.isPressed ? 0.95 : 1.0)
            .shadow(radius: 5)
    }
}

struct MomentumButton: View {
    var props: MomentumButtonProps
        
    var body: some View {
        Button(action: {
            Task {
                do {
                    props.onClick()
                }
            }
        }) {
            HStack {
                if let icon = props.icon {
                    icon
                        .resizable()
                        .scaledToFit()
                        .frame(width: 20, height: 20)
                }
                Text(props.displayName)
                    .font(.appCaptionTwo)
                    .foregroundColor(
                        props.disabled ? Color.disabledContainerOn :
                            props.variant == MomentumButtonVariant.none ? Color.momentumPrimary : Color.primaryOn
                        
                    )
            }
            .frame(maxWidth: props.variant == MomentumButtonVariant.none ? nil : .infinity)
        }
        .buttonStyle(MomentumButtonStyling(disabled: props.disabled, variant: props.variant))
        .frame(maxWidth: props.variant == MomentumButtonVariant.none ? nil : .infinity, alignment: .bottom)
        .disabled(props.disabled)
        .id(props.id)
    }
}
