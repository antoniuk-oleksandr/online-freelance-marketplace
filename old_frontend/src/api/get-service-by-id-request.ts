import {getHost} from "@/utils/utils";
import axios from "axios";
import {Service} from "@/types/Service";

export const getServiceByIdRequest = async (id: string) => {
    const host = getHost();
    const url = `http://${host}/api/v1/services/${id}`;

    try {
        const {data, status} = await axios.get(url);

        return {
            data: data.service as Service,
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