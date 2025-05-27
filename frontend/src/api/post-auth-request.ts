import { getHost } from "@/utils/utils";
import axios from "axios";
import { PostAuthRequestResponse } from "@/types/PostAuthRequestResponse";

export const postAuthRequest = async (
    endpoint: string,
    token?: boolean | string,
    body?: any,
): Promise<PostAuthRequestResponse> => {
    const host = getHost();
    const url = `${host}/api/v1/auth/${endpoint}`;

    try {
        const headers = getRequestHeaders(token);
        const response = await axios.post(url, body, {
            withCredentials: token === true, headers
        });

        return {
            data: response.data,
            status: response.status,
        }
    } catch (e) {
        return {
            data: (e as any).response.data,
            status: (e as any).status,
        }
    }
}

const getRequestHeaders = (token?: boolean | string) => {
    if (typeof token === 'string') {
        return { Authorization: `Bearer ${token}` }
    }
    return {}
}
