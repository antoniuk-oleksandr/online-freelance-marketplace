import type { OrderByIdDeliveryTabData } from "./OrderByIdDeliveryTabData"
import type { ResponseError } from "./ResponseErrorEnum"

export type MyProfileDeliveryRequestResponse = {
  data: OrderByIdDeliveryTabData,
  status: 200
} | {
  data: ResponseError,
  status: 400 | 404 | 500
}
