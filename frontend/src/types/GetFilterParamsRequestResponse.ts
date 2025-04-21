import { type ResponseError } from "@/types/ResponseErrorEnum";
import type { FilterParams } from "./FilterParams";

export type GetFilterParamsRequestResponse = {
    data: FilterParams,
    status: 200,
} | {
    data: ResponseError,
    status: 404 | 500 | 401,
}

