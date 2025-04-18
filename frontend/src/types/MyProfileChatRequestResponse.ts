import type { ChatData } from "./ChatData"
import type { ResponseError } from "./ResponseErrorEnum"

export type MyProfileChatRequestResponse = {
  data: ChatData,
  status: 200,
} | {
  data: ResponseError,
  status: 404 | 400 | 500,
}
