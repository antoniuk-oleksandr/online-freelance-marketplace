import type { MyProfileEditPageFormData } from "./MyProfileEditPageFormData"
import type { ResponseError } from "./ResponseErrorEnum"

export type MyProfileDataRequestResponse = {
  status: 200,
  data: MyProfileEditPageFormData
} | {
  status: 400 | 404 | 500,
  data: ResponseError,
}
