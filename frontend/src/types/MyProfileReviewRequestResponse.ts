import type { OrderReview } from "./OrderReview"
import type { ResponseErrorEnum } from "./ResponseErrorEnum"
import type { StatusEnum } from "./StatusEnum"


export type MyProfileReviewRequestResponse = {
  data: OrderReview,
  status: 200,
} | {
  data: {
    error: ResponseErrorEnum,
    status: StatusEnum | undefined
  },
  status: 400 | 401 | 404 | 500,
}
