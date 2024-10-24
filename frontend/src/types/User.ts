import {Review} from "@/types/Review";
import {UserService} from "@/types/UserService";

export type User = {
    id: number,
    firstName: string,
    surname: string,
    rating: number,
    level: number,
    reviewsCount: number,
    avatar: string | null,
    createdAt: Date,
    about: string | null,
    languages: string[],
    skills: string[],
    reviews: Review[],
    services: UserService[],
}