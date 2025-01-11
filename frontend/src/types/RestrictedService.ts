import {Package} from "./Package";

export type RestrictedService = {
    id: number,
    title: string,
    image: string,
    rating: number,
    reviewCount: number,
    packages: Package[],
};
