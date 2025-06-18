import type { OrderReview } from "./OrderReview"
import type { ResponseError } from "./ResponseErrorEnum"

export type PostOrderReviewRequestResponse = {
  data: OrderReview,
  status: 201
} | {
  data: ResponseError,
  status: 400 | 401 | 404 | 500
}
