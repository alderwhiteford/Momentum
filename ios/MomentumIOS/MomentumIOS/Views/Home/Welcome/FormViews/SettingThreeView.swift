//
//  SettingTwoView.swift
//  MomentumIOS
//
//  Created by Alder Whiteford on 8/23/24.
//

import Foundation
import SwiftUI

struct SettingThreeView: View {
    @Binding var state: Int
    @Binding var date: Date
    
    @EnvironmentObject var authManager: AuthManager
    
    var body: some View {
        GeometryReader { geometry in
            VStack {
                VStack(alignment: .leading) {
                    ProgressBar(
                        currentStep: 1,
                        numberOfSteps: 4,
                        title: "Goal Setting"
                    )
                    
                    Spacer().frame(maxHeight: 16)
                    
                    MomentumButton(props: MomentumButtonProps(
                        id: "setting-three-back",
                        displayName: "",
                        onClick: {
                            state = state - 1
                        },
                        disabled: false,
                        icon: Image("back"),
                        variant: MomentumButtonVariant.none))
                    
                    Text("When would you like to achieve this?")
                        .font(.lightAppSubheader)
                    
                    Spacer().frame(maxHeight: 24)
                    
                    Text("Setting deadlines makes it far easier to track progress and stick to a plan.")
                        .font(.appCaption)
                    
                    Spacer().frame(maxHeight: 24)
                    
                    Calendar(date: $date)
                    
                    Spacer().frame(maxHeight: 16)
                    
                    Text("Try to be as specific as possible, but you can always change this later if needed.")
                        .font(.lightAppCaptionThree)
                    
                    Spacer()
                    
                    InfoDrawer(
                        props: InfoDrawerProps(
                            id: "setting-three-info",
                            title: "Deadline Tips",
                            header: "Setting a deadline",
                            subheader: "Understanding your \"why\" or the underlying reason for your goal is essential for maintaining motivation and ensuring your efforts align with what truly matters to you. Here’s how to clarify your purpose:",
                            sections: [
                                InfoDrawerSection(
                                    title: "Connect with Your Values",
                                    description: "Ensure your goal reflects your core personal values. This alignment imbues your journey with deeper meaning and satisfaction."
                                ),
                                InfoDrawerSection(
                                    title: "Identify Personal Benefits",
                                    description: "Consider how achieving this goal will benefit you personally. Whether it’s growth, happiness, or fulfillment, recognizing these benefits can fuel your determination."
                                ),
                                InfoDrawerSection(
                                    title: "Visualize Success",
                                    description: "Imagine achieving your goal. How does it feel? Visualization can strengthen your commitment and clarify your why."
                                ),
                                InfoDrawerSection(
                                    title: "Consider the Broader Impact",
                                    description: "Think about how your goal affects others. Goals that contribute to a larger good often carry a powerful motivational force."
                                ),
                                InfoDrawerSection(
                                    title: "Revisit and Reflect",
                                    description: "Your why might evolve. Regular reflection ensures your goal continues to resonate with your aspirations and values."
                                )
                            ],
                            prompts: InfoDrawerPrompts(
                                bulleted: true,
                                questions: [
                                    "How does this goal align with your core values?",
                                    "What personal benefits do you anticipate from achieving this goal?",
                                    "Imagine achieving your goal. What does success look like and feel like to you?",
                                    "How does your goal contribute to the well-being of others or the larger community?",
                                    "How might your why evolve over time, and how will you stay connected to it?"
                                ]
                            )
                        )
                    )
                    
                    Spacer().frame(maxHeight: 16)
                                        
                    MomentumButton(props: MomentumButtonProps(
                        id: "setting-two-button",
                        displayName: "Continue",
                        onClick: {
                            state = state + 1
                        },
                        disabled: false,
                        icon: nil,
                        variant: MomentumButtonVariant.filled
                    ))
                    
                    // TODO: WILL HAVE TO REMOVE
                    Button(action: {
                        Task {
                            do {
                                try await authManager.signOut()
                            }
                            catch {
                                print("Failed to signout!")
                            }
                        }
                    }) {
                        HStack {
                            Text("Sign out")
                                .foregroundColor(.white)
                                .font(.appCaption)
                        }
                        .frame(maxWidth: .infinity)
                    }
                    .frame(alignment: .bottom)
                }
                .frame(
                    maxWidth: .infinity,
                    maxHeight: .infinity,
                    alignment: .topLeading
                )
            }
            .padding(geometry.size.width * 0.05)
            .padding(.bottom, geometry.size.width * 0.075)
            .padding(.top, geometry.size.height * 0.05)
            .frame(
                maxWidth: .infinity,
                maxHeight: .infinity
            )
            .background(Color("Background"))
            .foregroundColor(Color("Background-On"))
            .ignoresSafeArea()
        }
    }
}
