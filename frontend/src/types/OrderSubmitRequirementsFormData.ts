export type OrderSubmitRequirementsFormData = {
  answers: {
    content: string,
    questionId: number,
  }[],
  customerMessage: string,
  files: File[],
  orderId: string,
}
