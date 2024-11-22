import {ResponseError, ResponseErrorEnum} from "@/types/ResponseErrorEnum.ts";
import {PostConfirmEmailRequestResponse} from "@/types/PostConfirmEmailRequestResponse.ts";
import {ResponseMessageEnum} from "@/types/ResponseMessageEnum.ts";
import {ResponseMessage} from "@/types/ResponseMessage.ts";

export const handleConfirmEmailResponse = (
    response: PostConfirmEmailRequestResponse,
    setError: (error: ResponseErrorEnum | undefined | null) => void,
) => {
    const {data, status} = response;
    const errorMessage = (data as ResponseError).error;
    const responseMessage = (data as ResponseMessage).message;

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