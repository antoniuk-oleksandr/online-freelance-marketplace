import {ResponseErrorEnum} from "@/types/ResponseErrorEnum";
import {OrderSubmitRequirementsData} from "@/types/OrderSubmitRequirementsData";

export type GetPublicKeyRequestResponse = {
    data: OrderSubmitRequirementsData
    status: 200,
} | {
    data: {
        error: ResponseErrorEnum,
    }
    status: 404 | 500,
}
