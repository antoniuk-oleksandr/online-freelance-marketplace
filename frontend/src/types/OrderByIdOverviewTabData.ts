import type { StatusEnum } from "./StatusEnum";

export type OrderByIdOverviewTabData = {
  id: number,
  deliveryDate: number,
  createdAt: number,
  totalPrice: number,
  status: StatusEnum,
  subtotal: number,
  serviceFee: number,
  service: {
    title: string,
    image: string,
    package: {
      name: string,
      description: string,
      deliveryTime: number,
    },
  },
  freelancer: {
    username: string,
    id: number,
    avatar: string,
  }
}
