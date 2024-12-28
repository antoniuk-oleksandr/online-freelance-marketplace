import {ResponseErrorEnum} from "@/types/ResponseErrorEnum.ts";
import {Service} from "@/types/Service.ts";

export type GetUserByIdRequestResponse = {
    data: {
        service: Service,
        hasMoreReviews: boolean
        reviewsCursor?: string
    },
    status: 200
} | {
    error: ResponseErrorEnum,
    status: 404 | 500
}