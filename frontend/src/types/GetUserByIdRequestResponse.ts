import {ResponseError} from "@/types/ResponseErrorEnum.ts";
import {UserByIdData} from "@/types/GetUserByIdData.ts";

export type GetUserByIdRequestResponse = {
    data: UserByIdData,
    status: 200
} | {
    data: ResponseError,
    status: 429 | 404
}