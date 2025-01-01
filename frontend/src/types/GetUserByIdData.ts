import type {User} from "@/types/User.ts";

export type UserByIdData = {
    user: User,
    hasMoreReviews: boolean
    reviewsCursor?: string
    hasMoreServices: boolean
    servicesCursor?: string
}