import {getHost} from "@/utils/utils.ts";
import axios from "axios";

export const request = async <T>(
    endpoint: string,
    method: "GET" | "POST" | "PUT" | "DELETE",
    token?: string,
    body?: any,
) => {
    const host = getHost();
    const url = `http://${host}/api/v1${endpoint}`;

    const methodFunc = method === "GET"
        ? axios.get : method === "POST"
            ? axios.post : method === "PUT"
                ? axios.put : axios.delete;

    const config = token !== undefined
        ? {headers: {Authorization: `Bearer ${token}`}}
        : undefined;

    try {
        let response = ["GET", "DELETE"].includes(method)
            ? await methodFunc(url, config)
            : await methodFunc(url, body, config);

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