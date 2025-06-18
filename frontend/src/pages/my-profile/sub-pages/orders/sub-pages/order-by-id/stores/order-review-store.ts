import type { OrderReview } from "@/types/OrderReview";
import { writable } from "svelte/store";

export const orderReviewStore = writable<undefined | OrderReview>()
