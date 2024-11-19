import {getHost} from "@/utils/utils";
import axios from "axios";
import type {Service} from "@/types/Service.ts";

type ServiceByIdRequestResponse = {
    data: Service | null,
    status: number
}

export const getServiceByIdRequest = async (id: string)
    : Promise<ServiceByIdRequestResponse> => {
    const host = getHost();
    const url = `http://${host}/api/v1/services/${id}`;

    try {
        const {data, status} = await axios.get(url);

        return {
            data: data as Service,
            status
        };
    } catch (error) {
        return {
            data: null,
            status: (error as any).response.status
        }
    }
}