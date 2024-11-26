import type {SignInData} from "@/types/SignData.ts";
import {postAuthRequest} from "@/api/post-auth-request.ts";
import {ResponseErrorEnum} from "@/types/ResponseErrorEnum.ts";

export const handlePasswordResetRequest = async (
    body: SignInData,
    token: string,
    setError: (value: undefined | null | ResponseErrorEnum) => void,
    setReset: (value: boolean) => void
) => {
    const response = await postAuthRequest("reset-password", token, {password: body.password})

    if (response.status !== 200) {
        setError(response.data.error);
        return;
    } else setReset(true);

    return response;
}