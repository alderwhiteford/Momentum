//
//  Calendar.swift
//  MomentumIOS
//
//  Created by Alderr Whiteford on 8/26/24.
//

import Foundation
import SwiftUI

struct Calendar: View {
    @Binding var date: Date
    
    var body: some View {
        DatePicker(
            "Select Date",
            selection: $date,
            in: Date()...,
            displayedComponents: .date
        )
        .datePickerStyle(.automatic)
        .colorScheme(.dark)
        .accentColor(.primaryContainerOn)
    }
}
