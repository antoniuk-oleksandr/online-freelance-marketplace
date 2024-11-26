import {ResponseError} from "@/types/ResponseErrorEnum.ts";
import {ResponseMessage} from "@/types/ResponseMessage.ts";

export type PostAuthRequestResponse = {
    data: ResponseError,
    status: number,
}