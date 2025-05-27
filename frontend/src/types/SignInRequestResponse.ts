import type { ResponseError } from "./ResponseErrorEnum"
import type { SignInData } from "./SignInData"

export type SignInRequestResponse = {
  data: SignInData,
  status: 200,
} | {
  data: ResponseError,
  status: 400 | 404 | 500,
}
