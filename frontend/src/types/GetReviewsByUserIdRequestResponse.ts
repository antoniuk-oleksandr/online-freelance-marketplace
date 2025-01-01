import {Review} from "@/types/Review.ts";
import {ResponseErrorEnum} from "@/types/ResponseErrorEnum.ts";

export type GetReviewsByUserIdRequestResponse = {
    data: {
        reviews: Review[],
        hasMoreReviews: boolean,
        reviewsCursor?: string
    },
    status: 200
} | {
    error: ResponseErrorEnum,
    status: 404 | 500
}