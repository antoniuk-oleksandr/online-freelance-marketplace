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
        data.images = data.images.map((image) => getFile(image));
        data.freelancer.avatar = getFile(data.freelancer.avatar);
        data.reviews = data.reviews && data.reviews.map((review) => {
            review.customer.avatar = getFile(review.customer.avatar);
            return review;
        });

        setService(data);
    }

    setStatus(status);
}