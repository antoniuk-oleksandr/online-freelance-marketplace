import {ResponseError} from "@/types/ResponseErrorEnum.ts";
import {HomeData} from "@/types/HomeData.ts";

export type GetHomeDataRequestResponse = {
    data: HomeData,
    status: 200
} | {
    data: ResponseError,
    status: 500,
}