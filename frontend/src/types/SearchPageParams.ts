import { SortType } from "@/types/SortType";
import { OrderType } from "@/types/OrderType";

export type SearchPageParams = {
    query: string | null,
    sort: SortType | null,
    order: OrderType | null,
    skill: string[],
    language: string[],
    category: string[],
    priceFrom: string | null,
    priceTo: string | null,
    levelFrom: string | null,
    levelTo: string | null,
    ratingFrom: string | null,
    ratingTo: string | null,
    deliveryTimeFrom: string | null,
    deliveryTimeTo: string | null,
}
