import type { RefreshTokenRequestResponse } from "@/types/RefreshTokenRequestResponse";
import { ResponseErrorEnum, type ResponseError } from "@/types/ResponseErrorEnum";
import { getHost } from "@/utils/utils";
import axios from "axios";
import Cookies from "js-cookie";
import { jwtDecode } from "jwt-decode";

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

    console.log("Requesting", method, url, body);
    
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

const getAccessToken = async (host: string): Promise<RefreshTokenRequestResponse> => {
    const accessToken = Cookies.get("accessToken");
    if (accessToken) {
        const decodedJWT = jwtDecode(accessToken);
        const expired = new Date(decodedJWT.exp as number * 1000) < new Date()
        if (!expired) return { data: { accessToken }, status: 200 };
    }

    const response = await tryToRefreshToken(host);
    if (response.status === 200) {
        const decodedData = jwtDecode(response.data.accessToken);
        Cookies.set("accessToken", response.data.accessToken, { expires: decodedData.exp });
        return response;
    } else {
        return response;
    }
}

const tryToRefreshToken = async (host: string): Promise<RefreshTokenRequestResponse> => {
    const refreshToken = Cookies.get("refreshToken");
    if (!refreshToken) return {
        data: { error: ResponseErrorEnum.ExpiredToken },
        status: 401,
    }

    const url = `${host}/api/v1/auth/refresh-token`;

    try {
        const headers = { Authorization: `Bearer ${refreshToken}` };
        return await axios.post(url, undefined, { headers });
    } catch (e) {
        return {
            data: (e as any).response.data,
            status: (e as any).status,
        }
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
