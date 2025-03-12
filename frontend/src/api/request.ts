import { getHost } from "@/utils/utils";
import axios from "axios";

export const request = async <T>(
    endpoint: string,
    method: "GET" | "POST" | "PUT" | "DELETE",
    token?: string,
    body?: any,
) => {
    const host = getHost();
    const url = `${host}/api/v1${endpoint}`;

    const methodFunc = method === "GET"
        ? axios.get : method === "POST"
            ? axios.post : method === "PUT"
                ? axios.put : axios.delete;

    const headers: Record<string, string> = token
        ? { Authorization: `Bearer ${token}` }
        : {};

    if (body instanceof FormData) {
        headers["Content-Type"] = "multipart/form-data";
    } else if (body) {
        headers["Content-Type"] = "application/json";
    }

    try {
        let response = ["GET", "DELETE"].includes(method)
            ? await methodFunc(url, { headers })
            : await methodFunc(url, body, { headers });

        return {
            data: response.data,
            status: response.status,
        } as T
    } catch (e) {
        return {
            data: (e as any).response.data,
            status: (e as any).status,
        } as T
    }
}
