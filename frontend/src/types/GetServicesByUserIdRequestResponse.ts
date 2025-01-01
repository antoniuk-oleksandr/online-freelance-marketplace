import {ResponseErrorEnum} from "@/types/ResponseErrorEnum.ts";
import {UserService} from "@/types/UserService.ts";

export type GetServicesByUserIdRequestResponse = {
    data: {
        services: UserService[],
        hasMoreServices: boolean,
        servicesCursor?: string
    },
    status: 200
} | {
    error: ResponseErrorEnum,
    status: 404 | 500
}