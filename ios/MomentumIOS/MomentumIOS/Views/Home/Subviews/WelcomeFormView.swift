//
//  SettingAbstractView.swift
//  MomentumIOS
//
//  Created by Alder Whiteford on 9/4/24.
//

import Foundation
import SwiftUI

struct WelcomeFormView: View {
    // Form state:
    @Binding var state: Int
    
    // Progress bar state:
    @State private var progressBarState: CGFloat = 0.0
    
    // Form variables:
    @State var goal: String = ""
    @State var estimatedCompletionAt: Date = Date.now
    @State var theWhy: String = ""
    @State var whenSuccess: SuccessIndicators = SuccessIndicators.target
    
    // The Content for each of the pages:
    let pagesContent: [WelcomeForm] = [
        WelcomePageTwoContent,
        WelcomePageThreeContent,
        WelcomePageFourContent,
        WelcomePageFiveContent,
    ]
    
    func onSubmit() -> Void {
        print(goal)
        print(estimatedCompletionAt)
        print(theWhy)
        print(whenSuccess)
    }
    
    func nextDisabled() -> Bool {
        if state == 1 {
            return goal.isEmpty
        }
        if state == 3 {
            return theWhy.isEmpty
        }
        
        return false
    }
    
    var body: some View {
        // The Content for this view
        let viewContent = pagesContent[state - 1]
        
        // The state associated with this view
        let viewState: Any = [$goal, $estimatedCompletionAt, $theWhy, $whenSuccess][state - 1]
                
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
                            color: nil,
                            alignment: nil,
                            variant: MomentumButtonVariant.none))
                    }
                    
                    Text(viewContent.title)
                        .font(.lightAppSubheader)
                    
                    Spacer().frame(maxHeight: 24)
                    
                    Text(viewContent.subTitle)
                        .font(.appCaption)
                    
                    Spacer().frame(maxHeight: 24)
                    
                    switch (viewContent.formInput) {
                        case WelcomeFormInput.text:
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
                        case WelcomeFormInput.select:
                            MomentumButton(props: MomentumButtonProps(
                                id: "achieve-success",
                                displayName: "When I complete or achieve something",
                                onClick: {
                                    whenSuccess = SuccessIndicators.completion
                                    onSubmit()
                                },
                                disabled: false,
                                icon: Image(.check),
                                color: .container,
                                alignment: .leading,
                                variant: .filled)
                            )
                            MomentumButton(props: MomentumButtonProps(
                                id: "target-success",
                                displayName: "When I reach a target metric",
                                onClick: {
                                    whenSuccess = SuccessIndicators.target
                                    onSubmit()
                                },
                                disabled: false,
                                icon: Image(.target),
                                color: .container,
                                alignment: .leading,
                                variant: .filled)
                            )
                            MomentumButton(props: MomentumButtonProps(
                                id: "satisfy-success",
                                displayName: "When I feel satisfied",
                                onClick: {
                                    whenSuccess = SuccessIndicators.satisfied
                                    onSubmit()
                                },
                                disabled: false,
                                icon: Image(.satisfy),
                                color: .container,
                                alignment: .leading,
                                variant: .filled)
                            )
                        case WelcomeFormInput.date:
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
                            disabled: nextDisabled(),
                            icon: nil,
                            color: nil,
                            alignment: nil,
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
            .onAppear {
                print(state)
            }
        }
    }    
}

struct PreviewProvider: View {
    @State private var settingState: Int = 4
    var body: some View {
        WelcomeFormView(state: $settingState)
    }
}

#Preview {
    PreviewProvider()
}
