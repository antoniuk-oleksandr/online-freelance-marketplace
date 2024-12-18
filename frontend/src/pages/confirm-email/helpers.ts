import {ResponseError, ResponseErrorEnum} from "@/types/ResponseErrorEnum.ts";
import {PostAuthRequestResponse} from "@/types/PostAuthRequestResponse.ts";
import {ResponseMessageEnum} from "@/types/ResponseMessageEnum.ts";
import {ResponseMessage} from "@/types/ResponseMessage.ts";

export const handleConfirmEmailResponse = (
    response: PostAuthRequestResponse,
    setError: (error: ResponseErrorEnum | undefined | null) => void,
) => {
    const {data, status} = response;
    const errorMessage = (data as ResponseError).error;
    const responseMessage = (data as unknown as ResponseMessage).message;

    if (status === 200 && responseMessage === ResponseMessageEnum.EmailConfirmed) {
        setError(null);
    } else if (errorMessage === ResponseErrorEnum.ExpiredToken) {
        setError(ResponseErrorEnum.ExpiredToken);
    } else if (errorMessage === ResponseErrorEnum.InvalidToken) {
        setError(ResponseErrorEnum.InvalidToken);
    } else {
        setError(ResponseErrorEnum.UnexpectedError);
    }
}