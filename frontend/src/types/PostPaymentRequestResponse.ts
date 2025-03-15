import { ResponseErrorEnum } from "@/types/ResponseErrorEnum";

export type PostPaymentRequestResponse = {
    data: {
        success: boolean,
        orderId: number,
    },
    status: 201,
} | {
    data: {
        error: ResponseErrorEnum,
    }
    status: 404 | 500,
}
