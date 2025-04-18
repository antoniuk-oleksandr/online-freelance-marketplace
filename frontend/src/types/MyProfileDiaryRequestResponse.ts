import type { DiaryItem } from "./DiaryItem"
import type { ResponseError } from "./ResponseErrorEnum"

export type MyProfileDiaryRequestResponse = {
  data: {
    dates: DiaryItem[],
    firstDate: number,
    lastDate: number,
  },
  status: 200
} | {
  data: ResponseError,
  status: 500 | 404 | 400
}
