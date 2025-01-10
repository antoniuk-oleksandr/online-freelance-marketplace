import {ResponseError} from "@/types/ResponseErrorEnum.ts";
import {Service} from "@/types/Service.ts";

export type GetServiceByUserIdRequestResponse = {
    data: {
        service: Service,
        hasMoreReviews: boolean
        reviewsCursor?: string
    },
    status: 200
} | {
    data: ResponseError,
    status: 404 | 500
}