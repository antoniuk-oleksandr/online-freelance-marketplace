import {ResponseError} from "@/types/ResponseErrorEnum.ts";
import {PostAuthRequestResponse} from "@/types/PostAuthRequestResponse.ts";
import {ResponseMessageEnum} from "@/types/ResponseMessageEnum.ts";
import {ResponseMessage} from "@/types/ResponseMessage.ts";
import {errorStore} from "@/common-stores/error-store.ts";

export const handleConfirmEmailResponse = (
    response: PostAuthRequestResponse,
) => {
    const {data, status} = response;
    const errorMessage = (data as ResponseError).error;
    const responseMessage = (data as unknown as ResponseMessage).message;

    if (status === 200 && responseMessage === ResponseMessageEnum.EmailConfirmed) {
        return;
    } else {
        errorStore.set({shown: true, error: errorMessage});
    }
}