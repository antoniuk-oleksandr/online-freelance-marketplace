import axios from "axios";
import {getHost} from "@/utils/utils";

export const getUserByIdRequest = async (id: string) => {
    const host = getHost();
    const url = `http://${host}/api/v1/users/${id}`;

    try {
        return await axios.get(url);
    } catch (error) {
        console.error(error);
        return {
            data: null,
            status: (error as any).response.status
        }
    }
}