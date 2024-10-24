import {User} from "@/types/User";
import { Package } from "./Package";
import {Review} from "@/types/Review";

export type Service = {
    id: number,
    title: string,
    description: string,
    images: string[],
    createdAt: Date,
    rating: number,
    category: string,
    freelancer: User,
    packages: Package[],
    reviews: Review[]
}