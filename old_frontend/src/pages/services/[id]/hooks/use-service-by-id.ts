import {useRouter} from "next/router";
import {useEffect, useState} from "react";
import {Service} from "@/types/Service";
import {getServiceByIdRequest} from "@/api/get-service-by-id-request";
import {getFile} from "@/utils/utils";

export const useServiceById = () => {
    const [service, setService] = useState<Service | null>(null);
    const [status, setStatus] = useState<number | null>(null);
    const router = useRouter();

    useEffect(() => {
        if (!router.isReady) return;

        const getData = async () => {
            const id = router.query.id as string;
            const {data, status} = await getServiceByIdRequest(id);
            if (data && status === 200) {
                data.images = data.images.map((image) => getFile(image));
                data.freelancer.avatar = getFile(data.freelancer.avatar);
                data.reviews = data.reviews.map((review) => {
                    review.customer.avatar = getFile(review.customer.avatar);
                    return review;
                });

                setService(data);
                setStatus(status);
            } else setStatus(status)
        }

        getData();
    }, [router]);

    return {service, status};
}