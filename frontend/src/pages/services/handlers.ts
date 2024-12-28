import type {UpdateFunc} from "@/types/UpdateFunc.ts";
import type {GetUserByIdRequestResponse} from "@/types/GetServiceByIdRequestResponse.ts";
import {request} from "@/api/request.ts";
import {GetReviewsByServiceIdResponse} from "@/types/GetReviewsByServiceIdResponse.ts";
import {errorStore} from "@/common-stores/error-store.ts";
import {Review} from "@/types/Review.ts";

export const handleShowMoreFreelancesButtonClick = async (
    setServiceResponse: UpdateFunc<GetUserByIdRequestResponse | undefined>
) => {
    setServiceResponse(async (prev) => {
        if (!prev || prev.status !== 200) return prev;

        const {id} = prev.data.service;
        const {reviewsCursor} = prev.data;

        const link = `/freelances/${id}/reviews?reviewsCursor=${reviewsCursor}`;
        const response = await request<GetReviewsByServiceIdResponse>(link, "GET");

        if (response.status !== 200) {
            errorStore.set({shown: true, error: response.error});
            return undefined;
        }

        const {reviews: newReviews, hasMoreReviews, reviewsCursor: newReviewsCursor} = response.data;

        return {
            ...prev,
            data: {
                ...prev.data,
                service: {
                    ...prev.data.service,
                    reviews: [...prev.data.service.reviews, ...newReviews],
                },
                hasMoreReviews,
                reviewsCursor: newReviewsCursor,
            },
        };
    });
};
