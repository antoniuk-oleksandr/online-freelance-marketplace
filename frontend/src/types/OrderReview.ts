export type OrderReview = {
  content: string,
  rating: number,
  customer: {
    id: number,
    firstName: string,
    surname: string,
    avatar: string | null,
    username: string
  }
}
