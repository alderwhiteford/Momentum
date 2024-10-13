//
//  Welcome.swift
//  MomentumIOS
//
//  Created by Alder Whiteford on 9/5/24.
//

import Foundation

struct WelcomeStepTag: Identifiable {
    let title: String
    let subTitle: String
    let icon: String
    var id: String { title }
}

struct WelcomeInitial {
    let title: String
    let stepTags: [WelcomeStepTag]
    let buttonText: String
}

enum WelcomeFormInput {
    case text
    case date
    case select
}

struct WelcomeForm {
    let title: String
    let subTitle: String
    let infoDrawer: InfoDrawerProps
    let canGoBack: Bool
    let buttonText: String?
    let formInput: WelcomeFormInput
    let formInputText: String?
    let formInputSubText: String?

}

var WelcomePageOneContent: WelcomeInitial = WelcomeInitial(
    title: "Welcome to Momentum.",
    stepTags: [
        WelcomeStepTag(title: "Set", subTitle: "Set your goal the right way using a psychology-backed process.", icon: "flags"),
        WelcomeStepTag(title: "Plan", subTitle: "Break up your goal into smaller milestones, tasks, and habits.", icon: "event"),
        WelcomeStepTag(title: "Manage", subTitle: "Manage tasks and habits daily, making your goal actionable.", icon: "checklist"),
        WelcomeStepTag(title: "Reflect", subTitle: "Celebrate wins and note obstacles with daily check-ins.", icon: "cognition"),
        WelcomeStepTag(title: "Track", subTitle: "Track your progress and get actionable insights and feedback.", icon: "changes")
    ],
    buttonText: "Continue"
)

var WelcomePageTwoContent: WelcomeForm = WelcomeForm(
    title: "What is your goal?",
    subTitle: "Simply writing down your goal drastically improves your chances of success.",
    infoDrawer: InfoDrawerProps(
        id: "welcome-two-info",
        title: "Goal Setting Tips",
        header: "Choosing your goals",
        subheader: "Setting the right goal is the first step on your path to fulfillment. Here’s how to find your direction:",
        sections: [
            InfoDrawerSection(
                title: "Align with Your Passions",
                description: "Consider what excites you and aligns with your core values. Whether improving a skill, advancing in your career, or contributing to your community, your goals should reflect what truly matters to you."
            ),
            InfoDrawerSection(
                title: "Dream Big, Start Small",
                description: "Aim for ambitious goals with achievable steps. This balance ensures your goals are both inspiring and reachable."
            ),
            InfoDrawerSection(
                title: "Be Unique",
                description: "Let your goals be inspired by others, but ensure they are tailored to your aspirations and situation. Your goals should be a reflection of your personal journey, not a Content of someone else's."
            ),
            InfoDrawerSection(
                title: "Stay Flexible",
                description: "Remember, it’s okay to adjust your goals as you grow and explore new interests. Life is a journey of discovery."
            ),
        ],
        prompts: InfoDrawerPrompts(
            bulleted: false,
            questions: [
                "Which activities or causes resonate with you? How can your goals increase your involvement?",
                "Which skills do you want to improve or learn? How would mastering them change your life or career?",
                "Consider your key milestones. How do they fit into your long-term vision?",
                "What impact do you aim for in your community or the world? How will your goals facilitate this impact?"
            ]
        )
    ),
    canGoBack: false,
    buttonText: "Continue",
    formInput: WelcomeFormInput.text,
    formInputText: "Goal",
    formInputSubText: "Try to describe your goal in one sentence. Keep this general and high-level to begin with."
)

var WelcomePageThreeContent: WelcomeForm = WelcomeForm(
    title: "When would you like to achieve this?",
    subTitle: "When would you like to achieve this?",
    infoDrawer: InfoDrawerProps(
        id: "welcome-three-info",
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
    ),
    canGoBack: true,
    buttonText: "Continue",
    formInput: WelcomeFormInput.date,
    formInputText: nil,
    formInputSubText: "Try to be as specific as possible, but you can always change this later if needed."
)

var WelcomePageFourContent: WelcomeForm = WelcomeForm(
    title: "Why is this goal important to you?",
    subTitle: "Identifying your ‘why’ will help motivate you to overcome obstacles.",
    infoDrawer: InfoDrawerProps(
        id: "welcome-four-info",
        title: "Defining your why",
        header: "Start with why",
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
    ),
    canGoBack: true,
    buttonText: "Continue",
    formInput: WelcomeFormInput.text,
    formInputText: "Why",
    formInputSubText: "Write one or two sentences about why accomplishing your goal would be meaningful to you."
)

var WelcomePageFiveContent: WelcomeForm = WelcomeForm(
    title: "How will you know when you've been successful?",
    subTitle: "Every goal is different, so consider what success will look like for you.",
    infoDrawer: InfoDrawerProps(
        id: "welcome-four-info",
        title: "Defining your why",
        header: "Start with why",
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
    ),
    canGoBack: true,
    buttonText: nil,
    formInput: WelcomeFormInput.select,
    formInputText: nil,
    formInputSubText: nil
)
