import {getServiceByIdRequest} from "@/api/get-service-by-id-request.ts";
import {getFile} from "@/utils/utils.ts";
import {Service} from "@/types/Service.ts";

export const tryToGetServiceById = async (
    id: string,
    setService: (service: Service) => void,
    setStatus: (status: number) => void
): Promise<void> => {
    const {data, status} = await getServiceByIdRequest(id);

    if (status === 200 && data) {
        setService(data);
    }

    setStatus(status);
}