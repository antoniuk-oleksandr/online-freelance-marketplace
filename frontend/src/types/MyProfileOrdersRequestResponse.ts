import type { MyProfileOrdersData } from "./MyProfileOrdersData"
import type { ResponseError } from "./ResponseErrorEnum"

export type MyProfileOrdersRequestResponse = {
  data: MyProfileOrdersData,
  status: 200
} | {
  data: ResponseError,
  status: 400 | 404 | 500
}
