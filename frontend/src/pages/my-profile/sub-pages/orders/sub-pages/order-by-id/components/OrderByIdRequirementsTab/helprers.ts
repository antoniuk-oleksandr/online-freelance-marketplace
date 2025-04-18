import { errorStore } from "@/common-stores/error-store"
import type { MyProfileRequirementsRequestResponse } from "@/types/MyProfileRequirementsRequestResponse"
import type { ServiceQA } from "@/types/ServiceQA";

export const makeMyProfileRequirementsRequest = (orderId: string): MyProfileRequirementsRequestResponse => {
  const response = {
    data: {
      questionsAnswers: [
        {
          question: 'What is the deadline for this project?',
          answer: 'The deadline for this project is 10 days from the date of order.',
        },
        {
          question: 'What is the budget for this project?',
          answer: 'The budget for this project is $100.',
        },
        {
          question: 'What is the deadline for this project?',
          answer: 'The deadline for this project is 10 days from the date of order.',
        },
        {
          question: 'What is the budget for this project?',
          answer: 'The budget for this project is $100.',
        },
        {
          question: 'What is the deadline for this project?',
          answer: 'The deadline for this project is 10 days from the date of order.',
        },
        {
          question: 'What is the budget for this project?',
          answer: 'The budget for this project is $100.',
        },
      ] as ServiceQA[]
    },
    status: 200,
  } as MyProfileRequirementsRequestResponse

  if (response.status !== 200) {
    errorStore.set({ shown: true, error: response.data.error });
  }

  return response;
}
