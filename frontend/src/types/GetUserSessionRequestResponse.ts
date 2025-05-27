import { ResponseErrorEnum } from "@/types/ResponseErrorEnum";
import type { UserSessionData } from "./UserSessionData";

export type GetUserSessionRequestResponse = {
  data: UserSessionData,
  status: 200,
} | {
  data: {
    error: ResponseErrorEnum,
  }
  status: 400 | 404 | 500,

}
