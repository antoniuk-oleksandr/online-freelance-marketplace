import type { ResponseError } from "./ResponseErrorEnum"
import type { ServiceQA } from "./ServiceQA"

export type MyProfileRequirementsRequestResponse = {
  data: {
    questionsAnswers: ServiceQA[]
  },
  status: 200,
} | {
  data: ResponseError,
  status: 400 | 404 | 500,
}
