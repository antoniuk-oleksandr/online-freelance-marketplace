import { errorStore } from "@/common-stores/error-store"
import type { MyProfileRequirementsRequestResponse } from "@/types/MyProfileRequirementsRequestResponse"
import type { ServiceQA } from "@/types/ServiceQA";

export const makeMyProfileRequirementsRequest = (orderId: string): MyProfileRequirementsRequestResponse => {
  const response = {
    data: {
      questionsAnswers: [
        {
          question: 'What is the timeline for completing the project?',
          answer: 'The project needs to be completed within 10 days from the date of order.',
        },
        {
          question: 'Do you have any design preferences or style guidelines?',
          answer: 'Yes, I’d like the design to be clean, modern, and use a natural color palette, similar to the brand’s aesthetic.',
        },
        {
          question: 'Are there any reference materials you can share?',
          answer: 'I’ll send over a few logos and branding materials that reflect the style I’m aiming for.',
        },
        {
          question: 'What is the target audience for the logo?',
          answer: 'The target audience is health-conscious individuals aged 25-40, primarily interested in wellness and sustainable living.',
        },
        {
          question: 'Will you need multiple logo variations (e.g., color, black and white, etc.)?',
          answer: 'Yes, I would like both color and monochrome versions of the logo to be delivered.',
        },
      ] as ServiceQA[],
    },
    status: 200,
  } as MyProfileRequirementsRequestResponse;

  if (response.status !== 200) {
    errorStore.set({ shown: true, error: response.data.error });
  }

  return response;
}

