import {ResponseError} from "@/types/ResponseErrorEnum.ts";
import {RestrictedService} from "@/types/RestrictedService.ts";

export type GetRestrictedServiceByIdResponse = {
    data: RestrictedService,
    status: 200
} | {
    data: ResponseError,
    status: 404 | 500
}