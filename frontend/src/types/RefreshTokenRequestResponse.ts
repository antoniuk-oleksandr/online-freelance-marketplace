import type { ResponseError } from "./ResponseErrorEnum";

export type RefreshTokenRequestResponse = {
  data: {
    accessToken: string,
  },
  status: 200,
} | {
  data: ResponseError,
  status: 401 | 500,
}
