import type {SignInData} from "@/types/SignData.ts";
import {postAuthRequest} from "@/api/post-auth-request.ts";
import {ResponseErrorEnum} from "@/types/ResponseErrorEnum.ts";
import {errorStore} from "@/common-stores/error-store.ts";

export const handlePasswordResetRequest = async (
    body: SignInData,
    token: string,
    setReset: (value: boolean) => void
) => {
    const response = await postAuthRequest("reset-password", token, {password: body.password})

    if (response.status !== 200) {
        errorStore.set({shown: true, error: response.data.error});
        return;
    } else setReset(true);

    return response;
}