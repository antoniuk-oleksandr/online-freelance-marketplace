import type { ResponseErrorEnum } from "./ResponseErrorEnum"

export type PostSignOutRequestResponse = {
  data: { message: string }
  status: 200
} | {
  data: {
    error: ResponseErrorEnum,
  }
  status: 400 | 500,
}
