import type { MyProfileServicesData } from "./MyProfileServicesData"
import type { ResponseError } from "./ResponseErrorEnum"

export type MyProfileServicesRequestResponse = {
  data: MyProfileServicesData,
  status: 200
} | {
  data: ResponseError,
  status: 400 | 404 | 500
}
