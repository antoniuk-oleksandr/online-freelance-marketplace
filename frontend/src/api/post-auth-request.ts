import {getHost} from "@/utils/utils.ts";
import axios from "axios";
import {PostAuthRequestResponse} from "@/types/PostAuthRequestResponse.ts";

export const postAuthRequest = async (
    endpoint: string,
    token?: string,
    body?: any,
): Promise<PostAuthRequestResponse> => {
    const host = getHost();
    const url = `http://${host}/api/v1/auth/${endpoint}`;

    const config = token !== undefined ? {
        headers: {
            Authorization: `Bearer ${token}`
        }
    } : undefined;

    try {
        const response = await axios.post(url, body, config);

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