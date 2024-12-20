import type {User} from "@/types/User.ts";
import {ResponseError} from "@/types/ResponseErrorEnum.ts";

export type GetUserByIdRequestResponse = {
    data: User,
    status: 200
} | {
    data: ResponseError,
    status: 429 | 404
}