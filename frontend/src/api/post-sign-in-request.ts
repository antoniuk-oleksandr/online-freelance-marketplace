import {SignInData} from "@/types/SignData.ts";
import {getHost} from "@/utils/utils.ts";
import axios from "axios";
import {TokenResponse} from "@/types/TokenResponse.ts";
import {ResponseError} from "@/types/ResponseErrorEnum.ts";

type PostSignInRequestResponse = {
    data: TokenResponse | ResponseError,
    status: number
}

export const postSignInRequest = async (
    data: SignInData
) : Promise<PostSignInRequestResponse> => {
    const host = getHost();
    const url = `http://${host}/api/v1/auth/sign-in`;

    try {
        const response = await axios.post(url, {
            usernameOrEmail: data.usernameOrEmail,
            password: data.password
        });
        return {
            data: response.data as TokenResponse,
            status: response.status
        }
    }
    catch (e){
        return {
            data: (e as any).response.data,
            status: (e as any).response.status
        }
    }
}