import {request} from "@/api/request";
import {errorStore} from "@/common-stores/error-store.ts";
import {GetUserByIdRequestResponse} from "@/types/GetUserByIdRequestResponse";
import {userByIdStore} from "@/pages/users/stores/user-by-id-store.ts";
import {GetReviewsByUserIdRequestResponse} from "@/types/GetReviewsByUserIdRequestResponse.ts";
import {UserByIdData} from "@/types/GetUserByIdData.ts";
import {GetServicesByUserIdRequestResponse} from "@/types/GetServicesByUserIdRequestResponse.ts";

export const tryToGetUserById = async (id: string) => {
    const response = await request<GetUserByIdRequestResponse>(`/users/${id}`, 'GET');
    if (response.status != 200) {
        errorStore.set({shown: true, error: response.data.error});
    } else userByIdStore.set(response.data)
}

export const getMoreReviewsByUserId = async (
    prevUserData: UserByIdData | null | undefined
) => {
    if (!prevUserData) return;

    const response = await request<GetReviewsByUserIdRequestResponse>(
        `/users/${prevUserData.user.id}/reviews?cursor=${prevUserData.reviewsCursor}`,
        'GET'
    );
    if (response.status != 200) {
        errorStore.set({shown: true, error: response.error});
    } else userByIdStore.update((prev) => {
        if (!prev || !prev.user.reviews) return prev;

        return {
            ...prev,
            user: {
                ...prev.user,
                reviews: [...prev.user.reviews, ...response.data.reviews],
            },
            hasMoreReviews: response.data.hasMoreReviews,
            reviewsCursor: response.data.reviewsCursor
        }
    })
}

export const getMoreServicesByUserId = async (
    prevUserData: UserByIdData | null | undefined
) => {
    if (!prevUserData) return;

    const response = await request<GetServicesByUserIdRequestResponse>(
        `/users/${prevUserData.user.id}/services?cursor=${prevUserData.servicesCursor}`,
        'GET'
    );
    if (response.status != 200) {
        errorStore.set({shown: true, error: response.error});
    } else userByIdStore.update((prev) => {
        if (!prev || !prev.user.services) return prev;

        return {
            ...prev,
            user: {
                ...prev.user,
                services: [...prev.user.services, ...response.data.services],
            },
            hasMoreServices: response.data.hasMoreServices,
            servicesCursor: response.data.servicesCursor
        }
    })
}