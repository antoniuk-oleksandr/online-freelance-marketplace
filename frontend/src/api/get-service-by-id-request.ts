import {getHost} from "@/utils/utils";
import axios from "axios";
import {Service} from "@/types/Service";

export const getServiceByIdRequest = async (id: string) => {
    const host = getHost();
    const url = `http://${host}/api/services/${id}?page=0&size=30`;

    try {
        const {data, status} = await axios.get(url);

        return {
            data: data as Service,
            status
        };
    } catch (error) {
        console.error(error);
        return {
            data: null,
            status: (error as any).response.status
        }
    }
}