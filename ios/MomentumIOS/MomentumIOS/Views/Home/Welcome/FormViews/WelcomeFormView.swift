//
//  SettingAbstractView.swift
//  MomentumIOS
//
//  Created by Alder Whiteford on 9/4/24.
//

import Foundation
import SwiftUI

enum SuccessIndicators: String {
    case target
    case satisfied
    case completion
}

struct WelcomeFormView: View {
    // Form state:
    @Binding var state: Int
    
    // Progress bar state:
    @State private var progressBarState: CGFloat = 0.0
    
    // Form variables:
    @State var goal: String = ""
    @State var date: Date = Date.now
    @State var why: String = ""
    @State var successIndicators: [SuccessIndicators] = []
    
    // The Content for each of the pages:
    let pagesContent: [WelcomeForm] = [
        WelcomePageTwoContent,
        WelcomePageThreeContent,
        WelcomePageFourContent,
        WelcomePageFiveContent,
    ]
    
    var body: some View {
        // The Content for this view
        let viewContent = pagesContent[state - 1]
        
        // The state associated with this view
        let viewState: Any = [$goal, $date, $why, $successIndicators][state - 1]
        
        GeometryReader { geometry in
            VStack {
                VStack(alignment: .leading) {
                    ProgressBar(
                        progress: $progressBarState,
                        currentStep: state - 1,
                        numberOfSteps: 4,
                        title: "Goal Setting"
                    )
                    
                    Spacer().frame(maxHeight: 16)
                    
                    if viewContent.canGoBack {
                        MomentumButton(props: MomentumButtonProps(
                            id: "welcome-back-\(state)",
                            displayName: "",
                            onClick: {
                                state = state - 1
                            },
                            disabled: false,
                            icon: Image("back"),
                            variant: MomentumButtonVariant.none))
                    }
                    
                    Text(viewContent.title)
                        .font(.lightAppSubheader)
                    
                    Spacer().frame(maxHeight: 24)
                    
                    Text(viewContent.subTitle)
                        .font(.appCaption)
                    
                    Spacer().frame(maxHeight: 24)
                    
                    if viewContent.formInput == WelcomeFormInput.text {
                        TextField(
                            viewContent.formInputText ?? "",
                            text: viewState as! Binding<String>,
                            axis: .vertical
                        )
                        .padding(12)
                        .textInputAutocapitalization(.never)
                        .disableAutocorrection(true)
                        .clipShape(RoundedRectangle(cornerRadius: 4))
                        .font(.appCaption)
                        .overlay(
                            RoundedRectangle(cornerRadius: 4)
                                .stroke(Color("Outline-Dark"), lineWidth: 1)
                        )
                    }
                    
                    if viewContent.formInput == WelcomeFormInput.date {
                        Calendar(date: viewState as! Binding<Date>)
                    }

                    if let subInput = viewContent.formInputSubText {
                        Spacer().frame(maxHeight: 8)
                        
                        Text(subInput)
                            .font(.lightAppCaptionThree)
                    }
                    
                    Spacer()
                    
                    InfoDrawer(
                        props: viewContent.infoDrawer
                    )
                    
                    Spacer().frame(maxHeight: 16)
                    
                    if let buttonText = viewContent.buttonText {
                        MomentumButton(props: MomentumButtonProps(
                            id: "welcome-button-\(state)",
                            displayName: buttonText,
                            onClick: {
                                state = state + 1
                                progressBarState = 0
                            },
                            disabled: goal.isEmpty,
                            icon: nil,
                            variant: MomentumButtonVariant.filled
                        ))
                    }
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


