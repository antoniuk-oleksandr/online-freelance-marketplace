import { getAccessToken } from "@/utils/token-utils";
import { getHost } from "@/utils/utils";
import axios from "axios";

export const request = async <T>(
    method: "GET" | "POST" | "PUT" | "DELETE",
    endpoint: string,
    body?: any,
    token?: boolean,
) => {
    const host = getHost();
    const url = `${host}/api/v1${endpoint}`;

    let accessToken;
    if (token) {
        const response = await getAccessToken(host);
        if (response.status === 200) accessToken = response.data.accessToken;
        else return response;
    }

    const methodFunc = defineRequestMethod(method);
    const headers = defineRequestHeaders(accessToken, body);

    try {
        let response = ["GET", "DELETE"].includes(method)
            ? await methodFunc(url, { headers })
            : await methodFunc(url, body, { headers });

        return {
            data: response.data,
            status: response.status,
        } as T
    } catch (e) {
        console.error(e);
        return {
            data: (e as any).response.data,
            status: (e as any).status,
        } as T
    }
}



const defineRequestMethod = (method: "GET" | "POST" | "PUT" | "DELETE") => {
    switch (method) {
        case "GET":
            return axios.get;
        case "POST":
            return axios.post;
        case "PUT":
            return axios.put;
        case "DELETE":
            return axios.delete;
    }
}

const defineRequestHeaders = (accessToken: string | undefined, body?: any) => {
    const headers: Record<string, string> = accessToken
        ? { Authorization: `Bearer ${accessToken}` }
        : {};

    if (body instanceof FormData) {
        headers["Content-Type"] = "multipart/form-data";
    } else if (body) {
        headers["Content-Type"] = "application/json";
    }

    return headers;
}
