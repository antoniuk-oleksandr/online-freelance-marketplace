import type { ResponseError } from "./ResponseErrorEnum"

export type PostRequirementsByOrderIdRequestResponse = {
  data: {
    message: string,
  }
  status: 200
} | {
  data: ResponseError,
  status: 400 | 401 | 404 | 500
}
