import {ResponseError} from "@/types/ResponseErrorEnum.ts";
import {ResponseMessage} from "@/types/ResponseMessage.ts";

export type PostConfirmEmailRequestResponse = {
    data: ResponseMessage | ResponseError,
    status: number,
}