import {BestFreelancer} from "@/types/BestFreelancer.ts";
import {BestReview} from "@/types/BestReview.ts";
import {BestService} from "@/types/BestService.ts";
import {KeyMetrics} from "@/types/KeyMetrics.ts";

export type HomeData = {
    bestFreelancers: BestFreelancer[],
    bestReviews: BestReview[],
    bestServices: BestService[]
    keyMetrics: KeyMetrics,
}