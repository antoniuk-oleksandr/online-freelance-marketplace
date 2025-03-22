import type { MyProfileRequestsData } from "./MyProfileRequestsData"
import type { ResponseError } from "./ResponseErrorEnum"

export type MyProfileRequestsRequestResponse = {
  data: MyProfileRequestsData,
  status: 200
} | {
  data: ResponseError,
  status: 400 | 404 | 500
}

