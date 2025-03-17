import type { OrderTableData } from "./OrderTableData"

export type MyProfileOrdersData = {
  orders: OrderTableData[],
  cursor: string | null,
  hasMore: boolean
  totalPages: number
}
