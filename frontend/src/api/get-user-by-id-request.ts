import axios from "axios";
import {getHost} from "@/utils/utils";

export const getUserByIdRequest = async (id: string) => {
    const host = getHost();
    const url = `http://${host}/api/users/${id}?page=0&size=30`;

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