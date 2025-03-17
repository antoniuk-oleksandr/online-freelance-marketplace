import type { StatusEnum } from "./StatusEnum"

export type OrderTableData = {
  id: number,
  title: string,
  status: StatusEnum,
  price: number,
  date: number,
  image: string,
}
