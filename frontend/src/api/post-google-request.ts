import {getHost} from "@/utils/utils.ts";
import axios from "axios";
import {TokenResponse} from "@/types/TokenResponse.ts";
import {ResponseError} from "@/types/ResponseErrorEnum.ts";

type PostGoogleRequestResponse = {
    data: TokenResponse | ResponseError,
    status: number
}

export const postGoogleRequest = async (
    code: string
): Promise<PostGoogleRequestResponse> => {
    const host = getHost();
    const url = `http://${host}/api/v1/auth/google`;

    try {
        const response = await axios.post(url, {code});
        return {
            data: response.data,
            status: response.status
        }
    } catch (e) {
        return {
            data: (e as any).response.data,
            status: (e as any).response.status
        }
    }
}