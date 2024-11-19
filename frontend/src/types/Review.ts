import { User } from "@/types/User";
import { UserReview } from "@/types/UserReview";

export type Review = {
  id: number;
  content: string;
  rating: number;
  createdAt: number;
  endedAt: number;
  customer: User;
  service: UserReview;
};
