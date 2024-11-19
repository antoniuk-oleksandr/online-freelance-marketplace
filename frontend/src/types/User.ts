import {Review} from "@/types/Review";
import {UserService} from "@/types/UserService";
import {Language} from "@/types/Language";
import {Skill} from "@/types/Skill";

export type User = {
    id: number,
    firstName: string,
    surname: string,
    username: string,
    rating: number | null,
    level: number,
    reviewsCount: number | null,
    avatar: string | null,
    createdAt: Date,
    about: string | null,
    languages: Language[] | null,
    skills: Skill[] | null,
    reviews: Review[] | null,
    services: UserService[] | null,
};
