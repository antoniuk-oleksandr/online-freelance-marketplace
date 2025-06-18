import type { DeliveryData } from "./DeliveryData"
import type { ResponseError } from "./ResponseErrorEnum"

export type MyProfileDeliveryRequestResponse = {
  data: DeliveryData,
  status: 200
} | {
  data: ResponseError,
  status: 400 | 404 | 500
}
