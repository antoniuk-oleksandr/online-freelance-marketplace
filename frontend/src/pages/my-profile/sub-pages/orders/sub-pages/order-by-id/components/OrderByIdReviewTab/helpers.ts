import type { OrderReview } from "@/types/OrderReview";
import { writable } from "svelte/store";
import { z } from "zod";

export const reviewTabFormSchema = z.object({
  rating: z.number().min(1, { message: "Review rating is required." }).max(5),
  reviewMessage: z.string().max(2000, { message: "Review message must be less than 2000 characters." }).optional(),
});
