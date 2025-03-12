import {ResponseErrorEnum} from "@/types/ResponseErrorEnum.ts";

export type GetPublicKeyRequestResponse = {
    data: string
    status: 200,
} | {
    data: {
        error: ResponseErrorEnum,
    }
    status: 404 | 500,
}