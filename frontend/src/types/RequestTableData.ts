import type { StatusEnum } from "./StatusEnum";

export type RequestTableData = {
  id: string,
  image: string,
  title: string,
  status: StatusEnum,
  price: number,
  customerFirstName: string,
  customerSurname: string,
  customerAvatar: string,
  date: number,
}
