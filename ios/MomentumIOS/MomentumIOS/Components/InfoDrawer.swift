import SwiftUI

struct InfoDrawerSection {
    let title: String
    let description: String
}

struct InfoDrawerPrompts {
    let bulleted: Bool
    let questions: [String]
}

struct InfoDrawerProps {
    let id: String
    let title: String
    let header: String
    let subheader: String
    let sections: [InfoDrawerSection]
    let prompts: InfoDrawerPrompts?
}

struct DrawerContent: View {
    var props: InfoDrawerProps
    
    var body: some View {
        VStack(alignment: .leading) {
            HStack {
                Text(props.title)
                    .font(.appCaptionTwo)
                    .foregroundColor(.momentumPrimary)
            }
            .frame(maxWidth: .infinity, alignment: .center)
            
            Spacer().frame(maxHeight: 30)
            
            VStack(alignment: .leading, spacing: 8) {
                Text(props.header)
                    .font(.appBody)
                
                Spacer().frame(maxHeight: 4)
                
                Text(props.subheader)
                    .font(.appCaptionTwo)
                
                Spacer().frame(maxHeight: 4)
                
                ForEach(props.sections.indices, id: \.self) { index in
                    Text("\(props.sections[index].title): ")
                        .font(.boldAppCaptionTwo)
                    + Text(props.sections[index].description)
                        .font(.appCaptionTwo)
                    
                    Spacer().frame(maxHeight: 4)
                }
                
                if let prompts = props.prompts {
                    Text("Prompts:")
                        .font(.boldAppCaptionTwo)
                    ForEach(prompts.questions.indices, id: \.self) { index in
                        HStack (alignment: .top) {
                            Text(prompts.bulleted ? "•" : "\(index + 1).")
                            Text(prompts.questions[index])
                        }
                        .font(.appCaptionTwo)
                    }
                }
            }
        }
        .padding(16)
        .frame(maxWidth: .infinity, maxHeight: .infinity, alignment: .top)
        .foregroundColor(.backgroundOn)
        .background(Color.surfaceContainer3)
        .presentationDragIndicator(.visible)
    }
}

struct InfoDrawer: View {
    @State var isPresented: Bool = false
    
    var props: InfoDrawerProps
    
    var body: some View {
        VStack {
            MomentumButton(props: MomentumButtonProps(
                id: "display-drawer",
                displayName: "Goal Setting Tips",
                onClick: {
                    isPresented = true
                },
                disabled: false,
                icon: Image(.information),
                variant: MomentumButtonVariant.none
            ))
            }
            .sheet(isPresented: $isPresented) {
                DrawerContent(props: props)
                    .frame(maxWidth: .infinity, maxHeight: .infinity, alignment: .top)
            }
            .frame(maxWidth: .infinity, alignment: .center)
    }
}
