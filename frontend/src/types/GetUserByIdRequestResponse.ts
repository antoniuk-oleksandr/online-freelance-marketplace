import {ResponseErrorEnum} from "@/types/ResponseErrorEnum.ts";
import {UserByIdData} from "@/types/GetUserByIdData.ts";

export type GetUserByIdRequestResponse = {
    data: UserByIdData,
    status: 200
} | {
    error: ResponseErrorEnum,
    status: 429 | 404
}