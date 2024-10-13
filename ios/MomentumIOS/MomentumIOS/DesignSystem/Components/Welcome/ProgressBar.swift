//
//  ProgressBar.swift
//  MomentumIOS
//
//  Created by Alder Whiteford on 8/21/24.
//

import Foundation
import SwiftUI

enum StepError: Error {
    case invalidStep
}

struct BezierPathShape: Shape {
    var progress: CGFloat

    func path(in rect: CGRect) -> Path {
        var path = Path()
        let endX = rect.width * progress
        
        path.move(to: CGPoint(x: 0, y: rect.midY))
        path.addLine(to: CGPoint(x: endX, y: rect.midY))

        return path
    }
    
    var animatableData: CGFloat {
        get { progress }
        set { progress = newValue }
    }
}

struct ProgressBar: View {
    @Binding var progress: CGFloat
    var currentStep: Int
    
    var numberOfSteps: Int
    var title: String
    
    var body: some View {
        VStack (alignment: .leading, spacing: 10) {
            HStack {
                ForEach(0..<self.numberOfSteps, id: \.self) { index in
                    VStack{
                        if index == currentStep {
                            BezierPathShape(progress: progress * CGFloat(1))
                                .stroke(Color("MomentumPrimary"), lineWidth: 8)
                                .frame(height: 8)
                                .onAppear {
                                    animateProgress()
                                }
                                .onChange(of: currentStep) { () in
                                    resetAndAnimateProgress()
                                }
                        }
                        
                    }
                        .frame(maxWidth: .infinity)
                        .frame(height: 8)
                        .background(Color(index < currentStep ? "MomentumPrimary" : "Surface-Container-2"))
                        .clipShape(RoundedRectangle(cornerRadius: 8))
                }
            }
            HStack {
                VStack (alignment: .leading) {
                    Text(self.title)
                        .foregroundColor(Color("MomentumPrimary"))
                        .font(.appCaptionThree)
                }
                .frame(maxWidth: .infinity, alignment: .leading)
                VStack (alignment: .trailing) {
                    Text("Step \(self.currentStep + 1)/\(self.numberOfSteps)")
                        .foregroundColor(Color("Background-On-Dark"))
                        .font(.appCaptionThree)
                }
                .frame(maxWidth: .infinity, alignment: .trailing)
            }
        }
        .frame(maxWidth: .infinity)
    }
    
    private func resetAndAnimateProgress() {
        progress = 0.0
        animateProgress()
    }
    
    private func animateProgress() {
        withAnimation(.easeInOut(duration: 1.0)) {
            progress = 1.0
        }
    }
}
