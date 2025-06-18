import type { ResponseError } from "./ResponseErrorEnum"

export type PutUpdateOrderStatusByOrderIdRequestResponse = {
  data: {
    success: true,
  },
  status: 200
} | {
  data: ResponseError,
  status: 400 | 401 | 404 | 500
}
