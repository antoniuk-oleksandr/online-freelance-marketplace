import type { OrderByIdOverviewTabData } from "./OrderByIdOverviewTabData"
import type { ResponseError } from "./ResponseErrorEnum"

export type MyProfileOverviewRequestResponse = {
  data: OrderByIdOverviewTabData,
  status: 200
} | {
  data: ResponseError,
  status: 500 | 404 | 400
}
