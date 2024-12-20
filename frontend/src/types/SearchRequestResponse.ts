import {UserService} from "@/types/UserService.ts";
import {ResponseError} from "@/types/ResponseErrorEnum.ts";

export type SearchRequestResponse = {
    data: {
        services: null | UserService[],
        cursor: string | null,
        hasMore: boolean,
    }
    status: 200
} | {
    data: ResponseError,
    status: 429 | 500
}