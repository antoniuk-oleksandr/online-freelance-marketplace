import type { ResponseError } from "./ResponseErrorEnum"

export type GetEmailAvailabilityRequestResponse = {
  data: {
    available: boolean,
  },
  status: 200
} | {
  data: ResponseError,
  status: 400 | 404 | 500
}
