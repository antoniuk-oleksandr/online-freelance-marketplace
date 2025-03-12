import {ResponseErrorEnum} from "@/types/ResponseErrorEnum.ts";
import {OrderSubmitRequirementsData} from "@/types/OrderSubmitRequirementsData.ts";

export type GetPublicKeyRequestResponse = {
    data: OrderSubmitRequirementsData
    status: 200,
} | {
    data: {
        error: ResponseErrorEnum,
    }
    status: 404 | 500,
}