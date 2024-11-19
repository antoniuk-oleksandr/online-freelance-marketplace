import axios from "axios";
import {getHost} from "@/utils/utils";
import type {User} from "@/types/User.ts";

export type GetUserByIdRequestResponse = {
    data: User | null,
    status: number
}

export const getUserByIdRequest = async (id: string): Promise<GetUserByIdRequestResponse> => {
    const host = getHost();
    const url = `http://${host}/api/v1/users/${id}`;

    try {
        const response = await axios.get(url);
        return {
            data: response.data as User,
            status: response.status as number
        }
    } catch (error) {
        return {
            data: null,
            status: (error as any).response.status as number
        }
    }
}